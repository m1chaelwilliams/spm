package cmd

import (
	"flag"
	"fmt"
	"spm/data"

	"github.com/fatih/color"
)

type EditCmd struct {
	*defaultCmd
}

func NewEditCmd() *EditCmd {

	flagSet := flag.NewFlagSet("edit", flag.ContinueOnError)
	flagSet.String("target", "", "name of the target project")
	flagSet.String("name", "", "name of the target project")
	flagSet.String("path", "", "path of the target project")

	return &EditCmd{
		defaultCmd: newDefaultCmd(flagSet),
	}
}

func (e *EditCmd) Execute(args []string, projData *data.ProjectData) error {

	err := e.flagSet.Parse(args[2:])
	if err != nil {
		return err
	}

	targetName := e.flagSet.Lookup("target").Value.String()
	if targetProj, exists := projData.FindProject(targetName); exists {

		name := e.flagSet.Lookup("name").Value.String()
		path := e.flagSet.Lookup("path").Value.String()

		if len(name) > 0 {
			targetProj.Name = name
		}
		if len(path) > 0 {
			targetProj.Path = path
		}

		fmt.Println(color.GreenString("%s edited successfully", targetName))
		return nil

	}

	return fmt.Errorf("%s does not exist", targetName)
}
