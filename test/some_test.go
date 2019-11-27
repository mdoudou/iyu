package test

import (
	"fmt"
	"net"
	"regexp"
	"testing"
)

type A struct {
	List []string `json:"list"`
}

func TestSome(t *testing.T) {
	regex := `[\P{Han}\W]+`
	fmt.Println(regexp.MatchString(regex, "汉字"))
}

func getMacAddrs() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v", err)
		return macAddrs
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}

		macAddrs = append(macAddrs, macAddr)
	}
	return macAddrs
}
