package config

import (
	"github.com/joho/godotenv"
	"github.com/krack8/lighthouse/pkg/log"
	"os"
)

var ServerPort = "8080"
var PageLimit = int64(10)
var isK8 = false
var KubeConfigFile = "dev-config.yaml"
var NoAuth = false

func IsK8() bool {
	return isK8
}

func InitEnvironmentVariables() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Logger.Errorw("Failed to Load environment file", "err", err.Error())
		os.Exit(1)
	}
	if os.Getenv("NO_AUTH") == "TRUE" {
		NoAuth = true
		log.Logger.Infow("Started with NO AUTH enabled", "[NO-AUTH]", NoAuth)
	} else {
		log.Logger.Infow("Started with NO AUTH disabled", "[NO-AUTH]", NoAuth)
	}
	KubeConfigFile = os.Getenv("KUBE_CONFIG_FILE")
}

func IsNoAuth() bool {
	return NoAuth
}
