package dto

import corev1 "k8s.io/api/core/v1"

type TaintList struct {
	Taint []corev1.Taint `json:"taint"`
}

type UnTaintKeys struct {
	Keys []string `json:"keys"`
}
