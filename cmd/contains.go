package cmd

import (
	"flag"
	"fmt"
	"strings"
	"tpm/data"

	"github.com/fatih/color"
)

type ContainsCmd struct {
	*defaultCmd
}

func NewContainsCmd() *ContainsCmd {

	flagSet := flag.NewFlagSet("contains", flag.ContinueOnError)
	flagSet.String("name", "", "text to compare against")

	return &ContainsCmd{
		newDefaultCmd(flagSet),
	}
}

func (c *ContainsCmd) Execute(args []string, projData *data.ProjectData) error {
	err := c.flagSet.Parse(args[2:])
	if err != nil {
		return err
	}

	text := c.flagSet.Lookup("name").Value.String()

	similarProjects := make([]*data.Project, 0)
	for _, project := range projData.Projects {
		if strings.Contains(project.Name, text) {
			similarProjects = append(similarProjects, project)
		}
	}

	if len(similarProjects) < 1 {
		fmt.Println(color.YellowString("no projects contain: %s\n", text))
	} else {
		for _, project := range similarProjects {
			fmt.Printf("- %s\n", project.ToStringDetailed())
		}
	}

	return nil
}
