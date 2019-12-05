package hardware

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"iads/lib/common"
	"log"
	"os"
	"strings"
)

func OutActivationCode(bmcMac string) (code string, err error) {
	cmd := fmt.Sprintf("echo -n '%s'| xxd -r -p | openssl dgst -sha1 -mac HMAC -macopt hexkey:8544E3B47ECA58F9583043F8 | awk '{print $2}' | cut -c 1-24", bmcMac)
	code, err = common.ExecShellLinux(cmd)
	return code, err
}

func MacAddColon(srcMac string) (destMac string) {
	destMac = ""
	srcMacLen := len(srcMac)
	if srcMacLen == 17 || srcMacLen == 12 {
		if strings.Index(srcMac, ":") == -1 {
			for i, ch := range srcMac {
				destMac = destMac + string(ch)
				if i%2 != 0 && i != (srcMacLen-1) {
					destMac = destMac + ":"
				}
			}
		} else {
			return srcMac
		}
	} else {
		return destMac
	}

	return destMac
}

func HandleXlsx(fileName string) (err error) {
	xlsxFile, err := xlsx.OpenFile(fileName)
	if err != nil {
		println(err)
		os.Exit(1)
	}
	sheet := xlsxFile.Sheet["Sheet1"]
	if sheet == nil {
		println("Sheet1 is not exist.")
		os.Exit(1)
	}
	for i, row := range sheet.Rows {
		if len(row.Cells) > 0 {
			mac := row.Cells[0].Value
			mac = MacAddColon(mac)
			println("Crack MAC: ", mac, "  -- Done.")
			if false == RegMac(mac) || mac == "" {
				println("error: mac addr is wrong!")
				println("row: ", i)
			}
			code, err := OutActivationCode(mac)
			if err != nil {
				log.Println(err.Error())
				return err
			}
			cell := row.AddCell()
			cell.Value = code
		}
	}
	err = xlsxFile.Save("codeInfo.xlsx")

	return err
}
