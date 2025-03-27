package k8s

import (
	"bytes"
	"context"
	"github.com/gorilla/websocket"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/log"
	"io"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	_v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	"strconv"
	"strings"
	"time"
)

type PodServiceInterface interface {
	GetPodList(c context.Context, p GetPodListInputParams) (interface{}, error)
	GetPodDetails(c context.Context, p GetPodDetailsInputParams) (interface{}, error)
	GetPodStats(c context.Context, p GetPodStatsInputParams) (interface{}, error)
	GetPodLogs(c context.Context, p GetPodLogsInputParams) (interface{}, error)
	GetPodLogsStream(c context.Context, p GetPodLogsStreamInputParams, conn *websocket.Conn) error
	DeployPod(c context.Context, p DeployPodInputParams) (interface{}, error)
	DeletePod(c context.Context, p DeletePodInputParams) (interface{}, error)
}

type podService struct{}

var ps podService

func PodService() *podService {
	return &ps
}

const (
	PodLogBoolTrue     = "y"
	PodLogBoolFalse    = "n"
	TailLinesThreshold = int64(2500)
)

type OutputPodList struct {
	Result    []corev1.Pod
	Resource  string
	Remaining int64
	Total     int
}

type GetPodListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputPodList
}

type PodExecInputParams struct {
	PodName       string
	NamespaceName string
	ContainerName string
}

func (p *GetPodListInputParams) Find(c context.Context, podClient _v1.PodInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "pod")
	filteredPods := []corev1.Pod{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		podList, err := podClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get pod list", "err", err.Error())
			return err
		}

		for _, pod := range podList.Items {
			if strings.Contains(pod.Name, p.Search) {
				filteredPods = append(filteredPods, pod)
			}
		}
		length = len(filteredPods)
		nextPageToken = podList.Continue
		if podList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		podList, err := podClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get pod list", "err", err.Error())
			return err
		}
		for _, namespace := range podList.Items {
			if strings.Contains(namespace.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredPods
	p.output.Total = len(filteredPods)
	p.output.Remaining = int64(remaining)

	log.Logger.Info("pods", "count", len(filteredPods))
	return nil
}

func (p *GetPodListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching pod list")
	podClient := GetKubeClientSet().CoreV1().Pods(p.NamespaceName)

	limit := config.PageLimit
	if p.Limit != "" {
		limit, _ = strconv.ParseInt(p.Limit, 10, 64)
	}
	listOptions := metav1.ListOptions{Limit: limit, Continue: p.Continue}
	if p.Labels != nil {
		labelSelector := metav1.LabelSelector{MatchLabels: p.Labels}
		listOptions = metav1.ListOptions{
			LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
		}
	}
	var err error
	var podList *corev1.PodList
	if p.Search != "" {
		err = p.Find(c, podClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get pod list", "err", err.Error())
			return err
		}
		return nil
	} else {
		podList, err = podClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get pod list", "err", err.Error())
			return err
		}

		podList, err = podClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get pod list", "err", err.Error())
			return err
		}
		remaining := podList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: podList.Continue}
				res, err := podClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get pod list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = podList.Items
		p.output.Total = len(podList.Items)
		p.output.Resource = podList.Continue
	}
	return nil
}

