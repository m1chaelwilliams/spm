package cmd

import (
	"flag"
	"fmt"
	"sort"
	"tpm/data"

	"github.com/fatih/color"
)

type HelpCmd struct {
	*defaultCmd
}

func NewHelpCmd() *HelpCmd {
	return &HelpCmd{
		defaultCmd: newDefaultCmd(flag.NewFlagSet("help", flag.ContinueOnError)),
	}
}

func (h *HelpCmd) Execute(args []string, projData *data.ProjectData) error {

	fmt.Println(color.BlueString("TMP | Terminal Project Manager"))

	keys := make([]string, 0)
	for key := range commands {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(commands[key].Help())
	}

	return nil
}

func (h *HelpCmd) Help() string {
	return ""
}
