package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/log"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	v1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	"strconv"
	"strings"
)

type CronJobServiceInterface interface {
	GetCronJobList(c context.Context, p GetCronJobListInputParams) (interface{}, error)
	GetCronJobDetails(c context.Context, p GetCronJobInputParams) (interface{}, error)
	DeployCronJob(c context.Context, p DeployCronJobInputParams) (interface{}, error)
	DeleteCronJob(c context.Context, p DeleteCronJobInputParams) (interface{}, error)
}

type cronJobService struct{}

var cjs cronJobService

func CronJobService() *cronJobService {
	return &cjs
}

type OutputCronJobList struct {
	Result    []batchv1.CronJob
	Resource  string
	Remaining int64
	Total     int
}

type GetCronJobListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputCronJobList
}

type GetCronJobInputParams struct {
	NamespaceName string
	CronJobName   string
	output        batchv1.CronJob
}

func (p *GetCronJobListInputParams) Find(jobClient v1.CronJobInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "cron job")
	filteredCronJobs := []batchv1.CronJob{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		cronJobList, err := jobClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get  cronjob list", "err", err.Error())
			return err
		}

		for _, job := range cronJobList.Items {
			if strings.Contains(job.Name, p.Search) {
				filteredCronJobs = append(filteredCronJobs, job)
			}
		}
		length = len(filteredCronJobs)
		nextPageToken = cronJobList.Continue
		if cronJobList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		cronJobList, err := jobClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get cronjob list", "err", err.Error())
			return err
		}
		for _, cronJob := range cronJobList.Items {
			if strings.Contains(cronJob.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredCronJobs
	p.output.Total = len(filteredCronJobs)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetCronJobListInputParams) Process() error {
	log.Logger.Debugw("fetching cronjob list")
	cronJobClient := config.GetKubeClientSet().BatchV1().CronJobs(p.NamespaceName)
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
	var cronJobList *batchv1.CronJobList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(cronJobClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get cronjob list", "err", err.Error())
			return err
		}
		return nil
	} else {
		cronJobList, err = cronJobClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get cronjob list", "err", err.Error())
			return err
		}

		cronJobList, err = cronJobClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get cronjob list", "err", err.Error())
			return err
		}
		remaining := cronJobList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: cronJobList.Continue}
				res, err := cronJobClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get cronjob list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = cronJobList.Items
		p.output.Total = len(cronJobList.Items)
		p.output.Resource = cronJobList.Continue
	}
	return nil
}

func (svc *cronJobService) GetCronJobList(c context.Context, p GetCronJobListInputParams) (interface{}, error) {
	err := p.Process()
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

func (p *GetCronJobInputParams) Process() error {
	log.Logger.Debugw("fetching cronjob details of ....", p.NamespaceName)
	cronJobClient := config.GetKubeClientSet().BatchV1().CronJobs(p.NamespaceName)
	output, err := cronJobClient.Get(context.Background(), p.CronJobName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get cronjob ", p.CronJobName, "err", err.Error())
		return err
	}
	p.output = *output
	return nil
}

func (svc *cronJobService) GetCronJobDetails(c context.Context, p GetCronJobInputParams) (interface{}, error) {
	err := p.Process()
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployCronJobInputParams struct {
	CronJob *batchv1.CronJob
	output  *batchv1.CronJob
}

func (p *DeployCronJobInputParams) Process(c context.Context) error {
	cronJobClient := config.GetKubeClientSet().BatchV1().CronJobs(p.CronJob.Namespace)
	_, err := cronJobClient.Get(context.Background(), p.CronJob.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating cron job in namespace "+p.CronJob.Namespace, "value", p.CronJob.Name)
		p.output, err = cronJobClient.Create(context.Background(), p.CronJob, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create cron job in namespace "+p.CronJob.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("cron job created")
	} else {
		log.Logger.Infow("cron job exist in namespace "+p.CronJob.Namespace, "value", p.CronJob.Name)
		log.Logger.Infow("Updating cron job in namespace "+p.CronJob.Namespace, "value", p.CronJob.Name)
		p.output, err = cronJobClient.Update(context.Background(), p.CronJob, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update cron job ", p.CronJob.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("cron job updated")
	}
	return nil
}

func (svc *cronJobService) DeployCronJob(c context.Context, p DeployCronJobInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeleteCronJobInputParams struct {
	NamespaceName string
	CronJobName   string
}

func (p *DeleteCronJobInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting cronJob of ....", p.NamespaceName)
	cronJobClient := config.GetKubeClientSet().BatchV1().CronJobs(p.NamespaceName)
	_, err := cronJobClient.Get(context.Background(), p.CronJobName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get cronJob ", p.CronJobName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = cronJobClient.Delete(context.Background(), p.CronJobName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete cronJob ", p.CronJobName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *cronJobService) DeleteCronJob(c context.Context, p DeleteCronJobInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
