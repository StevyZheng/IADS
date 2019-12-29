package util

type VersionText struct {
	Version    string
	UpdateText string
}

var versionVar = []VersionText{
	{
		"1.0.2",
		`version: 1.0.2
1. Add Zstack reboot text.
2. Optimized the server rest API functionality.
3. Add smc oob active function.`,
	},
	{
		"1.0.1",
		`version: 1.0.1
1. Compatible with all linux,include ubuntu.
2. Add version change description.
3. Add CentOS reboot test module.`,
	},
}

func VersionPrint() {
	for _, v := range versionVar {
		println(v.UpdateText)
		println("")
	}
}
