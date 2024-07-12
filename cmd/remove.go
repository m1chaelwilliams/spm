package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"tpm/data"

	"github.com/fatih/color"
)

type RemoveCmd struct {
	*defaultCmd
}

func NewRemoveCmd() *RemoveCmd {

	flagSet := flag.NewFlagSet("remove", flag.ContinueOnError)
	flagSet.String("name", "", "name of the project (blank for dirname as name)")

	return &RemoveCmd{
		defaultCmd: newDefaultCmd(flagSet),
	}
}

func (r *RemoveCmd) Execute(args []string, projData *data.ProjectData) error {

	err := r.flagSet.Parse(args[2:])
	if err != nil {
		return err
	}

	name := r.flagSet.Lookup("name").Value.String()

	if name == "" {
		if len(r.flagSet.Args()) > 0 {
			name = r.flagSet.Arg(0)
		} else {
			return fmt.Errorf("no project name provided")
		}
	}
	if name == "." {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		name = filepath.Base(cwd)
	}

	err = projData.RemoveProject(name)
	if err != nil {
		return err
	}

	fmt.Println(color.GreenString("%s removed\n", name))
	return nil
}

func (r *RemoveCmd) Serialize(args []string, projData *data.ProjectData) error {
	return projData.Serialize()
}
