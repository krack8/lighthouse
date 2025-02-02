package k8s

import (
	"context"
	"errors"
	"fmt"
	cfg "github.com/krack8/lighthouse/pkg/config"
	"github.com/krack8/lighthouse/pkg/log"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"strconv"
	"strings"
)

type DeploymentServiceInterface interface {
	GetDeploymentList(c context.Context, p GetDeploymentListInputParams) (interface{}, error)
	GetDeploymentDetails(c context.Context, p GetDeploymentDetailsInputParams) (interface{}, error)
	GetDeploymentStats(c context.Context, p GetDeploymentStatsInputParams) (interface{}, error)
	DeployDeployment(c context.Context, p DeployDeploymentInputParams) (interface{}, error)
	DeleteDeployment(c context.Context, p DeleteDeploymentInputParams) (interface{}, error)
	GetDeploymentPodList(c context.Context, p GetDeploymentPodListInputParams) (interface{}, error)
}

type deploymentService struct{}

var ds deploymentService

func DeploymentService() *deploymentService {
	return &ds
}

type OutputDeploymentList struct {
	Result    []appsv1.Deployment
	Resource  string
	Remaining int64
	Total     int
}

type GetDeploymentListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputDeploymentList
}

func (p *GetDeploymentListInputParams) PostProcess(c context.Context) error {
	for idx, _ := range p.output.Result {
		p.output.Result[idx].ManagedFields = nil
	}
	return nil
}

func (p *GetDeploymentListInputParams) Find(c context.Context, deploymentClient v1.DeploymentInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "deployment")
	filteredDeployments := []appsv1.Deployment{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		deploymentList, err := deploymentClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get deployment list", "err", err.Error())
			return err
		}

		for _, deployment := range deploymentList.Items {
			if strings.Contains(deployment.Name, p.Search) {
				filteredDeployments = append(filteredDeployments, deployment)
			}
		}
		length = len(filteredDeployments)
		nextPageToken = deploymentList.Continue
		if deploymentList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		deploymentList, err := deploymentClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get deployment list", "err", err.Error())
			return err
		}
		for _, deployment := range deploymentList.Items {
			if strings.Contains(deployment.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredDeployments
	p.output.Total = len(filteredDeployments)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetDeploymentListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching deployment list")
	deploymentClient := cfg.GetKubeClientSet().AppsV1().Deployments(p.NamespaceName)
	limit := cfg.PageLimit
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
	var deploymentList *appsv1.DeploymentList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, deploymentClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get deployment list", "err", err.Error())
			return err
		}
		return nil
	} else {
		deploymentList, err = deploymentClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get deployment list", "err", err.Error())
			return err
		}

		deploymentList, err = deploymentClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get deployment list", "err", err.Error())
			return err
		}
		remaining := deploymentList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: deploymentList.Continue}
				res, err := deploymentClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get deployment list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = deploymentList.Items
		p.output.Total = len(deploymentList.Items)
		p.output.Resource = deploymentList.Continue
	}
	return nil
}

