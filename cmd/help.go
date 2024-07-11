package cmd

import (
	"flag"
	"fmt"
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
	for _, cmd := range commands {
		fmt.Println(cmd.Help())
	}

	return nil
}

func (h *HelpCmd) Help() string {
	return ""
}
