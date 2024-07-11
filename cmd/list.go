package cmd

import (
	"flag"
	"fmt"
	"tpm/data"

	"github.com/fatih/color"
)

type List struct {
	*defaultCmd
	detailed *bool
}

func NewList() *List {
	flagSet := flag.NewFlagSet("list", flag.ContinueOnError)
	detailed := flagSet.Bool("detailed", false, "give more info for each project")

	return &List{
		defaultCmd: newDefaultCmd(flagSet),
		detailed:   detailed,
	}
}

func (l *List) Execute(args []string, projData *data.ProjectData) error {

	err := l.flagSet.Parse(args[2:])
	if err != nil {
		return err
	}

	fmt.Println(color.BlueString("Projects:\n"))
	if *l.detailed {
		for _, project := range projData.Projects {
			fmt.Printf("%s\n", project.ToStringDetailed())
		}
	} else {
		for _, project := range projData.Projects {
			fmt.Printf("%s\n", project.ToString())
		}
	}

	return nil
}
