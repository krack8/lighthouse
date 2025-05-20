package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/log"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"strconv"
	"strings"
)

type SecretServiceInterface interface {
	GetSecretList(c context.Context, p GetSecretListInputParams) (interface{}, error)
	GetSecretDetails(c context.Context, p GetSecretDetailsInputParams) (interface{}, error)
	DeploySecret(c context.Context, p DeploySecretInputParams) (interface{}, error)
	DeleteSecret(c context.Context, p DeleteSecretInputParams) (interface{}, error)
}

type secretService struct{}

var ss secretService

func SecretService() *secretService {
	return &ss
}

const (
	SecretApiVersion = "v1"
	SecretKind       = "Secret"
)

type OutputSecretList struct {
	Result    []corev1.Secret
	Resource  string
	Remaining int64
	Total     int
}

type GetSecretListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputSecretList
}

func (p *GetSecretListInputParams) Find(c context.Context, secretClient v1.SecretInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "secret")
	filteredSecrets := []corev1.Secret{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		secretList, err := secretClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get secret list", "err", err.Error())
			return err
		}

		for _, secret := range secretList.Items {
			if strings.Contains(secret.Name, p.Search) {
				filteredSecrets = append(filteredSecrets, secret)
			}
		}
		length = len(filteredSecrets)
		nextPageToken = secretList.Continue
		if secretList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		secretList, err := secretClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get secret list", "err", err.Error())
			return err
		}
		for _, secret := range secretList.Items {
			if strings.Contains(secret.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredSecrets
	p.output.Total = len(filteredSecrets)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetSecretListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching secret list")
	secretClient := GetKubeClientSet().CoreV1().Secrets(p.NamespaceName)
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
	var secretList *corev1.SecretList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, secretClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get secret list", "err", err.Error())
			return err
		}
		return nil
	} else {
		secretList, err = secretClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get secret list", "err", err.Error())
			return err
		}

		secretList, err = secretClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get secret list", "err", err.Error())
			return err
		}
		remaining := secretList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: secretList.Continue}
				res, err := secretClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get secret list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = secretList.Items
		p.output.Total = len(secretList.Items)
		p.output.Resource = secretList.Continue
	}
	return nil
}

func (p *GetSecretListInputParams) PostProcess(ctx context.Context) error {
	for i := 0; i < len(p.output.Result); i++ {
		p.output.Result[i].ManagedFields = nil
		p.output.Result[i].APIVersion = SecretApiVersion
		p.output.Result[i].Kind = SecretKind
	}
	return nil
}

func (svc *secretService) GetSecretList(c context.Context, p GetSecretListInputParams) (interface{}, error) {
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

type GetSecretDetailsInputParams struct {
	NamespaceName string
	SecretName    string
	output        corev1.Secret
}

func (p *GetSecretDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching secret details of ....", p.NamespaceName)
	secretsClient := GetKubeClientSet().CoreV1().Secrets(p.NamespaceName)
	output, err := secretsClient.Get(context.Background(), p.SecretName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get secret ", p.SecretName, "err", err.Error())
		return err
	}
	p.output = *output
	p.output.ManagedFields = nil
	p.output.APIVersion = SecretApiVersion
	p.output.Kind = SecretKind
	return nil
}

func (svc *secretService) GetSecretDetails(c context.Context, p GetSecretDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeploySecretInputParams struct {
	Secret *corev1.Secret
	output *corev1.Secret
}

func (p *DeploySecretInputParams) PostProcess(c context.Context) error {
	p.output.ManagedFields = nil
	return nil
}

func (p *DeploySecretInputParams) Process(c context.Context) error {
	secretClient := GetKubeClientSet().CoreV1().Secrets(p.Secret.Namespace)
	_, err := secretClient.Get(context.Background(), p.Secret.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating secret in namespace "+p.Secret.Namespace, "value", p.Secret.Name)
		p.output, err = secretClient.Create(context.Background(), p.Secret, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create secret in namespace "+p.Secret.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("secret created")
	} else {
		log.Logger.Infow("Secret exist in namespace "+p.Secret.Namespace, "value", p.Secret.Name)
		log.Logger.Infow("Updating secret in namespace "+p.Secret.Namespace, "value", p.Secret.Name)
		p.output, err = secretClient.Update(context.Background(), p.Secret, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update secret ", p.Secret.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("secret updated")
	}
	return nil
}

func (svc *secretService) DeploySecret(c context.Context, p DeploySecretInputParams) (interface{}, error) {
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

type DeleteSecretInputParams struct {
	NamespaceName string
	SecretName    string
}

func (p *DeleteSecretInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting secret of ....", p.NamespaceName)
	secretClient := GetKubeClientSet().CoreV1().Secrets(p.NamespaceName)
	_, err := secretClient.Get(context.Background(), p.SecretName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get secret ", p.SecretName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = secretClient.Delete(context.Background(), p.SecretName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete secret ", p.SecretName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *secretService) DeleteSecret(c context.Context, p DeleteSecretInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
