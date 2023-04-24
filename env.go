package go_kit

import (
	"crypto/md5"
	"fmt"
	"os"
	"practice.com/go-kit/constant"
)

var (
	appLogDir   string
	appMode     string
	appRegion   string
	appZone     string
	appHost     string
	appInstance string
	appPodIP    string
	appPodName  string
)

func InitEnv() {
	appLogDir = os.Getenv(constant.EnvAppLogDir)
	appMode = os.Getenv(constant.EnvAppMode)
	appRegion = os.Getenv(constant.EnvAppRegion)
	appZone = os.Getenv(constant.EnvAppZone)
	appHost = os.Getenv(constant.EnvAppHost)
	appInstance = os.Getenv(constant.EnvAppInstance)
	if appInstance == "" {
		appInstance = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s:%s", HostName(), AppID()))))
	}
	appPodIP = os.Getenv(constant.EnvPodIP)
	appPodName = os.Getenv(constant.EnvPodName)
}

func AppLogDir() string {
	return appLogDir
}

func SetAppLogDir(logDir string) {
	appLogDir = logDir
}

func AppMode() string {
	if appMode == "" {
		return "dev"
	}
	return appMode
}

func SetAppMode(mode string) {
	appMode = mode
}

func AppRegion() string {
	return appRegion
}

func SetAppRegion(region string) {
	appRegion = region
}

func AppZone() string {
	if appZone == "" {
		return "unknown"
	}
	return appZone
}

func SetAppZone(zone string) {
	appZone = zone
}

func AppHost() string {
	return appHost
}

func SetAppHost(host string) {
	appHost = host
}

func AppInstance() string {
	return appInstance
}

func SetAppInstance(instance string) {
	appInstance = instance
}