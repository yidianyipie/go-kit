package errcode

import "fmt"

const (
	systemErrorCode uint32 = 1 * 1000000
)

var (
	PackageNum = map[string]uint32{
		"night-watcher": 1,
		"dbcm":          2,
		"fip":           3,
		"netops":        4,
		"nginx-manager": 5,
		"dns-manager":   6,
		"yide-manager":  7,
		"cdn-manager":   8,
		"mozis":         9,
		"sso":           10,
		"heimdallr":     11,
		"enos":          12,
		"cmp":           13,
		"workflow":      14,
	}
)

var (
	packageCode uint32
)

func SetPackageCode(packageName string) {
	packageOriginNum, exist := PackageNum[packageName]
	if !exist {
		panic(fmt.Sprintf("package(%s) 未注册", packageName))
	}
	packageCode = packageOriginNum * 10000
}
