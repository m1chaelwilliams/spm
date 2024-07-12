package cmd

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"tpm/data"

	"github.com/fatih/color"
)

//go:embed templates/tpmproj.json
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

	f, err := os.Create("tpmproj.json")
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
