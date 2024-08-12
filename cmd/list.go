package cmd

import (
	"flag"
	"fmt"
	"spm/data"

	"github.com/fatih/color"
)

type List struct {
	*defaultCmd
	detailed *bool
	sortBy   *string
}

func NewList() *List {
	flagSet := flag.NewFlagSet("list", flag.ContinueOnError)
	detailed := flagSet.Bool("detailed", false, "give more info for each project")
	sortBy := flagSet.String("sortBy", "alphabet", "sort by this metadata field.")

	return &List{
		defaultCmd: newDefaultCmd(flagSet),
		detailed:   detailed,
		sortBy:     sortBy,
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
