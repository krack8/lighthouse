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

func InitEnvironmentVariables() {
	RunMode = os.Getenv("RUN_MODE")
	if RunMode == "" {
		RunMode = DEVELOP
	}
	log.Logger.Infow("RUN MODE:", "value", RunMode)
	if RunMode != PRODUCTION {
		err := godotenv.Load("../.env-example")
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
	KubeConfigFile = os.Getenv("KUBE_CONFIG_FILE")
	isK8 = os.Getenv("IS_K8")
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