func (svc *deploymentService) GetDeploymentList(c context.Context, p GetDeploymentListInputParams) (interface{}, error) {
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

type GetDeploymentDetailsInputParams struct {
	NamespaceName  string
	DeploymentName string
	output         appsv1.Deployment
}

func (p *GetDeploymentDetailsInputParams) PostProcess(c context.Context) error {
	p.output.ManagedFields = nil
	return nil
}

func (p *GetDeploymentDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching deployment details of ....", p.NamespaceName)
	deploymentsClient := cfg.GetKubeClientSet().AppsV1().Deployments(p.NamespaceName)
	output, err := deploymentsClient.Get(context.Background(), p.DeploymentName, metav1.GetOptions{})
	/////
	//var replicasets []string
	//for _, i := range output.Status.Conditions {
	//	if i.Type == "Progressing" {
	//		content := i.Message
	//		re := regexp.MustCompile(`\"(.*)\"`)
	//		match := re.FindStringSubmatch(content)
	//		if len(match) > 1 {
	//			fmt.Println("match found -", match[1])
	//			replicasets = append(replicasets, match[1])
	//		} else {
	//			fmt.Println("match not found")
	//		}
	//	}
	//}
	//fmt.Println(replicasets)
	////
	if err != nil {
		log.Logger.Errorw("Failed to get deployment ", p.DeploymentName, "err", err.Error())
		return err
	}

	p.output = *output
	return nil
}

func (svc *deploymentService) GetDeploymentDetails(c context.Context, p GetDeploymentDetailsInputParams) (interface{}, error) {
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

type DeployDeploymentInputParams struct {
	Deployment *appsv1.Deployment
	output     *appsv1.Deployment
}

func (p *DeployDeploymentInputParams) PostProcess(c context.Context) error {
	p.output.ManagedFields = nil
	return nil
}

func (p *DeployDeploymentInputParams) Process(c context.Context) error {
	deploymentClient := cfg.GetKubeClientSet().AppsV1().Deployments(p.Deployment.Namespace)
	_, err := deploymentClient.Get(context.Background(), p.Deployment.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating deployment in namespace "+p.Deployment.Namespace, "value", p.Deployment.Name)
		p.output, err = deploymentClient.Create(context.Background(), p.Deployment, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create deployment in namespace "+p.Deployment.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("deployment created")
	} else {
		log.Logger.Infow("Deployment exist in namespace "+p.Deployment.Namespace, "value", p.Deployment.Name)
		log.Logger.Infow("Updating deployment in namespace "+p.Deployment.Namespace, "value", p.Deployment.Name)
		p.output, err = deploymentClient.Update(context.Background(), p.Deployment, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update deployment ", p.Deployment.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("deployment updated")
	}
	return nil
}

func (svc *deploymentService) DeployDeployment(c context.Context, p DeployDeploymentInputParams) (interface{}, error) {
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

type DeleteDeploymentInputParams struct {
	NamespaceName  string
	DeploymentName string
}

func (p *DeleteDeploymentInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting deployment of ....", p.NamespaceName)
	deploymentClient := cfg.GetKubeClientSet().AppsV1().Deployments(p.NamespaceName)
	_, err := deploymentClient.Get(context.Background(), p.DeploymentName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get deployment ", p.DeploymentName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = deploymentClient.Delete(context.Background(), p.DeploymentName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete deployment ", p.DeploymentName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *deploymentService) DeleteDeployment(c context.Context, p DeleteDeploymentInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}

type StatsDeployment struct {
	Total       int
	Ready       int
	NotReady    int
	TotalCPU    float64
	TotalMemory float64
}

func (s *StatsDeployment) New() *StatsDeployment {
	return &StatsDeployment{Total: 0, Ready: 0, NotReady: 0}
}

type GetDeploymentStatsInputParams struct {
	NamespaceName string
	Labels        map[string]string
	Search        string
	output        *StatsDeployment
}

func (p *GetDeploymentStatsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching deployment list stats")
	deploymentClient := cfg.GetKubeClientSet().AppsV1().Deployments(p.NamespaceName)

	listOptions := metav1.ListOptions{}
	if p.Labels != nil {
		labelSelector := metav1.LabelSelector{MatchLabels: p.Labels}
		listOptions = metav1.ListOptions{
			LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
		}
	}

	deploymentList, err := deploymentClient.List(context.Background(), listOptions)
	if err != nil {
		log.Logger.Errorw("Failed to get deployment list stats", "err", err.Error())
		return err
	}

	totalCPU := float64(0)
	totalMemory := float64(0)

	if p.Search != "" {
		listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		filteredDeployments := []appsv1.Deployment{}

		for _, deployment := range deploymentList.Items {
			if strings.Contains(deployment.Name, p.Search) {
				filteredDeployments = append(filteredDeployments, deployment)
			}
		}

		p.output = p.output.New()
		for _, obj := range filteredDeployments {
			p.output.Total += int(obj.Status.Replicas)
			p.output.Ready += int(obj.Status.ReadyReplicas)
			podMetricsList, err := cfg.GetMetricsClientSet().MetricsV1beta1().PodMetricses(p.NamespaceName).List(context.TODO(), metav1.ListOptions{LabelSelector: fmt.Sprintf("app=%s", obj.Labels["app"])})
			if err != nil {
				log.Logger.Errorw("Failed to get pod metrics", "err", err.Error())
				return err
			}
			for _, podMetrics := range podMetricsList.Items {
				for _, container := range podMetrics.Containers {
					totalCPU += float64(container.Usage.Cpu().MilliValue()) / 1000.0
					totalMemory += float64(container.Usage.Memory().Value()) / (1024 * 1024 * 1024)
				}
			}
		}

		p.output.NotReady = p.output.Total - p.output.Ready
		p.output.TotalCPU = totalCPU
		p.output.TotalMemory = totalMemory
		return nil
	}

	p.output = p.output.New()
	for _, obj := range deploymentList.Items {
		p.output.Total += int(obj.Status.Replicas)
		p.output.Ready += int(obj.Status.ReadyReplicas)

		podMetricsList, err := cfg.GetMetricsClientSet().MetricsV1beta1().PodMetricses(p.NamespaceName).List(context.TODO(), metav1.ListOptions{LabelSelector: fmt.Sprintf("app=%s", obj.Labels["app"])})
		if err != nil {
			log.Logger.Errorw("Failed to get pod metrics", "err", err.Error())
			return err
		}
		for _, podMetrics := range podMetricsList.Items {
			for _, container := range podMetrics.Containers {
				totalCPU += float64(container.Usage.Cpu().MilliValue()) / 1000.0
				totalMemory += float64(container.Usage.Memory().Value()) / (1024 * 1024 * 1024)
			}
		}
	}
	p.output.NotReady = p.output.Total - p.output.Ready
	p.output.TotalCPU = totalCPU
	p.output.TotalMemory = totalMemory
	return nil
}

func (svc *deploymentService) GetDeploymentStats(c context.Context, p GetDeploymentStatsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeploymentPodOutput struct {
	PodList     []corev1.Pod
	Resource    string
	Remaining   int64
	TotalCPU    float64
	TotalMemory float64
}

type GetDeploymentPodListInputParams struct {
	NamespaceName  string
	DeploymentName string
	Replicaset     string
	Limit          string
	Labels         map[string]string
	Search         string
	Continue       string
	output         DeploymentPodOutput
}

const PodTemplateHash = "pod-template-hash"

func (p *GetDeploymentPodListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching replicaset details of ...."+p.NamespaceName, "service", "deployment-pod-list")
	replicaSetClient := cfg.GetKubeClientSet().AppsV1().ReplicaSets(p.NamespaceName)
	replicaSet, err := replicaSetClient.Get(context.Background(), p.Replicaset, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get deployment pod list"+p.DeploymentName, "err", err.Error())
		return err
	}
	if _, isKeyExists := replicaSet.Labels[PodTemplateHash]; !isKeyExists {
		return errors.New("unable to fetch pod list")
	}
	if replicaSet.Labels[PodTemplateHash] != "" {
		podClient := cfg.GetKubeClientSet().CoreV1().Pods(p.NamespaceName)

		limit := cfg.PageLimit
		if p.Limit != "" {
			limit, _ = strconv.ParseInt(p.Limit, 10, 64)
		}
		listOptions := metav1.ListOptions{Limit: limit, Continue: p.Continue}
		if p.Labels == nil {
			p.Labels = make(map[string]string)
		}
		p.Labels["pod-template-hash"] = replicaSet.Labels["pod-template-hash"]
		labelSelector := metav1.LabelSelector{MatchLabels: p.Labels}
		if p.Labels != nil {
			listOptions = metav1.ListOptions{
				LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
			}
		}
		if p.Search != "" {
			listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		}
		//FieldSelector: fmt.Sprintf("spec.ports[0].nodePort=%s", port),
		podList, err := podClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get pod list", "err", err.Error())
			return err
		}
		p.output.PodList = podList.Items
		totalCPU := float64(0)
		totalMemory := float64(0)
		for idx, pod := range p.output.PodList {
			p.output.PodList[idx].ManagedFields = nil
			podMetrics, err := cfg.GetMetricsClientSet().MetricsV1beta1().PodMetricses(p.NamespaceName).Get(context.TODO(), pod.Name, metav1.GetOptions{})
			if err != nil {
				log.Logger.Errorw("Failed to get pod metrics", "err", err.Error())
				return err
			}
			for _, container := range podMetrics.Containers {
				totalCPU += float64(container.Usage.Cpu().MilliValue()) / 1000.0
				totalMemory += float64(container.Usage.Memory().Value()) / (1024 * 1024 * 1024)
			}
		}
		p.output.TotalCPU = totalCPU
		p.output.TotalMemory = totalMemory
		remaining := podList.RemainingItemCount

		if remaining != nil {
			p.output.Remaining = *remaining
		} else {
			p.output.Remaining = 0
		}

		p.output.Resource = podList.Continue
	} else {
		return errors.New("unable to fetch pod list")
	}
	/////
	//var replicasets []string
	//for _, i := range output.Status.Conditions {
	//	if i.Type == "Progressing" {
	//		content := i.Message
	//		re := regexp.MustCompile(`\"(.*)\"`)
	//		match := re.FindStringSubmatch(content)
	//		if len(match) > 1 {
	//			fmt.Println("match found -", match[1])
	//			replicasets = append(replicasets, match[1])
	//		} else {
	//			fmt.Println("match not found")
	//		}
	//	}
	//}
	//fmt.Println(replicasets)
	////
	return nil
}

func (svc *deploymentService) GetDeploymentPodList(c context.Context, p GetDeploymentPodListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}
