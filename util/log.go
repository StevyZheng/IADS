package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"iads/lib/file"
	"iads/lib/stringx"
	"io"
	. "os"
	"strings"
	"time"
)

type ErrorMap struct {
	Error    string
	Analysis string
}

var errVar = []ErrorMap{
	{"error", "normal."},
	{"fail", "normal."},
	{"io error", "io error."},
	{"mce error", "hardware error: like cpu,memory."},
}

type MessagesType struct {
	Time    time.Time
	Trigger string
	Msg     string
}

func messagesError() (msgVar []MessagesType, err error) {
	monthVar := []string{"jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sept", "Oct", "Nov", "Dec"}
	msgVar = make([]MessagesType, 0)

	logFiles, err := file.ListFiles("/var/log")
	for _, v := range logFiles {
		if -1 < strings.Index(v, "messages") || -1 < strings.Index(v, "kerl") {
			msgFd, err := OpenFile(v, O_RDONLY, 0)
			if err != nil {
				println("Open %s failed!", v)
				return nil, err
			}
			defer msgFd.Close()
			rd := bufio.NewReader(msgFd)
			for {
				data, _, eof := rd.ReadLine()
				if eof == io.EOF {
					break
				}
				line := string(data)
				for _, v := range errVar {
					if -1 < strings.Index(line, v.Error) {
						msg := MessagesType{}
						splitVar := stringx.SplitString(line, " ")
						for i, v := range monthVar {
							if v == splitVar[0] {
								timeTpm := fmt.Sprintf("%d-%d-%s %s", 2019, i+1, splitVar[1], splitVar[2])
								lt, _ := time.ParseInLocation("2006-01-02 15:04:05", timeTpm, time.Local)
								msg.Time = lt
								msg.Trigger = splitVar[3] + splitVar[4]
								msg.Msg = strings.Join(splitVar[4:], " ")
								break
							}
						}
						msgVar = append(msgVar, msg)
						break
					}
				}
			}
		}
	}
	return msgVar, err
}

func MessagesErrorPrint() (err error) {
	msgVar, err := messagesError()
	t := table.NewWriter()
	alignT := []text.Align{text.AlignCenter, text.AlignCenter, text.AlignCenter, text.AlignCenter, text.AlignCenter, text.AlignCenter}
	t.SetOutputMirror(Stdout)
	t.SetAlignHeader(alignT)
	t.SetAlign(alignT)
	t.Style().Options.SeparateRows = true
	t.Style().Box = table.StyleBoxBold
	t.SetAutoIndex(true)

	t.AppendHeader(table.Row{"TIME", "PROCESS", "MSG"})

	for _, v := range msgVar {
		t.AppendRow(table.Row{v.Time, v.Trigger, v.Msg})
	}
	t.Render()
	return err
}

func MessagesErrorJson() (string, error) {
	msgVar, err := messagesError()
	b, err := json.MarshalIndent(msgVar, "", "    ")
	if err != nil {
		println("Json error:", err)
	}
	return string(b), err
}
