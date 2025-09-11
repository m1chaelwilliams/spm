package cmd

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/fatih/color"

	"github.com/vimichael/spm/data"
)

type List struct {
	*defaultCmd
	detailed     *bool
	sortBy       *string
	outputAsJson *bool
}

func NewList() *List {
	flagSet := flag.NewFlagSet("list", flag.ContinueOnError)
	detailed := flagSet.Bool("detailed", false, "give more info for each project")
	sortBy := flagSet.String("sortBy", "alphabet", "sort by this metadata field.")
	outputAsJson := flagSet.Bool("json", false, "format output as JSON")

	return &List{
		defaultCmd:   newDefaultCmd(flagSet),
		detailed:     detailed,
		sortBy:       sortBy,
		outputAsJson: outputAsJson,
	}
}

func (l *List) Execute(args []string, projData *data.ProjectData) error {
	err := l.flagSet.Parse(args[2:])
	if err != nil {
		return err
	}

	// if json output is requested
	if *l.outputAsJson {
		data := map[string][]*data.Project{
			"data": projData.Projects,
		}
		contents, err := json.Marshal(data)
		if err != nil {
			return err
		}

		fmt.Printf(string(contents))
		return nil
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
