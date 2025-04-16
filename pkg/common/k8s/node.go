package k8s

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/krack8/lighthouse/pkg/common/dto"
	"github.com/krack8/lighthouse/pkg/common/log"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

type NodeServiceInterface interface {
	GetNodeList(c context.Context, p GetNodeListInputParams) (interface{}, error)
	GetNodeDetails(c context.Context, p GetNodeInputParams) (interface{}, error)
	NodeCordon(c context.Context, p NodeCordonInputParams) (interface{}, error)
	NodeTaint(c context.Context, p NodeTaintInputParams) (interface{}, error)
	NodeUnTaint(c context.Context, p NodeUnTaintInputParams) (interface{}, error)
}

type nodeService struct{}

var ns nodeService

func NodeService() *nodeService {
	return &ns
}

const (
	NODE_API_VERSION = "v1"
	NODE_KIND        = "Node"
)

type ListOutput struct {
	Result    []corev1.Node
	Metrics   []v1beta1.NodeMetrics
	GraphView GraphView `json:"graph_view"`
}

type GraphView struct {
	DeployedPodCount      int     `json:"deployed_pod_count"`
	PodCapacity           int     `json:"pod_capacity"`
	NodeCpuCapacity       float64 `json:"node_cpu_capacity"`
	NodeMemoryCapacity    float64 `json:"node_memory_capacity"`
	NodeCpuAllocatable    float64 `json:"node_cpu_allocatable"`
	NodeMemoryAllocatable float64 `json:"node_memory_allocatable"`
	NodeCpuUsage          float64 `json:"node_cpu_usage"`
	NodeMemoryUsage       float64 `json:"node_memory_usage"`
}

type GetNodeListInputParams struct {
	Search string
	Labels map[string]string
	output ListOutput
}

type DetailsOutput struct {
	Result           corev1.Node
	Metrics          v1beta1.NodeMetrics
	DeployedPodCount int `json:"deployed_pod_count"`
}

type GetNodeInputParams struct {
	NodeName string
	output   DetailsOutput
}

func (p *GetNodeListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching node list")
	nodesClient := GetKubeClientSet().CoreV1().Nodes()
	listOptions := metav1.ListOptions{}
	if p.Labels != nil {
		labelSelector := metav1.LabelSelector{MatchLabels: p.Labels}
		listOptions = metav1.ListOptions{
			LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
		}
	}
	if p.Search != "" {
		listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
	}
	nodeList, err := nodesClient.List(context.Background(), listOptions)
	if err != nil {
		log.Logger.Errorw("Failed to process get node list", "err", err.Error())
		return err
	}
	p.output.GraphView.PodCapacity = 0
	p.output.GraphView.NodeCpuAllocatable = 0
	p.output.GraphView.NodeMemoryAllocatable = 0
	p.output.GraphView.NodeCpuCapacity = 0
	p.output.GraphView.NodeMemoryAllocatable = 0
	for _, node := range nodeList.Items {
		p.output.GraphView.PodCapacity = p.output.GraphView.PodCapacity + int(node.Status.Capacity.Pods().Value())
		p.output.GraphView.NodeCpuCapacity = float64(node.Status.Capacity.Cpu().MilliValue()/1000.0) + p.output.GraphView.NodeCpuCapacity
		p.output.GraphView.NodeMemoryCapacity = (node.Status.Capacity.Memory().AsApproximateFloat64() / (1024 * 1024 * 1024)) + p.output.GraphView.NodeMemoryCapacity
		p.output.GraphView.NodeCpuAllocatable = float64(node.Status.Allocatable.Cpu().MilliValue()/1000.0) + p.output.GraphView.NodeCpuAllocatable
		p.output.GraphView.NodeMemoryAllocatable = (node.Status.Allocatable.Memory().AsApproximateFloat64() / (1024 * 1024 * 1024)) + p.output.GraphView.NodeMemoryAllocatable
	}
	metricsClient := GetMetricsClientSet().MetricsV1beta1().NodeMetricses()
	nodeMetricsList, err := metricsClient.List(context.Background(), listOptions)
	if err != nil {
		log.Logger.Errorw("Failed to process get node list metrics", "err", err.Error())
		p.output.Metrics = []v1beta1.NodeMetrics{}
		//return err
	}
	p.output.GraphView.NodeCpuUsage = 0
	p.output.GraphView.NodeMemoryUsage = 0
	if err == nil {
		p.output.Metrics = nodeMetricsList.Items
		for _, metric := range nodeMetricsList.Items {
			p.output.GraphView.NodeCpuUsage = float64(metric.Usage.Cpu().MilliValue()/1000.0) + p.output.GraphView.NodeCpuUsage
			p.output.GraphView.NodeMemoryUsage = metric.Usage.Memory().AsApproximateFloat64()/(1024*1024*1024) + p.output.GraphView.NodeMemoryUsage
		}
	}

	podList, err := GetKubeClientSet().CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to process get pod count", "err", err.Error())
	}
	// Count the total number of pods
	p.output.GraphView.DeployedPodCount = len(podList.Items)
	//for _, node := range nodeList.Items {
	//	log.Logger.Infow("%s\n", node.Name, "val", "node")
	//	for _, condition := range node.Status.Conditions {
	//		log.Logger.Infow(fmt.Sprintf("\t%s: %s\n", condition.Type, condition.Status), "val", "status")
	//	}
	//	log.Logger.
	//	Infow("Taint\n", "val", "taint")
	//	for _, taint := range node.Spec.Taints {
	//		log.Logger.Infow(fmt.Sprintf("\t%s: %s %s\n", taint.Effect, taint.Key, taint.Value), "val", "taints")
	//	}
	//}
	p.output.Result = nodeList.Items
	return nil
}

