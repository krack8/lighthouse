package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/log"
	"k8s.io/api/policy/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	_v1 "k8s.io/client-go/kubernetes/typed/policy/v1"
	"strconv"
	"strings"
)

type PodDisruptionBudgetsServiceInterface interface {
	GetPodDisruptionBudgetsList(c context.Context, p GetPodDisruptionBudgetsListInputParams) (interface{}, error)
	GetPodDisruptionBudgetsDetails(c context.Context, p GetPodDisruptionBudgetsDetailsInputParams) (interface{}, error)
	DeployPodDisruptionBudgets(c context.Context, p DeployPodDisruptionBudgetsInputParams) (interface{}, error)
	DeletePodDisruptionBudgets(c context.Context, p DeletePodDisruptionBudgetsInputParams) (interface{}, error)
}

type podDisruptionBudgetsService struct{}

var pdbs podDisruptionBudgetsService

func PodDisruptionBudgetsService() *podDisruptionBudgetsService {
	return &pdbs
}

const (
	POD_DISRUPTION_BUDGETS_API_VERSION = "policy/v1"
	POD_DISRUPTION_BUDGETS_KIND        = "PodDisruptionBudget"
)

type OutputPodDisruptionBudgetsList struct {
	Result    []v1.PodDisruptionBudget
	Resource  string
	Remaining int64
	Total     int
}

type GetPodDisruptionBudgetsListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputPodDisruptionBudgetsList
}

func (p *GetPodDisruptionBudgetsListInputParams) Find(c context.Context, podDisruptionBudgetsClient _v1.PodDisruptionBudgetInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "podDisruptionBudgets")
	filteredPodDisruptionBudgetses := []v1.PodDisruptionBudget{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		podDisruptionBudgetsList, err := podDisruptionBudgetsClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get podDisruptionBudgets list", "err", err.Error())
			return err
		}

		for _, podDisruptionBudgets := range podDisruptionBudgetsList.Items {
			if strings.Contains(podDisruptionBudgets.Name, p.Search) {
				filteredPodDisruptionBudgetses = append(filteredPodDisruptionBudgetses, podDisruptionBudgets)
			}
		}
		length = len(filteredPodDisruptionBudgetses)
		nextPageToken = podDisruptionBudgetsList.Continue
		if podDisruptionBudgetsList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		podDisruptionBudgetsList, err := podDisruptionBudgetsClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get podDisruptionBudgets list", "err", err.Error())
			return err
		}
		for _, podDisruptionBudgets := range podDisruptionBudgetsList.Items {
			if strings.Contains(podDisruptionBudgets.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredPodDisruptionBudgetses
	p.output.Total = len(filteredPodDisruptionBudgetses)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetPodDisruptionBudgetsListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching podDisruptionBudgets list")
	podDisruptionBudgetsClient := GetKubeClientSet().PolicyV1().PodDisruptionBudgets(p.NamespaceName)
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
	var podDisruptionBudgetsList *v1.PodDisruptionBudgetList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, podDisruptionBudgetsClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get podDisruptionBudgets list", "err", err.Error())
			return err
		}
		return nil
	} else {
		podDisruptionBudgetsList, err = podDisruptionBudgetsClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get podDisruptionBudgets list", "err", err.Error())
			return err
		}

		podDisruptionBudgetsList, err = podDisruptionBudgetsClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get podDisruptionBudgets list", "err", err.Error())
			return err
		}
		remaining := podDisruptionBudgetsList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: podDisruptionBudgetsList.Continue}
				res, err := podDisruptionBudgetsClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get podDisruptionBudgets list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = podDisruptionBudgetsList.Items
		p.output.Total = len(podDisruptionBudgetsList.Items)
		p.output.Resource = podDisruptionBudgetsList.Continue
	}
	return nil
}

func (svc *podDisruptionBudgetsService) GetPodDisruptionBudgetsList(c context.Context, p GetPodDisruptionBudgetsListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}
	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetPodDisruptionBudgetsDetailsInputParams struct {
	NamespaceName            string
	PodDisruptionBudgetsName string
	output                   v1.PodDisruptionBudget
}

func (p *GetPodDisruptionBudgetsDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching podDisruptionBudgets details of ....", p.NamespaceName)
	podDisruptionBudgetssClient := GetKubeClientSet().PolicyV1().PodDisruptionBudgets(p.NamespaceName)
	output, err := podDisruptionBudgetssClient.Get(context.Background(), p.PodDisruptionBudgetsName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get podDisruptionBudgets ", p.PodDisruptionBudgetsName, "err", err.Error())
		return err
	}
	p.output = *output
	p.output.APIVersion = POD_DISRUPTION_BUDGETS_API_VERSION
	p.output.Kind = POD_DISRUPTION_BUDGETS_KIND
	return nil
}

func (svc *podDisruptionBudgetsService) GetPodDisruptionBudgetsDetails(c context.Context, p GetPodDisruptionBudgetsDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployPodDisruptionBudgetsInputParams struct {
	PodDisruptionBudgets *v1.PodDisruptionBudget
	output               *v1.PodDisruptionBudget
}

func (p *DeployPodDisruptionBudgetsInputParams) Process(c context.Context) error {
	PodDisruptionBudgetsClient := GetKubeClientSet().PolicyV1().PodDisruptionBudgets(p.PodDisruptionBudgets.Namespace)
	_, err := PodDisruptionBudgetsClient.Get(context.Background(), p.PodDisruptionBudgets.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating podDisruptionBudgets in namespace "+p.PodDisruptionBudgets.Namespace, "value", p.PodDisruptionBudgets.Name)
		p.output, err = PodDisruptionBudgetsClient.Create(context.Background(), p.PodDisruptionBudgets, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create podDisruptionBudgets in namespace "+p.PodDisruptionBudgets.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("podDisruptionBudgets created")
	} else {
		log.Logger.Infow("podDisruptionBudgets exist in namespace "+p.PodDisruptionBudgets.Namespace, "value", p.PodDisruptionBudgets.Name)
		log.Logger.Infow("Updating podDisruptionBudgets in namespace "+p.PodDisruptionBudgets.Namespace, "value", p.PodDisruptionBudgets.Name)
		p.output, err = PodDisruptionBudgetsClient.Update(context.Background(), p.PodDisruptionBudgets, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update podDisruptionBudgets ", p.PodDisruptionBudgets.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("podDisruptionBudgets updated")
	}
	return nil
}

func (svc *podDisruptionBudgetsService) DeployPodDisruptionBudgets(c context.Context, p DeployPodDisruptionBudgetsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeletePodDisruptionBudgetsInputParams struct {
	NamespaceName            string
	PodDisruptionBudgetsName string
}

func (p *DeletePodDisruptionBudgetsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting PodDisruptionBudgets of ....", p.NamespaceName)
	PodDisruptionBudgetsClient := GetKubeClientSet().PolicyV1().PodDisruptionBudgets(p.NamespaceName)
	_, err := PodDisruptionBudgetsClient.Get(context.Background(), p.PodDisruptionBudgetsName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get PodDisruptionBudgets ", p.PodDisruptionBudgetsName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = PodDisruptionBudgetsClient.Delete(context.Background(), p.PodDisruptionBudgetsName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete PodDisruptionBudgets ", p.PodDisruptionBudgetsName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *podDisruptionBudgetsService) DeletePodDisruptionBudgets(c context.Context, p DeletePodDisruptionBudgetsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
