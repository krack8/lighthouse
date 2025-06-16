package config

import (
	"github.com/joho/godotenv"
	"github.com/krack8/lighthouse/pkg/common/log"
	"os"
	"strings"
)

const DEVELOP = "DEVELOP"
const PRODUCTION = "PRODUCTION"
const True = "true"
const False = "false"

var RunMode string
var ServerPort = "8080"
var PageLimit = int64(10)
var isK8 = "False"
var KubeConfigFile = "dev-config.yaml"
var Auth = true
var ControllerGrpcTlsEnabled = true
var ControllerGrpcSkipTlsVerification = false
var TlsServerCustomCa = ""
var AgentGroup string
var ControllerGrpcServerHost string
var AgentSecretName string
var ResourceNamespace string
var HelmChartVersion string

func InitEnvironmentVariables(filenames ...string) {
	RunMode = os.Getenv("RUN_MODE")
	if RunMode == "" {
		RunMode = DEVELOP
	}
	log.Logger.Infow("RUN MODE:", "value", RunMode)
	if RunMode != PRODUCTION {
		err := godotenv.Load(filenames...)
		if err != nil {
			log.Logger.Errorw("Failed to Load environment file", "err", err.Error())
			os.Exit(1)
		}
	}
	if strings.ToLower(os.Getenv("AUTH_ENABLED")) == False {
		Auth = false
		log.Logger.Infow("Started with AUTH disabled", "[Auth-Enabled]", Auth)
	} else {
		log.Logger.Infow("Started with AUTH enabled", "[AUTH-Enabled]", Auth)
	}
	if strings.ToLower(os.Getenv("CONTROLLER_GRPC_TLS_ENABLED")) == True {
		log.Logger.Infow("Controller grpc server tls is enabled. Using tls config", "[Grpc-Server-Tls]", ControllerGrpcTlsEnabled)
	} else {
		ControllerGrpcTlsEnabled = false
		log.Logger.Infow("Controller grpc server tls is disabled. Skipping tls config", "[Grpc-Server-Tls]", ControllerGrpcTlsEnabled)
	}
	if ControllerGrpcTlsEnabled {
		if strings.ToLower(os.Getenv("CONTROLLER_GRPC_SKIP_TLS_VERIFICATION")) == True {
			ControllerGrpcSkipTlsVerification = true
			log.Logger.Infow("Controller grpc server is skipping tls verification.", "[TLS-Skip-Verify]", ControllerGrpcSkipTlsVerification)
		} else {
			log.Logger.Infow("Controller grpc server is verifying tls.", "[TLS-Skip-Verify]", ControllerGrpcSkipTlsVerification)
		}
	} else {
		log.Logger.Infow("Server transport security is disabled.", "[TLS]", false)
	}
	KubeConfigFile = os.Getenv("KUBE_CONFIG_FILE")
	isK8 = os.Getenv("IS_K8")
	TlsServerCustomCa = os.Getenv("TLS_SERVER_CUSTOM_CA")
	AgentGroup = os.Getenv("AGENT_GROUP")
	ControllerGrpcServerHost = os.Getenv("CONTROLLER_GRPC_SERVER_HOST")
	AgentSecretName = os.Getenv("AGENT_SECRET_NAME")
	ResourceNamespace = os.Getenv("RESOURCE_NAMESPACE")
	HelmChartVersion = os.Getenv("HELM_CHART_VERSION")
	if os.Getenv("SERVER_PORT") != "" {
		ServerPort = os.Getenv("SERVER_PORT")
	}
}

func IsK8() bool {
	if strings.ToLower(isK8) == True {
		return true
	}
	return false
}

func IsAuth() bool {
	return Auth
}

func IsControllerGrpcTlsEnabled() bool {
	return ControllerGrpcTlsEnabled
}

func IsControllerGrpcSkipTlsVerification() bool {
	return ControllerGrpcSkipTlsVerification
}
