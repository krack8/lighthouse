package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/log"
	"istio.io/client-go/pkg/apis/networking/v1beta1"
	_v1beta1 "istio.io/client-go/pkg/clientset/versioned/typed/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"strconv"
	"strings"
)

type IstioGatewayServiceInterface interface {
	GetIstioGatewayList(c context.Context, p GetIstioGatewayListInputParams) (interface{}, error)
	GetIstioGatewayDetails(c context.Context, p GetIstioGatewayDetailsInputParams) (interface{}, error)
	DeployIstioGateway(c context.Context, p DeployIstioGatewayInputParams) (interface{}, error)
	DeleteIstioGateway(c context.Context, p DeleteIstioGatewayInputParams) (interface{}, error)
}

type istioGatewayService struct{}

var ists istioGatewayService

func IstioGatewayService() *istioGatewayService {
	return &ists
}

const (
	IstioGatewayApiVersion = "networking.istio.io/v1"
	IstioGatewayKind       = "Gateway"
)

type OutputIstioGatewayList struct {
	Result    []*v1beta1.Gateway
	Resource  string
	Remaining int64
	Total     int
}

type GetIstioGatewayListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputIstioGatewayList
}

func (p *GetIstioGatewayListInputParams) Find(c context.Context, istioGatewayClient _v1beta1.GatewayInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "istioGateway")
	filteredIstoGateways := []*v1beta1.Gateway{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		istioGatewayList, err := istioGatewayClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get istioGateway list", "err", err.Error())
			return err
		}

		for _, istioGateway := range istioGatewayList.Items {
			if strings.Contains(istioGateway.Name, p.Search) {
				filteredIstoGateways = append(filteredIstoGateways, istioGateway)
			}
		}
		length = len(filteredIstoGateways)
		nextPageToken = istioGatewayList.Continue
		if istioGatewayList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		istioGatewayList, err := istioGatewayClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get istioGateway list", "err", err.Error())
			return err
		}
		for _, istioGateway := range istioGatewayList.Items {
			if strings.Contains(istioGateway.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredIstoGateways
	p.output.Total = len(filteredIstoGateways)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetIstioGatewayListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching istioGateway list")
	istioGatewayClient := GetNetworkingV1Beta1ClientSet().Gateways(p.NamespaceName)
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
	var istioGatewayList *v1beta1.GatewayList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, istioGatewayClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get istioGateway list", "err", err.Error())
			return err
		}
		return nil
	} else {
		istioGatewayList, err = istioGatewayClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get istioGateway list", "err", err.Error())
			return err
		}
		remaining := istioGatewayList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: istioGatewayList.Continue}
				res, err := istioGatewayClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get istioGateway list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = istioGatewayList.Items
		p.output.Total = len(istioGatewayList.Items)
		p.output.Resource = istioGatewayList.Continue
	}
	return nil
}

func (p *GetIstioGatewayListInputParams) PostProcess(ctx context.Context) error {
	for i := 0; i < len(p.output.Result); i++ {
		p.output.Result[i].ManagedFields = nil
		p.output.Result[i].TypeMeta.APIVersion = IstioGatewayApiVersion
		p.output.Result[i].TypeMeta.Kind = IstioGatewayKind
	}
	return nil
}

func (svc *istioGatewayService) GetIstioGatewayList(c context.Context, p GetIstioGatewayListInputParams) (interface{}, error) {
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

type GetIstioGatewayDetailsInputParams struct {
	NamespaceName    string
	IstioGatewayName string
	output           v1beta1.Gateway
}

func (p *GetIstioGatewayDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching istioGateway details of ....", p.NamespaceName)
	istioGatewaysClient := GetNetworkingV1Beta1ClientSet().Gateways(p.NamespaceName)
	output, err := istioGatewaysClient.Get(context.Background(), p.IstioGatewayName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get istioGateway ", p.IstioGatewayName, "err", err.Error())
		return err
	}
	p.output = *output
	p.output.ManagedFields = nil
	p.output.TypeMeta.APIVersion = IstioGatewayApiVersion
	p.output.TypeMeta.Kind = IstioGatewayKind
	return nil
}

func (svc *istioGatewayService) GetIstioGatewayDetails(c context.Context, p GetIstioGatewayDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployIstioGatewayInputParams struct {
	IstioGateway *v1beta1.Gateway
	output       *v1beta1.Gateway
}

func (p *DeployIstioGatewayInputParams) Process(c context.Context) error {
	IstioGatewayClient := GetNetworkingV1Beta1ClientSet().Gateways(p.IstioGateway.Namespace)
	_, err := IstioGatewayClient.Get(context.Background(), p.IstioGateway.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating istioGateway in namespace "+p.IstioGateway.Namespace, "value", p.IstioGateway.Name)
		p.output, err = IstioGatewayClient.Create(context.Background(), p.IstioGateway, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create istioGateway in namespace "+p.IstioGateway.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("istioGateway created")
	} else {
		log.Logger.Infow("istioGateway exist in namespace "+p.IstioGateway.Namespace, "value", p.IstioGateway.Name)
		log.Logger.Infow("Updating istioGateway in namespace "+p.IstioGateway.Namespace, "value", p.IstioGateway.Name)
		p.output, err = IstioGatewayClient.Update(context.Background(), p.IstioGateway, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update istioGateway ", p.IstioGateway.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("istioGateway updated")
	}
	return nil
}

func (svc *istioGatewayService) DeployIstioGateway(c context.Context, p DeployIstioGatewayInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeleteIstioGatewayInputParams struct {
	NamespaceName    string
	IstioGatewayName string
}

func (p *DeleteIstioGatewayInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting IstioGateway of ....", p.NamespaceName)
	IstioGatewayClient := GetNetworkingV1Beta1ClientSet().Gateways(p.NamespaceName)
	_, err := IstioGatewayClient.Get(context.Background(), p.IstioGatewayName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get IstioGateway ", p.IstioGatewayName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = IstioGatewayClient.Delete(context.Background(), p.IstioGatewayName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete IstioGateway ", p.IstioGatewayName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *istioGatewayService) DeleteIstioGateway(c context.Context, p DeleteIstioGatewayInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