func (svc *podService) GetPodList(c context.Context, p GetPodListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type OutputPodDetails struct {
	Result corev1.Pod
	CPU    float64
	Memory float64
}
type GetPodDetailsInputParams struct {
	NamespaceName string
	PodName       string
	output        OutputPodDetails
}

func (p *GetPodDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching pod details of ....", p.NamespaceName)
	podsClient := GetKubeClientSet().CoreV1().Pods(p.NamespaceName)
	pod, err := podsClient.Get(context.Background(), p.PodName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get pod ", p.PodName, "err", err.Error())
		return err
	}
	podMetrics, err := GetMetricsClientSet().MetricsV1beta1().PodMetricses(p.NamespaceName).Get(context.TODO(), p.PodName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get pod metrics list", "err", err.Error())
	}
	p.output.CPU = 0
	p.output.Memory = 0
	if podMetrics != nil {
		for _, containerMetric := range podMetrics.Containers {
			p.output.CPU += float64(containerMetric.Usage.Cpu().MilliValue()) / 1000.0
			p.output.Memory += float64(containerMetric.Usage.Memory().Value()) / (1024 * 1024 * 1024)
		}
	}
	p.output.Result = *pod
	return nil
}

func (svc *podService) GetPodDetails(c context.Context, p GetPodDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployPodInputParams struct {
	Pod    *corev1.Pod
	output *corev1.Pod
}

func (p *DeployPodInputParams) PostProcess(c context.Context) error {
	p.output.ManagedFields = nil
	return nil
}

func (p *DeployPodInputParams) Process(c context.Context) error {
	podClient := GetKubeClientSet().CoreV1().Pods(p.Pod.Namespace)
	returnedPod, err := podClient.Get(context.Background(), p.Pod.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating pod in namespace "+p.Pod.Namespace, "value", p.Pod.Name)
		p.output, err = podClient.Create(context.Background(), p.Pod, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create pod in namespace "+p.Pod.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("pod created")
	} else {
		log.Logger.Infow("Pod exist in namespace "+p.Pod.Namespace, "value", p.Pod.Name)
		log.Logger.Infow("Updating pod in namespace "+p.Pod.Namespace, "value", p.Pod.Name)
		p.Pod.SetResourceVersion(returnedPod.ResourceVersion)
		p.output, err = podClient.Update(context.Background(), p.Pod, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update pod ", p.Pod.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("pod updated")
	}
	return nil
}

func (svc *podService) DeployPod(c context.Context, p DeployPodInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	_ = p.PostProcess(c)

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeletePodInputParams struct {
	NamespaceName string
	PodName       string
}

func (p *DeletePodInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting pod of ....", p.NamespaceName)
	podClient := GetKubeClientSet().CoreV1().Pods(p.NamespaceName)
	_, err := podClient.Get(context.Background(), p.PodName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get pod ", p.PodName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = podClient.Delete(context.Background(), p.PodName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete pod ", p.PodName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *podService) DeletePod(c context.Context, p DeletePodInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}

type StatsPod struct {
	Total   int
	Running int
	Pending int
	Failed  int
	CPU     float64
	Memory  float64
}

func (s *StatsPod) New() *StatsPod {
	return &StatsPod{Total: 0, Running: 0, Pending: 0, Failed: 0, CPU: 0, Memory: 0}
}

type GetPodStatsInputParams struct {
	NamespaceName string
	Search        string
	Labels        map[string]string
	output        *StatsPod
}

func (p *GetPodStatsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching pod list stats")
	podClient := GetKubeClientSet().CoreV1().Pods(p.NamespaceName)

	p.output = p.output.New()
	listOptions := metav1.ListOptions{}
	if p.Labels != nil {
		labelSelector := metav1.LabelSelector{MatchLabels: p.Labels}
		listOptions = metav1.ListOptions{
			LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
		}
	}
	podList, err := podClient.List(context.Background(), listOptions)
	if err != nil {
		log.Logger.Errorw("Failed to get pod list stats", "err", err.Error())
		return err
	}
	podMetricsList := []v1beta1.PodMetrics{}
	podMetrics, err := GetMetricsClientSet().MetricsV1beta1().PodMetricses(p.NamespaceName).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get pod metrics list", "err", err.Error())
	} else {
		podMetricsList = podMetrics.Items
	}

	if p.Search != "" {
		filteredPods := make(map[string]bool)
		for _, pod := range podList.Items {
			if strings.Contains(pod.Name, p.Search) {
				p.output.Total += 1
				filteredPods[pod.Name] = true
				switch pod.Status.Phase {
				case RUNNING:
					p.output.Running += 1
				case FAILED:
					p.output.Failed += 1
				case PENDING:
					p.output.Pending += 1
				}
			}
		}

		for _, podMetric := range podMetricsList {
			if filteredPods[podMetric.Name] {
				for _, containerMetric := range podMetric.Containers {
					p.output.CPU += float64(containerMetric.Usage.Cpu().MilliValue()) / 1000.0
					p.output.Memory += float64(containerMetric.Usage.Memory().Value()) / (1024 * 1024 * 1024)
				}
			}
		}

		return nil
	}

	p.output.Total = len(podList.Items)

	for _, obj := range podList.Items {
		switch obj.Status.Phase {
		case RUNNING:
			p.output.Running += 1
		case FAILED:
			p.output.Failed += 1
		case PENDING:
			p.output.Pending += 1
		}
	}
	for _, podMetric := range podMetricsList {
		for _, containerMetric := range podMetric.Containers {
			p.output.CPU += float64(containerMetric.Usage.Cpu().MilliValue()) / 1000.0
			p.output.Memory += float64(containerMetric.Usage.Memory().Value()) / (1024 * 1024 * 1024)
		}
	}
	return nil
}

func (svc *podService) GetPodStats(c context.Context, p GetPodStatsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetPodLogsInputParams struct {
	NamespaceName string
	Pod           string
	Container     string
	TailLines     *int64
	Timestamps    string
	SinceSeconds  *int64
	Previous      string
	output        string
}

func (p *GetPodLogsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching pod container logs", "debug", "logs")
	podClient := GetKubeClientSet().CoreV1().Pods(p.NamespaceName)
	podLogOptions := corev1.PodLogOptions{Follow: false}
	if p.Container != "" {
		podLogOptions.Container = p.Container
	}
	zero := int64(0)
	if p.TailLines != nil && *p.TailLines != zero {
		if *p.TailLines > TailLinesThreshold {
			threshold := TailLinesThreshold
			podLogOptions.TailLines = &threshold
		} else {
			podLogOptions.TailLines = p.TailLines
		}
	}

	switch p.Timestamps {
	case PodLogBoolTrue:
		podLogOptions.Timestamps = true
	case PodLogBoolFalse:
		podLogOptions.Timestamps = false
	}

	switch p.Previous {
	case PodLogBoolTrue:
		podLogOptions.Previous = true
	case PodLogBoolFalse:
		podLogOptions.Previous = false
	}

	if p.SinceSeconds != nil && *p.SinceSeconds > zero {
		podLogOptions.SinceSeconds = p.SinceSeconds
	}

	podLogRequest := podClient.GetLogs(p.Pod, &podLogOptions)
	podLogs, err := podLogRequest.Stream(context.Background())
	if err != nil {
		return err
	}
	defer podLogs.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return err
	}
	p.output = buf.String()
	return nil
}

func (svc *podService) GetPodLogs(c context.Context, p GetPodLogsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}
	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetPodLogsStreamInputParams struct {
	NamespaceName string
	Pod           string
	Container     string
	TailLines     *int64
	Timestamps    string
	SinceSeconds  *int64
	output        []byte
}

func (p *GetPodLogsStreamInputParams) Process(c context.Context, conn *websocket.Conn) error {
	log.Logger.Debugw("fetching pod container stream logs", "debug", "logs")
	podClient := GetKubeClientSet().CoreV1().Pods(p.NamespaceName)
	podLogOptions := corev1.PodLogOptions{Follow: true}
	if p.Container != "" {
		podLogOptions.Container = p.Container
	}
	zero := int64(0)
	if p.TailLines != nil && *p.TailLines != zero {
		if *p.TailLines > TailLinesThreshold {
			threshold := TailLinesThreshold
			podLogOptions.TailLines = &threshold
		} else {
			podLogOptions.TailLines = p.TailLines
		}
	}
	if p.SinceSeconds != nil && *p.SinceSeconds > zero {
		podLogOptions.SinceSeconds = p.SinceSeconds
	}
	// checking timestamps
	switch p.Timestamps {
	case PodLogBoolTrue:
		podLogOptions.Timestamps = true
	case PodLogBoolFalse:
		podLogOptions.Timestamps = false
	}

	podLogRequest := podClient.GetLogs(p.Pod, &podLogOptions)
	stream, err := podLogRequest.Stream(context.TODO())
	if err != nil {
		return err
	}
	defer stream.Close()
	for {
		buf := make([]byte, 2000)
		numBytes, err := stream.Read(buf)
		if err == io.EOF {
			break
		}
		if numBytes == 0 {
			//time.Sleep(1000 * time.Millisecond)
			continue
		}
		if err != nil {
			return err
		}
		p.output = buf[:numBytes]
		err = conn.WriteMessage(websocket.TextMessage, p.output)
		if err != nil {
			break
		}
		time.Sleep(time.Second)
	}
	return nil
}

func (svc *podService) GetPodLogsStream(c context.Context, p GetPodLogsStreamInputParams, conn *websocket.Conn) error {
	err := p.Process(c, conn)
	if err != nil {
		return err
	}
	return nil
}

type PodDetailsGrafanaInputParams struct {
	NamespaceName string
	PodName       string
	output        string
}

func (p *PodDetailsGrafanaInputParams) Process(c context.Context) error {
	log.Logger.Debugw("setting grafana pod details of ....", p.NamespaceName)
	//url := fmt.Sprintf("%s/d/%s?apiKey=%s", grafanaURL, dashboardID, apiKey)
	//p.output = url
	return nil
}

func (svc *podService) PodDetailsGrafana(c context.Context, p PodDetailsGrafanaInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}
