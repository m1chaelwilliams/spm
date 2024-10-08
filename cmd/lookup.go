package cmd

import (
	"flag"
	"fmt"
	"spm/data"
	"spm/utils"

	"github.com/fatih/color"
)

type LookupCmd struct {
	*defaultCmd
}

func NewLookupCmd() *LookupCmd {

	flagSet := flag.NewFlagSet("lookup", flag.ContinueOnError)
	flagSet.String("name", "", "name of the project")

	return &LookupCmd{
		defaultCmd: newDefaultCmd(
			flagSet,
		),
	}
}

func (l *LookupCmd) Execute(args []string, projData *data.ProjectData) error {

	err := l.flagSet.Parse(args[2:])
	if err != nil {
		return err
	}

	name := l.flagSet.Lookup("name").Value.String()

	if name == "" && len(l.flagSet.Args()) > 0 {
		name = l.flagSet.Arg(0)
	}

	if proj, exists := projData.FindProject(name); exists {
		fmt.Println(color.GreenString("Name: %s", proj.Name))
		fmt.Println(color.BlueString("Path: %s", proj.Path))
		fmt.Println(color.GreenString("Metadata:"))
		for key, value := range proj.MetaData {
			fmt.Printf("- %-20s = %s\n", color.BlueString(key), color.YellowString("%v", value))
		}

		proj.MetaData["last_queried"] = utils.GetDateStr()
		return nil
	}

	return fmt.Errorf("could not find project: %s", name)
}

func (l *LookupCmd) Serialize(args []string, projData *data.ProjectData) error {
	return projData.Serialize()
}
