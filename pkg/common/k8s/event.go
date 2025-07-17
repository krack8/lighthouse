package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/log"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
)

type EventServiceInterface interface {
	GetEventList(c context.Context, p GetEventListInputParams) (interface{}, error)
	GetEventDetails(c context.Context, p GetEventDetailsInputParams) (interface{}, error)
}

type eventService struct{}

var es eventService

func EventService() *eventService {
	return &es
}

type GetEventListInputParams struct {
	NamespaceName      string
	Search             string
	InvolvedObjectName string
	Labels             map[string]string
	output             []corev1.Event
}

func (p *GetEventListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching event list")
	eventClient := GetKubeClientSet().CoreV1().Events(p.NamespaceName)
	listOptions := metav1.ListOptions{}
	if p.Labels != nil {
		labelSelector := metav1.LabelSelector{MatchLabels: p.Labels}
		listOptions.LabelSelector = labels.Set(labelSelector.MatchLabels).String()
	}
	if p.InvolvedObjectName != "" {
		listOptions.FieldSelector = "involvedObject.name=" + p.InvolvedObjectName
		log.Logger.Infow("Event FieldSelector", "value", p.InvolvedObjectName)
	}
	if p.Search != "" {
		listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
	}
	eventList, err := eventClient.List(context.Background(), listOptions)
	if err != nil {
		log.Logger.Errorw("Failed to get event list", "err", err.Error())
		return err
	}
	p.output = eventList.Items
	return nil
}

func (svc *eventService) GetEventList(c context.Context, p GetEventListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetEventDetailsInputParams struct {
	NamespaceName string
	EventName     string
	output        corev1.Event
}

func (p *GetEventDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching event details of ....", p.NamespaceName)
	eventsClient := GetKubeClientSet().CoreV1().Events(p.NamespaceName)
	output, err := eventsClient.Get(context.Background(), p.EventName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get event ", p.EventName, "err", err.Error())
		return err
	}
	p.output = *output
	return nil
}

func (svc *eventService) GetEventDetails(c context.Context, p GetEventDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}
