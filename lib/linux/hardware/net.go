package hardware

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/vishvananda/netlink"
)

type NetInterface struct {
	Name       string
	IpAddr     string
	Netmask    string
	Gateway    string
	Mac        string
	Type       string
	LinkStatus bool
	ni         netlink.Link
}

type NetInfo struct {
	DNS        string
	Interfaces arraylist.List
}

func (e *NetInfo) NetInit() (err error) {
	interfaceArr, err := netlink.LinkList()
	for _, value := range interfaceArr {
		if value.Attrs().EncapType == "ether" {
			ni := NetInterface{}
			ni.Name = value.Attrs().Name
			ni.Type = value.Attrs().EncapType
			ni.Mac = value.Attrs().HardwareAddr.String()
			if value.Attrs().OperState.String() == "up" {
				ni.LinkStatus = true
			} else if value.Attrs().OperState.String() == "down" {
				ni.LinkStatus = false
			}
			addrArr, _ := netlink.AddrList(value, netlink.FAMILY_V4)
			routeArr, _ := netlink.RouteList(value, netlink.FAMILY_V4)
			if len(addrArr) > 0 {
				ni.IpAddr = addrArr[0].IP.String()
				ni.Netmask = addrArr[0].Mask.String()
			}
			if len(routeArr) > 0 {
				ni.Gateway = routeArr[0].Gw.String()
			}
			ni.ni = value
			e.Interfaces.Add(ni)
		}
	}
	return
}

func (e NetInfo) Print() {
	it := e.Interfaces.Iterator()
	for it.Next() {
		println(it.Value().(NetInterface).Gateway)
	}
}
