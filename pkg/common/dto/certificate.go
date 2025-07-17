package dto

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type KeyUsage string
type PrivateKeyRotationPolicy string
type PrivateKeyEncoding string
type PrivateKeyAlgorithm string
type ConditionStatus string
type CertificateConditionType string

type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              CertificateSpec   `json:"spec"`
	Status            CertificateStatus `json:"status"`
}

// Full specification (https://github.com/jetstack/cert-manager/blob/master/pkg/client/clientset/versioned/typed/certmanager/v1/certificate.go).
type CertificateSpec struct {
	Subject               *X509Subject           `json:"subject,omitempty"`
	CommonName            string                 `json:"commonName,omitempty"`
	Duration              *metav1.Duration       `json:"duration,omitempty"`
	RenewBefore           *metav1.Duration       `json:"renewBefore,omitempty"`
	DNSNames              []string               `json:"dnsNames,omitempty"`
	IPAddresses           []string               `json:"ipAddresses,omitempty"`
	URIs                  []string               `json:"uris,omitempty"`
	EmailAddresses        []string               `json:"emailAddresses,omitempty"`
	SecretName            string                 `json:"secretName"`
	Keystores             *CertificateKeystores  `json:"keystores,omitempty"`
	IssuerRef             ObjectReference        `json:"issuerRef"`
	IsCA                  bool                   `json:"isCA,omitempty"`
	Usages                []KeyUsage             `json:"usages,omitempty"`
	PrivateKey            *CertificatePrivateKey `json:"privateKey,omitempty"`
	EncodeUsagesInRequest *bool                  `json:"encodeUsagesInRequest,omitempty"`
}

type CertificateStatus struct {
	Conditions               []CertificateCondition `json:"conditions,omitempty"`
	LastFailureTime          *metav1.Time           `json:"lastFailureTime,omitempty"`
	NotBefore                *metav1.Time           `json:"notBefore,omitempty"`
	NotAfter                 *metav1.Time           `json:"notAfter,omitempty"`
	RenewalTime              *metav1.Time           `json:"renewalTime,omitempty"`
	Revision                 *int                   `json:"revision,omitempty"`
	NextPrivateKeySecretName *string                `json:"nextPrivateKeySecretName,omitempty"`
}

type CertificateCondition struct {
	Type               CertificateConditionType `json:"type"`
	Status             ConditionStatus          `json:"status"`
	LastTransitionTime *metav1.Time             `json:"lastTransitionTime,omitempty"`
	Reason             string                   `json:"reason,omitempty"`
	Message            string                   `json:"message,omitempty"`
}

type CertificatePrivateKey struct {
	RotationPolicy PrivateKeyRotationPolicy `json:"rotationPolicy,omitempty"`
	Encoding       PrivateKeyEncoding       `json:"encoding,omitempty"`
	Algorithm      PrivateKeyAlgorithm      `json:"algorithm,omitempty"`
	Size           int                      `json:"size,omitempty"` // Validated by webhook. Be mindful of adding OpenAPI validation- see https://github.com/jetstack/cert-manager/issues/3644
}

type ObjectReference struct {
	Name string `json:"name"`
	Kind string `json:"kind,omitempty"`
}

type X509Subject struct {
	Organizations       []string `json:"organizations,omitempty"`
	Countries           []string `json:"countries,omitempty"`
	OrganizationalUnits []string `json:"organizationalUnits,omitempty"`
	Localities          []string `json:"localities,omitempty"`
	Provinces           []string `json:"provinces,omitempty"`
	StreetAddresses     []string `json:"streetAddresses,omitempty"`
	PostalCodes         []string `json:"postalCodes,omitempty"`
	SerialNumber        string   `json:"serialNumber,omitempty"`
}

type CertificateKeystores struct {
	JKS    *JKSKeystore    `json:"jks,omitempty"`
	PKCS12 *PKCS12Keystore `json:"pkcs12,omitempty"`
}

type JKSKeystore struct {
	Create            bool              `json:"create"`
	PasswordSecretRef SecretKeySelector `json:"passwordSecretRef"`
}

type PKCS12Keystore struct {
	Create            bool              `json:"create"`
	PasswordSecretRef SecretKeySelector `json:"passwordSecretRef"`
}

type SecretKeySelector struct {
	LocalObjectReference `json:",inline"`
	Key                  string `json:"key,omitempty"`
}

type LocalObjectReference struct {
	Name string `json:"name"`
}

func (config *Certificate) GenerateUnstructured() *unstructured.Unstructured {
	if config == nil {
		return nil
	}
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       config.Kind,
			"apiVersion": config.APIVersion,
			"metadata": map[string]interface{}{
				"name":      config.Name,
				"namespace": config.Namespace,
				"labels":    config.Labels,
			},
			"spec": map[string]interface{}{
				"dnsNames":   config.Spec.DNSNames,
				"secretName": config.Spec.SecretName,
				"issuerRef": map[string]string{
					"kind": config.Spec.IssuerRef.Kind,
					"name": config.Spec.IssuerRef.Name,
				},
			},
		},
	}
}
