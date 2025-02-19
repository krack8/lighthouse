package config

import (
	"flag"
	"github.com/krack8/lighthouse/pkg/common/log"
	snapshotV1 "github.com/kubernetes-csi/external-snapshotter/client/v6/clientset/versioned/typed/volumesnapshot/v1"
	networkingv1beta1 "istio.io/client-go/pkg/clientset/versioned/typed/networking/v1beta1"
	apiExtension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
	"path/filepath"
)

var clientSet *kubernetes.Clientset
var dynamicClientSet *dynamic.DynamicClient
var snapshotV1ClientSet *snapshotV1.SnapshotV1Client
var apiExtensionCientSet *apiExtension.Clientset
var metricsClientSet *metrics.Clientset
var networkingV1beta1ClientSet *networkingv1beta1.NetworkingV1beta1Client

func InitiateKubeClientSet() {
	var kubeConfig *string
	var restConfig *rest.Config
	var err error

	if IsK8() {
		restConfig, err = clientcmd.BuildConfigFromFlags("", "")
	} else {
		if home := homedir.HomeDir(); home != "" {
			kubeConfig = flag.String("kubeconfig", filepath.Join(home, ".kube", KubeConfigFile), "(optional) absolute path to the kubeconfig file")
			restConfig, err = clientcmd.BuildConfigFromFlags("", *kubeConfig)
			log.Logger.Info(filepath.Join(home, ".kube", KubeConfigFile))

		} else {
			restConfig, err = clientcmd.BuildConfigFromFlags("", "")
		}
	}

	if err != nil {
		log.Logger.Errorw(err.Error())
	}

	clientSet, err = kubernetes.NewForConfig(restConfig)

	if err != nil {
		log.Logger.Errorw(err.Error())
	}

	dynamicClientSet, err = dynamic.NewForConfig(restConfig)
	if err != nil {
		panic(err)
	}

	snapshotV1ClientSet, err = snapshotV1.NewForConfig(restConfig)
	if err != nil {
		panic(err)
	}

	apiExtensionCientSet, err = apiExtension.NewForConfig(restConfig)
	if err != nil {
		panic(err)
	}

	metricsClientSet, err = metrics.NewForConfig(restConfig)
	if err != nil {
		panic(err)
	}

	networkingV1beta1ClientSet, err = networkingv1beta1.NewForConfig(restConfig)
	if err != nil {
		panic(err)
	}

}

func GetKubeClientSet() *kubernetes.Clientset {
	return clientSet
}

func GetDynamicClientSet() *dynamic.DynamicClient {
	return dynamicClientSet
}

func GetSnapshotV1ClientSet() *snapshotV1.SnapshotV1Client {
	return snapshotV1ClientSet
}

func GetApiExtensionClientSet() *apiExtension.Clientset {
	return apiExtensionCientSet
}

func GetMetricsClientSet() *metrics.Clientset {
	return metricsClientSet
}

func GetNetworkingV1Beta1ClientSet() *networkingv1beta1.NetworkingV1beta1Client {
	return networkingV1beta1ClientSet
}

const (
	MetricsAbsPath = "apis/metrics.k8s.io/v1beta1/namespaces/"
)

var (
	CertificateSGVR = schema.GroupVersionResource{
		Group:    "cert-manager.io",
		Version:  "v1",
		Resource: "certificates",
	}

	VolumeSnapshotSGVR = schema.GroupVersionResource{
		Group:    "snapshot.storage.k8s.io",
		Version:  "v1",
		Resource: "volumesnapshots",
	}

	VolumeSnapshotContentSGVR = schema.GroupVersionResource{
		Group:    "snapshot.storage.k8s.io",
		Version:  "v1",
		Resource: "volumesnapshotcontents",
	}
)
