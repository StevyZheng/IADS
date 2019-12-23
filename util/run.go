package util

import (
	"fmt"
	"iads/lib/common"
	"iads/lib/file"
	"iads/lib/linux/hardware"
	"iads/lib/net"
	"time"
)

type NodeIp struct {
	Bmc string
	Ip  string
}

func rebootSubFunc(node NodeIp) {
	bmcCmd := fmt.Sprintf("ipmitool -H %s -I lanplus -U ADMIN -P ADMIN power reset", node.Bmc)
	for {
		flag := net.IsPing(node.Ip, "2s")
		if flag {
			_, _ = common.ExecShellLinux(bmcCmd)
		} else {
			time.Sleep(4 * time.Second)
		}
	}
}

func RebootFunc() (err error) {
	if !file.IfFileExist("iplist.txt") {
		println("File iplist.txt is not exist, please create it first!")
		println("One line BMC address, next line system IP address, number of rows must be even.\neg. (only write address):\n\n192.168.1.1 --> bmc ip\n192.168.1.2 --> system ip\n192.168.1.8 --> bmc ip\n192.168.1.9 --> system ip")
		return
	}
	lines, err := file.ReadFileAsLine("iplist.txt")
	if err == nil {
		return
	}
	var nodes []NodeIp
	var node NodeIp
	for index, value := range lines {
		if index%2 == 0 {
			node.Bmc = value
		} else {
			node.Ip = value
			nodes = append(nodes, node)
		}
	}
	for _, value := range nodes {
		go rebootSubFunc(value)
	}
	for {
		time.Sleep(3600 * time.Second)
	}
	//return err
}

func SmcOobActiveFunc(bmcMac string) (code string, err error) {
	return hardware.OutActivationCode(bmcMac)
}
