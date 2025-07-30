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

type JobServiceInterface interface {
	GetJobList(c context.Context, p GetJobListInputParams) (interface{}, error)
	GetJobDetails(c context.Context, p GetJobInputParams) (interface{}, error)
	DeployJob(c context.Context, p DeployJobInputParams) (interface{}, error)
	DeleteJob(c context.Context, p DeleteJobInputParams) (interface{}, error)
}

type jobService struct{}

var js jobService

func JobService() *jobService {
	return &js
}

const (
	JobApiVersion = "batch/v1"
	JobKind       = "Job"
)

type Output struct {
	Result    []batchv1.Job
	Resource  string
	Remaining int64
	Total     int
}

type GetJobListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        Output
}

type GetJobInputParams struct {
	NamespaceName string
	JobName       string
	output        batchv1.Job
}

func (p *GetJobListInputParams) Find(jobClient v1.JobInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "job")
	filteredJobs := []batchv1.Job{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		jobList, err := jobClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get job list", "err", err.Error())
			return err
		}

		for _, job := range jobList.Items {
			if strings.Contains(job.Name, p.Search) {
				filteredJobs = append(filteredJobs, job)
			}
		}
		length = len(filteredJobs)
		nextPageToken = jobList.Continue
		if jobList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		jobList, err := jobClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get job list", "err", err.Error())
			return err
		}
		for _, job := range jobList.Items {
			if strings.Contains(job.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredJobs
	p.output.Total = len(filteredJobs)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetJobListInputParams) Process() error {
	log.Logger.Debugw("fetching job list of " + p.NamespaceName)
	jobClient := GetKubeClientSet().BatchV1().Jobs(p.NamespaceName)
	limit := config.PageLimit
	if p.Limit != "" {
		limit, _ = strconv.ParseInt(p.Limit, 10, 64)
	}
	listOptions := metav1.ListOptions{Limit: limit, Continue: p.Continue}
	if p.Labels != nil {
		labelSelector := metav1.LabelSelector{MatchLabels: p.Labels}
		listOptions.LabelSelector = labels.Set(labelSelector.MatchLabels).String()
	}
	var err error
	var jobList *batchv1.JobList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(jobClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get job list", "err", err.Error())
			return err
		}
		return nil
	} else {
		jobList, err = jobClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get job list", "err", err.Error())
			return err
		}

		jobList, err = jobClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get job list", "err", err.Error())
			return err
		}
		remaining := jobList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: jobList.Continue}
				res, err := jobClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get job list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = jobList.Items
		p.output.Total = len(jobList.Items)
		p.output.Resource = jobList.Continue
	}
	return nil
}

func (p *GetJobListInputParams) PostProcess(ctx context.Context) error {
	for i := 0; i < len(p.output.Result); i++ {
		p.output.Result[i].ManagedFields = nil
		p.output.Result[i].APIVersion = JobApiVersion
		p.output.Result[i].Kind = JobKind
	}
	return nil
}

func (svc *jobService) GetJobList(c context.Context, p GetJobListInputParams) (interface{}, error) {
	err := p.Process()
	if err != nil {
		return ErrorResponse(err)
	}
	_ = p.PostProcess(c)
	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

func (p *GetJobInputParams) Process() error {
	log.Logger.Debugw("fetching job details of ....", p.NamespaceName)
	jobs := GetKubeClientSet().BatchV1().Jobs(p.NamespaceName)
	output, err := jobs.Get(context.Background(), p.JobName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get job ", p.JobName, "err", err.Error())
		return err
	}
	p.output = *output
	p.output.ManagedFields = nil
	p.output.APIVersion = JobApiVersion
	p.output.Kind = JobKind
	return nil
}

func (svc *jobService) GetJobDetails(c context.Context, p GetJobInputParams) (interface{}, error) {
	err := p.Process()
	if err != nil {
		return ErrorResponse(err)
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployJobInputParams struct {
	Job    *batchv1.Job
	output *batchv1.Job
}

func (p *DeployJobInputParams) Process(c context.Context) error {
	jobClient := GetKubeClientSet().BatchV1().Jobs(p.Job.Namespace)
	_, err := jobClient.Get(context.Background(), p.Job.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating job in namespace "+p.Job.Namespace, "value", p.Job.Name)
		p.output, err = jobClient.Create(context.Background(), p.Job, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create job in namespace "+p.Job.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("job created")
	} else {
		log.Logger.Infow("job exist in namespace "+p.Job.Namespace, "value", p.Job.Name)
		log.Logger.Infow("Updating job in namespace "+p.Job.Namespace, "value", p.Job.Name)
		p.output, err = jobClient.Update(context.Background(), p.Job, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update job ", p.Job.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("job updated")
	}
	return nil
}

func (svc *jobService) DeployJob(c context.Context, p DeployJobInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return ErrorResponse(err)
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeleteJobInputParams struct {
	NamespaceName string
	JobName       string
}

func (p *DeleteJobInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting job of ....", p.NamespaceName)
	jobClient := GetKubeClientSet().BatchV1().Jobs(p.NamespaceName)
	_, err := jobClient.Get(context.Background(), p.JobName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get job ", p.JobName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = jobClient.Delete(context.Background(), p.JobName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete job ", p.JobName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *jobService) DeleteJob(c context.Context, p DeleteJobInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return ErrorResponse(err)
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
