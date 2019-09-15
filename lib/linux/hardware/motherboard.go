package hardware

import (
	"fmt"
	"iads/lib/common"
	"iads/lib/stringx"
)

type MotherboradInfo struct {
	Model    string
	BiosVer  string
	BiosDate string
	BmcVer   string
	BmcDate  string
}

func (e *MotherboradInfo) GetMbInfo() {
	tmpStr, err := common.ExecShellLinux("dmidecode")
	if err != nil {
		fmt.Println(err)
	}
	e.Model = stringx.SearchSplitStringColumnFirst(tmpStr, ".*Product Name.*", ":", 2)
	e.BiosVer = stringx.SearchSplitStringColumnFirst(tmpStr, ".*Version.*", ":", 2)
	e.BiosDate = stringx.SearchSplitStringColumnFirst(tmpStr, ".*Release Date.*", ":", 2)
}
