package go_kit

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"path/filepath"
	"github.com/yidianyipie/go-kit/constant"
	"runtime"
	"strings"
	"time"
)

var (
	startTime string
	goVersion string
)

// build info
/*
 */
var (
	appName         string
	appID           string
	hostName        string
	buildAppVersion string
	buildUser       string
	buildHost       string
	buildStatus     string
	buildTime       string
)

func init() {
	if appName == "" {
		appName = os.Getenv(constant.EnvAppName)
		if appName == "" {
			appName = filepath.Base(os.Args[0])
		}
	}
	if appID == "" {
		appID = os.Getenv(constant.EnvAppID)
		if appID == "" {
			appID = "1234567890"
		}
	}
	name, err := os.Hostname()
	if err != nil {
		name = "unknown"
	}
	hostName = name
	startTime = time.Now().Format("2006-01-02 15:04:05")
	SetBuildTime(buildTime)
	goVersion = runtime.Version()
	InitEnv()
}

// Name gets application name.
func Name() string {
	return appName
}

// SetName set app name
func SetName(s string) {
	appName = s
}

// AppID get appID
func AppID() string {
	return appID
}

// SetAppID set appID
func SetAppID(s string) {
	appID = s
}

// AppVersion get buildAppVersion
func AppVersion() string {
	return buildAppVersion
}

// SetAppVersion set appVersion
func SetAppVersion(s string) {
	buildAppVersion = s
}

// BuildTime get buildTime
func BuildTime() string {
	return buildTime
}

// BuildUser get buildUser
func BuildUser() string {
	return buildUser
}

// BuildHost get buildHost
func BuildHost() string {
	return buildHost
}

// SetBuildTime set buildTime
func SetBuildTime(param string) {
	buildTime = strings.Replace(param, "--", " ", 1)
}

// HostName get host name
func HostName() string {
	return hostName
}

// StartTime get start time
func StartTime() string {
	return startTime
}

// GoVersion get go version
func GoVersion() string {
	return goVersion
}

func LogDir() string {
	// LogDir gets application log directory.
	logDir := AppLogDir()
	if logDir == "" {
		if appPodIP != "" && appPodName != "" {
			// k8s 环境
			return fmt.Sprintf("/tmp/logs/%s/%s/", Name(), appPodName)
		}
		return fmt.Sprintf("/tmp/logs/%s/%s/", Name(), appInstance)
	}
	return fmt.Sprintf("%s/%s/%s/", logDir, Name(), appInstance)
}

// PrintVersion print formatted version info
func PrintVersion() {
	fmt.Printf("%-8s]> %-30s => %s\n", "go-kit", color.RedString("name"), color.BlueString(appName))
	fmt.Printf("%-8s]> %-30s => %s\n", "go-kit", color.RedString("appID"), color.BlueString(appID))
	fmt.Printf("%-8s]> %-30s => %s\n", "go-kit", color.RedString("region"), color.BlueString(AppRegion()))
	fmt.Printf("%-8s]> %-30s => %s\n", "go-kit", color.RedString("zone"), color.BlueString(AppZone()))
	fmt.Printf("%-8s]> %-30s => %s\n", "go-kit", color.RedString("appVersion"), color.BlueString(buildAppVersion))
	fmt.Printf("%-8s]> %-30s => %s\n", "go-kit", color.RedString("buildUser"), color.BlueString(buildUser))
	fmt.Printf("%-8s]> %-30s => %s\n", "go-kit", color.RedString("buildHost"), color.BlueString(buildHost))
	fmt.Printf("%-8s]> %-30s => %s\n", "go-kit", color.RedString("buildTime"), color.BlueString(BuildTime()))
	fmt.Printf("%-8s]> %-30s => %s\n", "go-kit", color.RedString("buildStatus"), color.BlueString(buildStatus))
}
