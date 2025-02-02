package k8s

import (
	"context"
	cfg "github.com/krack8/lighthouse/pkg/config"
	"github.com/krack8/lighthouse/pkg/dto"
	"github.com/krack8/lighthouse/pkg/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	"strconv"
	"strings"
)

type CertificateServiceInterface interface {
	GetCertificateList(c context.Context, p GetCertificateListInputParams) (interface{}, error)
	GetCertificateDetails(c context.Context, p GetCertificateDetailsInputParams) (interface{}, error)
	DeployCertificate(c context.Context, p DeployCertificateInputParams) (interface{}, error)
	DeleteCertificate(c context.Context, p DeleteCertificateInputParams) (interface{}, error)
}

type certificateService struct{}

var certs certificateService

func CertificateService() *certificateService {
	return &certs
}

type OutputCertificateList struct {
	Result    []*dto.Certificate
	Resource  string
	Remaining int64
	Total     int
}

type GetCertificateListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputCertificateList
}

func (p *GetCertificateListInputParams) Find(c context.Context, certificateClient dynamic.ResourceInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "certificate")
	filteredCertificates := []*dto.Certificate{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		unstructuredCertificateList, err := certificateClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get certificate list", "err", err.Error())
			return err
		} else {
			log.Logger.Info("get certificate list")
			if unstructuredCertificateList.Items != nil {
				for _, unstructured := range unstructuredCertificateList.Items {
					var certificate dto.Certificate
					err := runtime.DefaultUnstructuredConverter.
						FromUnstructured(unstructured.Object, &certificate)
					if err != nil {
						log.Logger.Errorw("Get certificate list structured conversion", "err", err.Error())
						return err
					}
					if strings.Contains(certificate.Name, p.Search) {
						filteredCertificates = append(filteredCertificates, &certificate)
					}
				}
			}
		}
		length = len(filteredCertificates)
		nextPageToken = unstructuredCertificateList.GetContinue()
		if nextPageToken == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		unstructuredCertificateList, err := certificateClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get certificate list", "err", err.Error())
			return err
		}
		if unstructuredCertificateList.Items != nil {
			for _, unstructured := range unstructuredCertificateList.Items {
				var certificate dto.Certificate
				err := runtime.DefaultUnstructuredConverter.
					FromUnstructured(unstructured.Object, &certificate)
				if err != nil {
					log.Logger.Errorw("Get certificate list structured conversion", "err", err.Error())
					return err
				}
				if strings.Contains(certificate.Name, p.Search) {
					remaining = remaining + 1
				}
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredCertificates
	p.output.Total = len(filteredCertificates)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetCertificateListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching certificate list")
	var certificateList = []*dto.Certificate{}
	var err error

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
	certificateClient := cfg.GetDynamicClientSet().Resource(cfg.CertificateSGVR).Namespace(p.NamespaceName)
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, certificateClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get pod list", "err", err.Error())
			return err
		}
		return nil
	} else {
		unstructuredCertificateList, err := certificateClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get certificate list", "err", err.Error())
			return err
		}
		if unstructuredCertificateList.Items != nil {
			for _, unstructured := range unstructuredCertificateList.Items {
				var certificate dto.Certificate
				err := runtime.DefaultUnstructuredConverter.
					FromUnstructured(unstructured.Object, &certificate)
				if err != nil {
					log.Logger.Errorw("Get certificate list structured conversion", "err", err.Error())
					return err
				}
				certificateList = append(certificateList, &certificate)
			}
		}
		remaining := unstructuredCertificateList.GetRemainingItemCount()
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: unstructuredCertificateList.GetContinue()}
				res, err := certificateClient.List(context.Background(), listOptions)
				if err != nil {
					log.Logger.Errorw("failed to get certificate remaining count", "err", err.Error())
					return err
				}
				p.output.Remaining = int64(len(res.Items))
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Resource = unstructuredCertificateList.GetContinue()
		p.output.Result = certificateList
		p.output.Total = len(certificateList)
	}
	return nil
}

func (svc *certificateService) GetCertificateList(c context.Context, p GetCertificateListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}
	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetCertificateDetailsInputParams struct {
	NamespaceName   string
	CertificateName string
	output          dto.Certificate
}

func (p *GetCertificateDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching certificate details of ....", p.NamespaceName)
	var certificate dto.Certificate

	certificatesClient := cfg.GetDynamicClientSet().Resource(cfg.CertificateSGVR)
	unstructuredCertificate, err := certificatesClient.Namespace(p.NamespaceName).Get(context.Background(), p.CertificateName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get certificate ", p.CertificateName, "err", err.Error())
		return err
	} else {
		err := runtime.DefaultUnstructuredConverter.
			FromUnstructured(unstructuredCertificate.Object, &certificate)
		if err != nil {
			log.Logger.Errorw("Get certificate details", "err", err.Error())
			return err
		}
	}
	p.output = certificate
	return nil
}

func (svc *certificateService) GetCertificateDetails(c context.Context, p GetCertificateDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployCertificateInputParams struct {
	Certificate *dto.Certificate
}

func (p *DeployCertificateInputParams) Process(c context.Context) error {
	certificatesClient := cfg.GetDynamicClientSet().Resource(cfg.CertificateSGVR)
	unstructuredCertificate := p.Certificate.GenerateUnstructured()
	if unstructuredCertificate == nil {
		log.Logger.Errorw("unstructured Certificate is nil")
		return ErrorUnstructuredNil
	}
	returnCertificate, err := certificatesClient.Namespace(p.Certificate.Namespace).Get(context.Background(), p.Certificate.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating certificate in namespace "+p.Certificate.Namespace, "value", p.Certificate.Name)
		_, err := certificatesClient.Namespace(p.Certificate.Namespace).Create(context.Background(), unstructuredCertificate, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create certificate in namespace "+p.Certificate.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("certificate created")
	} else {
		log.Logger.Infow("certificate exist in namespace "+p.Certificate.Namespace, "value", p.Certificate.Name)
		log.Logger.Infow("Updating certificate in namespace "+p.Certificate.Namespace, "value", p.Certificate.Name)
		unstructuredCertificate.SetResourceVersion(returnCertificate.GetResourceVersion())
		_, err = certificatesClient.Namespace(p.Certificate.Namespace).Update(context.Background(), unstructuredCertificate, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update certificate ", p.Certificate.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("certificate updated")
	}
	return nil
}

func (svc *certificateService) DeployCertificate(c context.Context, p DeployCertificateInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}

type DeleteCertificateInputParams struct {
	NamespaceName   string
	CertificateName string
}

func (p *DeleteCertificateInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting Certificate of ....", p.NamespaceName)
	certificatesClient := cfg.GetDynamicClientSet().Resource(cfg.CertificateSGVR)
	_, err := certificatesClient.Namespace(p.NamespaceName).Get(context.Background(), p.CertificateName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get Certificate ", p.CertificateName, "err", err.Error())
		return err
	}
	err = certificatesClient.Namespace(p.NamespaceName).Delete(context.Background(), p.CertificateName, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to delete Certificate ", p.CertificateName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *certificateService) DeleteCertificate(c context.Context, p DeleteCertificateInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
