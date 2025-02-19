package context

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	WsConnId   = "WsConnId"
	WsMsgId    = "WsMsgId"
	WsApiInput = "WsApiInput"
	UserKey    = "User"
)

func IsRequestFromWS(c *gin.Context) bool {
	if c.Keys == nil {
		return false
	}
	_, exists := c.Get(WsConnId)
	return exists
}

func AddWsConnIdToContext(c *gin.Context, val interface{}) {
	c.Set(WsConnId, val)
}

func GetWsConnId(c *gin.Context) string {
	if msgId, exists := c.Get(WsConnId); exists {
		return fmt.Sprintf("%v", msgId)
	}
	return ""
}

func AddRequestMsgIdToContext(c *gin.Context, val interface{}) {
	c.Set(WsMsgId, val)
}

func GetRequestMsgId(c *gin.Context) string {
	if msgId, exists := c.Get(WsMsgId); exists {
		return fmt.Sprintf("%v", msgId)
	}
	return ""
}

func AddInputToContext(c *gin.Context, input interface{}) {
	c.Set(WsApiInput, input)
}

func GetInputFromContext(c *gin.Context, in interface{}) error {
	// Binding input from context request if http
	if !IsRequestFromWS(c) {
		err := c.BindJSON(in)
		return err
	}

	input, exists := c.Get(WsApiInput)
	if !exists {
		return nil
	}

	// Binding input from context keys if websocket
	inputRaw, err := json.Marshal(input)
	if err != nil {
		return err
	}
	err = json.Unmarshal(inputRaw, in)
	if err != nil {
		return err
	}
	return nil
}
