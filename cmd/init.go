package cmd

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"

	"github.com/vimichael/spm/data"
)

//go:embed templates/spmproj.json
var template []byte

type InitCmd struct {
	*defaultCmd
}

func NewInitCmd() *InitCmd {
	flagSet := flag.NewFlagSet("init", flag.ContinueOnError)

	return &InitCmd{
		defaultCmd: newDefaultCmd(flagSet),
	}
}

func (i *InitCmd) Execute(args []string, projData *data.ProjectData) error {
	err := i.flagSet.Parse(args[2:])
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	path := cwd

	f, err := os.Create("github.com/vimichael/spmproj.json")
	if err != nil {
		return err
	}

	_, err = f.Write(template)
	if err != nil {
		return err
	}

	fmt.Println(color.GreenString("Created template at: %s", path))

	return nil
}

func (i *InitCmd) Serialize(args []string, projData *data.ProjectData) error {
	return projData.Serialize()
}
