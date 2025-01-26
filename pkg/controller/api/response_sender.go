package api

import (
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/context"
	"github.com/krack8/lighthouse/pkg/log"
	"net/http"
)

type ResponseDTO struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
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

func logRequestedTaskController(resource string, taskName string) {
	log.Logger.Infow("from "+resource, "task", taskName)
}
func logErrMarshalTaskController(taskName string, err error) {
	log.Logger.Errorw("unable to marshal "+taskName+" Task input", "err", err.Error())
}
