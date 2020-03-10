package util

import (
	"fmt"
	"iads/lib/common"
	"iads/lib/file"
	"iads/lib/linux/hardware"
	"iads/lib/net"
	"strconv"
	"time"
)

type NodeIp struct {
	Bmc      string
	Ip       string
	User     string
	Password string
	Count    int
}

func rebootSubFunc(node NodeIp) {
	for {
		flag, err := net.IsPing(node.Bmc, "2s")
		if !flag {
			println("ping bmc ip locked: " + err.Error())
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}
	bmcCmd := fmt.Sprintf("ipmitool -H %s -I lanplus -U ADMIN -P ADMIN power reset", node.Bmc)
	for {
		flag, err := net.IsPing(node.Ip, "2s")
		if err != nil {
		}
		if flag {
			_, _ = common.ExecShellLinux(bmcCmd)
			node.Count++
			println(node.Bmc + " power reset. times:" + strconv.Itoa(node.Count))
		}
		time.Sleep(8 * time.Second)
	}
}

func RebootFunc() (err error) {
	if !file.IfFileExist("iplist.txt") {
		println("File iplist.txt is not exist, please create it first!")
		println("One line BMC address, next line system IP address, number of rows must be even.\neg. (only write address):\n\n192.168.1.1 --> bmc ip\n192.168.1.2 --> system ip\n192.168.1.8 --> bmc ip\n192.168.1.9 --> system ip")
		return
	}
	lines, err := file.ReadFileAsLine("iplist.txt")
	if err != nil {
		return
	}
	var nodes []NodeIp
	var node NodeIp
	node.Count = 0
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
		time.Sleep(1000 * time.Hour)
	}
	//return err
}

func SmcOobActiveFunc(bmcMac string) (code string, err error) {
	return hardware.OutActivationCode(bmcMac)
}
