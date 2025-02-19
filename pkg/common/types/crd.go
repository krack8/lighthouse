package types

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CrdForList struct {
	Kind              string      `json:"kind"`
	Name              string      `json:"name"`
	NamePlural        string      `json:"name_plural"`
	Group             string      `json:"group"`
	Scope             interface{} `json:"scope"`
	Version           []string    `json:"version"`
	CreationTimestamp v1.Time     `json:"creationTimestamp"`
}
