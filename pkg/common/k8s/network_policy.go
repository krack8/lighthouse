package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/log"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	v1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	"strconv"
	"strings"
)

type NetworkPolicyServiceInterface interface {
	GetNetworkPolicyList(c context.Context, p GetNetworkPolicyListInputParams) (interface{}, error)
	GetNetworkPolicyDetails(c context.Context, p GetNetworkPolicyDetailsInputParams) (interface{}, error)
}

type networkPolicyService struct{}

var nps networkPolicyService

func NetworkPolicyService() *networkPolicyService {
	return &nps
}

type OutputNetworkPolicyList struct {
	Result    []networkingv1.NetworkPolicy
	Resource  string
	Remaining int64
	Total     int
}

type GetNetworkPolicyListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputNetworkPolicyList
}

func (p *GetNetworkPolicyListInputParams) Find(c context.Context, networkPolicyClient v1.NetworkPolicyInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "networkPolicy")
	filteredNetworkPolicies := []networkingv1.NetworkPolicy{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		networkPolicyList, err := networkPolicyClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get networkPolicy list", "err", err.Error())
			return err
		}

		for _, networkPolicy := range networkPolicyList.Items {
			if strings.Contains(networkPolicy.Name, p.Search) {
				filteredNetworkPolicies = append(filteredNetworkPolicies, networkPolicy)
			}
		}
		length = len(filteredNetworkPolicies)
		nextPageToken = networkPolicyList.Continue
		if networkPolicyList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		networkPolicyList, err := networkPolicyClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get networkPolicy list", "err", err.Error())
			return err
		}
		for _, networkPolicy := range networkPolicyList.Items {
			if strings.Contains(networkPolicy.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredNetworkPolicies
	p.output.Total = len(filteredNetworkPolicies)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetNetworkPolicyListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching networkPolicy list")
	networkPolicyClient := config.GetKubeClientSet().NetworkingV1().NetworkPolicies(p.NamespaceName)
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
	var networkPolicyList *networkingv1.NetworkPolicyList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, networkPolicyClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get networkPolicy list", "err", err.Error())
			return err
		}
		return nil
	} else {
		networkPolicyList, err = networkPolicyClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get networkPolicy list", "err", err.Error())
			return err
		}

		networkPolicyList, err = networkPolicyClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get networkPolicy list", "err", err.Error())
			return err
		}
		remaining := networkPolicyList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: networkPolicyList.Continue}
				res, err := networkPolicyClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get networkPolicy list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = networkPolicyList.Items
		p.output.Total = len(networkPolicyList.Items)
		p.output.Resource = networkPolicyList.Continue
	}
	return nil
}

func (svc *networkPolicyService) GetNetworkPolicyList(c context.Context, p GetNetworkPolicyListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetNetworkPolicyDetailsInputParams struct {
	NamespaceName     string
	NetworkPolicyName string
	output            networkingv1.NetworkPolicy
}

func (p *GetNetworkPolicyDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching networkPolicy details of ....", p.NamespaceName)
	networkPolicysClient := config.GetKubeClientSet().NetworkingV1().NetworkPolicies(p.NamespaceName)
	output, err := networkPolicysClient.Get(context.Background(), p.NetworkPolicyName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get networkPolicy ", p.NetworkPolicyName, "err", err.Error())
		return err
	}
	p.output = *output
	return nil
}

func (svc *networkPolicyService) GetNetworkPolicyDetails(c context.Context, p GetNetworkPolicyDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}
