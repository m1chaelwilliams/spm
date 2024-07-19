package cmd

import (
	"flag"
	"fmt"
	"spm/data"

	"github.com/fatih/color"
	"golang.design/x/clipboard"
)

type CopyPathCmd struct {
	*defaultCmd
}

func NewCopyPathCmd() *CopyPathCmd {

	flagSet := flag.NewFlagSet("copypath", flag.ContinueOnError)
	flagSet.String("name", "", "name of the project")

	return &CopyPathCmd{
		defaultCmd: newDefaultCmd(flagSet),
	}
}

func (c *CopyPathCmd) Execute(args []string, projData *data.ProjectData) error {

	err := c.flagSet.Parse(args[2:])
	if err != nil {
		return err
	}

	name := c.flagSet.Lookup("name").Value.String()
	if name == "" {
		if len(c.flagSet.Args()) > 0 {
			name = c.flagSet.Arg(0)
		}
	}

	var path *string = nil

	for _, project := range projData.Projects {
		if name == project.Name {
			path = &project.Path
		}
	}

	if path != nil {
		err := clipboard.Init()
		if err != nil {
			return err
		}

		clipboard.Write(clipboard.FmtText, []byte(*path))
		fmt.Println(color.GreenString("Copied: %s\n", *path))
	} else {
		return fmt.Errorf("project with name: %s not found", name)
	}

	return nil
}
