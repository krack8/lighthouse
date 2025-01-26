package k8s

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"

	_context "context"
	"github.com/krack8/lighthouse/pkg/context"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type ResponseDTO struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

var clientset *kubernetes.Clientset
var (
	ErrorUnstructuredNil   = errors.New("unstructured CustomResource is nil")
	ErrorFieldEmptyStr     = " field empty"
	ErrorFieldMismatch     = " field mismatched"
	ErrorNamespaceMismatch = "namespace mismatched"
)

const (
	RUNNING = "Running"
	FAILED  = "Failed"
	PENDING = "Pending"
)

func InitClient() {
	// Load kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to load kubeconfig: %v", err))
	}

	// Create Kubernetes client
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(fmt.Sprintf("Failed to create Kubernetes client: %v", err))
	}
}

func GetClientset() *kubernetes.Clientset {
	return clientset
}

func ListNamespaces(clientset kubernetes.Interface) ([]string, error) {
	namespaces, err := clientset.CoreV1().Namespaces().List(_context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list namespaces: %w", err)
	}

	var namespaceNames []string
	for _, ns := range namespaces.Items {
		namespaceNames = append(namespaceNames, ns.Name)
	}
	return namespaceNames, nil
}

var nilResponse ResponseDTO = ResponseDTO{}

var (
	UserAuthenticatedMsg = "User authenticated"
	UserLoggedOutMsg     = "Logged out from lighthouse"
	ErrNamespaceEmpty    = "namespace field empty"
	ClearRedis           = "role cleared"
)

var httpStatusMap = map[string]int{
	"success": http.StatusOK,
	"error":   http.StatusBadRequest,
}

func NilResponse() ResponseDTO {
	return nilResponse
}

func ErrorResponse(err error) (ResponseDTO, error) {
	return ResponseDTO{
		Status: "error",
		Msg:    err.Error(),
	}, nil
}

func SuccessResponse(data interface{}) (ResponseDTO, error) {
	return ResponseDTO{
		Status: "success",
		Data:   data,
	}, nil
}

func executeSendResponse(c *gin.Context, data interface{}, httpStatus int) {
	sendHttpResponse(c, data, httpStatus)
}

func SendResponse(c *gin.Context, response ResponseDTO) {
	executeSendResponse(c, response, httpStatusMap[response.Status])
}

func SendErrorResponse(c *gin.Context, msg string) {
	data := gin.H{
		"status": http.StatusBadRequest,
		"msg":    msg,
	}
	if context.IsRequestFromWS(c) {
		data["msgId"] = context.GetRequestMsgId(c)
	}
	executeSendResponse(c, data, http.StatusBadRequest)
}

func sendHttpResponse(c *gin.Context, data interface{}, httpStatus int) {
	c.JSON(httpStatus, data)
}

func SendForbiddenResponse(c *gin.Context, msg string) {
	data := gin.H{
		"status": http.StatusForbidden,
		"msg":    msg,
	}
	if context.IsRequestFromWS(c) {
		data["msgId"] = context.GetRequestMsgId(c)
	}
	executeSendResponse(c, data, http.StatusForbidden)
}

//func SendSuccessResponse(c *gin.Context, response ResponseDTO) {
//	response.Status = "success"
//	data := gin.H{
//		"status": http.StatusOK,
//		"data":   response,
//	}
//	if context.IsRequestFromWS(c) {
//		data["msgId"] = context.GetRequestMsgId(c)
//	}
//	executeSendResponse(c, data, http.StatusOK)
//}

func SendSuccessResponse(c *gin.Context, response ResponseDTO) {
	response.Status = "success"
	executeSendResponse(c, response, httpStatusMap[response.Status])
}
