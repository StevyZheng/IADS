package disk

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/spf13/cobra"
	"iads/lib/linux/hardware"
	"iads/lib/logging"
	"os"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print the disk list",
	Run: func(cmd *cobra.Command, args []string) {
		t := table.NewWriter()
		alignT := []text.Align{text.AlignCenter, text.AlignCenter, text.AlignCenter, text.AlignCenter, text.AlignCenter, text.AlignCenter}
		t.SetOutputMirror(os.Stdout)
		t.SetAlignHeader(alignT)
		t.SetAlign(alignT)
		t.Style().Options.SeparateRows = true
		t.Style().Box = table.StyleBoxBold
		t.SetAutoIndex(true)

		t.AppendHeader(table.Row{"DEV", "MODEL", "SN", "WWN", "SIZE", "TYPE"})
		disks, err := hardware.Disk{}.DiskList()
		if err != nil {
			logging.FatalPrintln("Disk list error.")
			return
		}
		for _, disk := range disks {
			t.AppendRow(table.Row{disk.DevName, disk.Model, disk.Serial, disk.Wwn, fmt.Sprintf("%dGB", disk.Size), disk.DevType})
		}
		t.Render()
	},
}
