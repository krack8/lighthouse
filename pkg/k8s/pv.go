package k8s

import (
	"context"
	cfg "github.com/krack8/lighthouse/pkg/config"
	"github.com/krack8/lighthouse/pkg/log"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"strconv"
	"strings"
)

type PvServiceInterface interface {
	GetPvList(c context.Context, p GetPvListInputParams) (interface{}, error)
	GetPvDetails(c context.Context, p GetPvDetailsInputParams) (interface{}, error)
	DeployPv(c context.Context, p DeployPvInputParams) (interface{}, error)
	DeletePv(c context.Context, p DeletePvInputParams) (interface{}, error)
}

type pvService struct{}

var pvs pvService

func PvService() *pvService {
	return &pvs
}

type OutputPvList struct {
	Result    []corev1.PersistentVolume
	Resource  string
	Remaining int64
	Total     int
}

type GetPvListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputPvList
}

func (p *GetPvListInputParams) PostProcess(c context.Context) error {
	for idx, _ := range p.output.Result {
		p.output.Result[idx].ManagedFields = nil
	}
	return nil
}

func (p *GetPvListInputParams) Find(c context.Context, persistentVolumeClient v1.PersistentVolumeInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "pv")
	filteredPV := []corev1.PersistentVolume{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		pvList, err := persistentVolumeClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get pv list", "err", err.Error())
			return err
		}

		for _, pv := range pvList.Items {
			if strings.Contains(pv.Name, p.Search) {
				filteredPV = append(filteredPV, pv)
			}
		}
		length = len(filteredPV)
		nextPageToken = pvList.Continue
		if pvList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		pvList, err := persistentVolumeClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get pv list", "err", err.Error())
			return err
		}
		for _, pv := range pvList.Items {
			if strings.Contains(pv.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredPV
	p.output.Total = len(filteredPV)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetPvListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching pv list")
	pvClient := cfg.GetKubeClientSet().CoreV1().PersistentVolumes()
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
	var pvList *corev1.PersistentVolumeList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, pvClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get pv list", "err", err.Error())
			return err
		}
		return nil
	} else {
		pvList, err = pvClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get pv list", "err", err.Error())
			return err
		}

		pvList, err = pvClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get pv list", "err", err.Error())
			return err
		}
		remaining := pvList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: pvList.Continue}
				res, err := pvClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get pv list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = pvList.Items
		p.output.Total = len(pvList.Items)
		p.output.Resource = pvList.Continue
	}
	return nil
}

func (svc *pvService) GetPvList(c context.Context, p GetPvListInputParams) (interface{}, error) {
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

type GetPvDetailsInputParams struct {
	PvName string
	output corev1.PersistentVolume
}

func (p *GetPvDetailsInputParams) PostProcess(c context.Context) error {
	p.output.ManagedFields = nil
	return nil
}

func (p *GetPvDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching pv details of ....", p.PvName)
	pvsClient := cfg.GetKubeClientSet().CoreV1().PersistentVolumes()
	output, err := pvsClient.Get(context.Background(), p.PvName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get pv ", p.PvName, "err", err.Error())
		return err
	}
	p.output = *output
	return nil
}

func (svc *pvService) GetPvDetails(c context.Context, p GetPvDetailsInputParams) (interface{}, error) {
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

type DeployPvInputParams struct {
	Pv     *corev1.PersistentVolume
	output *corev1.PersistentVolume
}

func (p *DeployPvInputParams) PostProcess(c context.Context) error {
	p.output.ManagedFields = nil
	return nil
}

func (p *DeployPvInputParams) Process(c context.Context) error {
	pvClient := cfg.GetKubeClientSet().CoreV1().PersistentVolumes()
	_, err := pvClient.Get(context.Background(), p.Pv.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating pv ", "value", p.Pv.Name)
		p.output, err = pvClient.Create(context.Background(), p.Pv, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create pv ", "err", err.Error())
			return err
		}
		log.Logger.Infow("pv created")
	} else {
		log.Logger.Infow("Pv exist ", "value", p.Pv.Name)
		log.Logger.Infow("Updating pv ", "value", p.Pv.Name)
		p.output, err = pvClient.Update(context.Background(), p.Pv, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update pv ", p.Pv.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("pv updated")
	}
	return nil
}

func (svc *pvService) DeployPv(c context.Context, p DeployPvInputParams) (interface{}, error) {
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

type DeletePvInputParams struct {
	NamespaceName string
	PvName        string
}

func (p *DeletePvInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting pv of ....", p.NamespaceName)
	pvClient := cfg.GetKubeClientSet().CoreV1().PersistentVolumes()
	_, err := pvClient.Get(context.Background(), p.PvName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get pv ", p.PvName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = pvClient.Delete(context.Background(), p.PvName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete pv ", p.PvName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *pvService) DeletePv(c context.Context, p DeletePvInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
