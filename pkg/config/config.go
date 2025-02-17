package config

import (
	"github.com/joho/godotenv"
	"github.com/krack8/lighthouse/pkg/log"
	"os"
)

const DEVELOP = "DEVELOP"
const PRODUCTION = "PRODUCTION"

var RunMode string
var ServerPort = "8080"
var PageLimit = int64(10)
var isK8 = "False"
var KubeConfigFile = "dev-config.yaml"
var Auth = false
var InternalServer = false
var SkipServerTlsVerification = false
var TlsServerCustomCa = ""
var WorkerGroup string
var ServerUrl string
var AgentSecretName string
var ResourceNamespace string

func InitEnvironmentVariables() {
	RunMode = os.Getenv("RUN_MODE")
	if RunMode == "" {
		RunMode = DEVELOP
	}
	log.Logger.Infow("RUN MODE:", "value", RunMode)
	if RunMode != PRODUCTION {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Logger.Errorw("Failed to Load environment file", "err", err.Error())
			os.Exit(1)
		}
	}
	if os.Getenv("AUTH_ENABLED") == "TRUE" {
		Auth = true
		log.Logger.Infow("Started with AUTH enabled", "[AUTH]", Auth)
	} else {
		log.Logger.Infow("Started with AUTH disabled", "[AUTH]", Auth)
	}
	if os.Getenv("IS_INTERNAL_SERVER") == "TRUE" {
		InternalServer = true
		log.Logger.Infow("Server is internal. Skipping tls config", "[Internal-Server]", InternalServer)
	} else {
		log.Logger.Infow("Server is external. Using tls config", "[Internal-Server]", InternalServer)
	}
	if os.Getenv("SKIP_SERVER_TLS_VERIFICATION") == "TRUE" {
		SkipServerTlsVerification = true
		log.Logger.Infow("Server is skipping tls verification.", "[TLS-Skip-Verify]", SkipServerTlsVerification)
	} else {
		log.Logger.Infow("Server is verifying tls.", "[TLS-Skip-Verify]", SkipServerTlsVerification)
	}
	KubeConfigFile = os.Getenv("KUBE_CONFIG_FILE")
	isK8 = os.Getenv("IS_K8")
	TlsServerCustomCa = os.Getenv("TLS_SERVER_CUSTOM_CA")
	WorkerGroup = os.Getenv("WORKER_GROUP")
	ServerUrl = os.Getenv("SERVER_URL")
	AgentSecretName = os.Getenv("AGENT_SECRET_NAME")
	ResourceNamespace = os.Getenv("RESOURCE_NAMESPACE")
}

func IsK8() bool {
	if isK8 == "TRUE" {
		return true
	}
	return false
}

func IsAuth() bool {
	return Auth
}

func IsInternalServer() bool {
	return InternalServer
}

func IsSkipServerTlsVerification() bool {
	return SkipServerTlsVerification
}
