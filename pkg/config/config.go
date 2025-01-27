package config

var ServerPort = "8080"
var PageLimit = int64(10)
var isK8 = false
var KubeConfigFile = "dev-config.yaml"

func IsK8() bool {
	return isK8
}
