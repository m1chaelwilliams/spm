package cmd

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"spm/data"
	"spm/utils"
	"strings"

	"github.com/fatih/color"
)

type AddCmd struct {
	*defaultCmd
	override *bool
}

func NewAddCmd() *AddCmd {

	addFlagSet := flag.NewFlagSet("add", flag.ContinueOnError)
	addFlagSet.String("name", "", "name of project (blank for dirname as name)")
	addFlagSet.String("path", ".", "path of project (blank for cwd)")
	override := addFlagSet.Bool("override", false, "overrides any existing entry with the same name")

	return &AddCmd{
		defaultCmd: newDefaultCmd(addFlagSet),
		override:   override,
	}
}

func (a *AddCmd) Execute(args []string, projData *data.ProjectData) error {

	err := a.flagSet.Parse(args[2:])
	if err != nil {
		log.Fatal(err)
	}

	name := a.flagSet.Lookup("name").Value.String()
	path := a.flagSet.Lookup("path").Value.String()

	fmt.Printf("Name: %s\n", name)

	if path == "." {
		if len(a.flagSet.Args()) > 0 {
			path = a.flagSet.Arg(0)
		} else {
			return errors.New("not enough arguments. must at least provide a path")
		}
	}

	// if cwd
	if path == "." {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		path = cwd

		if name == "" {
			name = filepath.Base(path)
		}
	}

	// ensure the path exists
	if _, err := os.Stat(path); err != nil {
		return err
	}

	fmt.Printf("Adding: %s, %s\n", name, path)

	// get date added
	dateStr := utils.GetDateStr()

	newProj := *data.NewProject(
		name,
		path,
		map[string]any{
			"date_added":   dateStr,
			"last_queried": dateStr,
		},
	)

	addProj := true

	dupProj := projData.CheckDuplicates(&newProj)
	if dupProj != nil {

		var response string

		if !*a.override {
			fmt.Println(color.YellowString("Duplicate Project Found:"))
			fmt.Println(dupProj.ToStringDetailed())
			fmt.Println("Overrite?")
			fmt.Println("Y: Yes | N: No")

			fmt.Scanln(&response)
			response = strings.ToLower(strings.Trim(response, " \t"))
		} else {
			response = "y"
		}

		switch response {
		case "n":
			addProj = false
		case "y":
			projData.ReplaceProject(&newProj)
			return nil
		}
	}

	if addProj {
		projData.Projects = append(projData.Projects, &newProj)
		fmt.Println(color.GreenString("Added entry sucessfully."))
	}

	return nil
}

func (a *AddCmd) Serialize(args []string, projData *data.ProjectData) error {
	return projData.Serialize()
}

func (a *AddCmd) Help() string {
	return utils.StringifyFlagSet(a.flagSet)
}