func (svc *nodeService) GetNodeList(c context.Context, p GetNodeListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

func (p *GetNodeInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching node details of ....", p.NodeName)
	nodesClient := GetKubeClientSet().CoreV1().Nodes()
	output, err := nodesClient.Get(context.Background(), p.NodeName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to fetch node details", "err", err.Error())
		return err
	}
	metricsClient := GetMetricsClientSet().MetricsV1beta1().NodeMetricses()
	nodeMetrics, err := metricsClient.Get(context.Background(), p.NodeName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to process get node metrics", "err", err.Error())
		p.output.Metrics = v1beta1.NodeMetrics{}
		//return err
	}
	fieldSelector := fmt.Sprintf("spec.nodeName=%s", p.NodeName)
	podList, err := GetKubeClientSet().CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{FieldSelector: fieldSelector})
	if err != nil {
		log.Logger.Errorw("Failed to fetch pod count", "err", err.Error())
	}
	p.output.DeployedPodCount = len(podList.Items)
	p.output.Result = *output
	p.output.Result.APIVersion = NODE_API_VERSION
	p.output.Result.Kind = NODE_KIND
	if err == nil {
		p.output.Metrics = *nodeMetrics
	}
	return nil
}

func (svc *nodeService) GetNodeDetails(c context.Context, p GetNodeInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type NodeCordonOutput struct {
	Node string `json:"node"`
	Task string `json:"task"`
}

type NodeCordonInputParams struct {
	NodeName string
	output   NodeCordonOutput
}

type patchNodeString struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value bool   `json:"value"`
}

const (
	UncordonStr = "uncordon"
	CordonStr   = "cordon"
)

func (p *NodeCordonInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching node details of ....", p.NodeName)
	nodesClient := GetKubeClientSet().CoreV1().Nodes()
	node, err := nodesClient.Get(context.Background(), p.NodeName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to fetch node details", "err", err.Error())
		return err
	}
	if node.Spec.Unschedulable {
		log.Logger.Infow("patch uncordon node "+p.NodeName, "info", "uncordon")
		payload := []patchNodeString{{
			Op:    "replace",
			Path:  "/spec/unschedulable",
			Value: false,
		}}
		payloadBytes, _ := json.Marshal(payload)
		_, err = nodesClient.Patch(context.Background(), p.NodeName, types.JSONPatchType, payloadBytes, metav1.PatchOptions{})
		if err != nil {
			log.Logger.Errorw("Failed to uncordon node", "err", err.Error())
			return err
		}
		p.output.Task = UncordonStr
	} else {
		log.Logger.Infow("patch cordon node "+p.NodeName, "info", "cordon")
		payload := []patchNodeString{{
			Op:    "replace",
			Path:  "/spec/unschedulable",
			Value: true,
		}}
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			log.Logger.Errorw("Error Marshal", "err", err.Error())
			return err
		}
		_, err = nodesClient.Patch(context.Background(), p.NodeName, types.JSONPatchType, payloadBytes, metav1.PatchOptions{})
		if err != nil {
			log.Logger.Errorw("Failed to cordon node", "err", err.Error())
			return err
		}
		p.output.Task = CordonStr
	}
	p.output.Node = p.NodeName
	return nil
}

func (svc *nodeService) NodeCordon(c context.Context, p NodeCordonInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type NodeTaintInputParams struct {
	NodeName  string
	TaintList *dto.TaintList
}

func (p *NodeTaintInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching node taint of ....", p.NodeName)
	nodesClient := GetKubeClientSet().CoreV1().Nodes()
	node, err := nodesClient.Get(context.Background(), p.NodeName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to fetch node details", "err", err.Error())
		return err
	}
	updated := false
	for i, newTaint := range p.TaintList.Taint {
	loop:
		for _, existingTaint := range node.Spec.Taints {
			if existingTaint.Key == newTaint.Key {
				existingTaint.Effect = newTaint.Effect
				existingTaint.Value = newTaint.Value
				node.Spec.Taints[i] = existingTaint
				updated = true
				break loop
			}
		}
		if !updated {
			node.Spec.Taints = append(node.Spec.Taints, newTaint)
		}
		updated = false
	}

	_, err = nodesClient.Update(context.Background(), node, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to update node taint", "err", err.Error())
		return err
	}
	return nil
}

func (svc *nodeService) NodeTaint(c context.Context, p NodeTaintInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}

type NodeUnTaintInputParams struct {
	NodeName string
	Keys     []string
}

func (p *NodeUnTaintInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching node untaint of ....", p.NodeName)
	nodesClient := GetKubeClientSet().CoreV1().Nodes()
	node, err := nodesClient.Get(context.Background(), p.NodeName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to fetch node details", "err", err.Error())
		return err
	}
	newTaints := []corev1.Taint{}
	if len(p.Keys) > 0 {
		found := false
		for _, taint := range node.Spec.Taints {
		loop2:
			for _, key := range p.Keys {
				if taint.Key == key {
					found = true
					break loop2
				}
			}
			if !found {
				newTaints = append(newTaints, taint)
			} else {
				found = false
			}
		}
	}

	node.Spec.Taints = newTaints

	_, err = nodesClient.Update(context.Background(), node, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to update node untaint", "err", err.Error())
		return err
	}
	return nil
}

func (svc *nodeService) NodeUnTaint(c context.Context, p NodeUnTaintInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
