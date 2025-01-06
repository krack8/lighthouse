package types

import "errors"

type CreateNamespaceDto struct {
	Name        string            `json:"name"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

func (payload CreateNamespaceDto) Validate() error {
	if payload.Name == "" {
		return errors.New("namespace name is empty")
	}
	return nil
}

type UpdateNamespaceDto struct {
	Labels map[string]string `json:"labels"`
}
