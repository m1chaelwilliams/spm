package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"tpm/data"
	"tpm/utils"

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

	if len(a.flagSet.Args()) > 0 {
		path = a.flagSet.Arg(0)
	}

	// if cwd
	if path == "." {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		path = cwd
	}

	// ensure the path exists
	if _, err := os.Stat(path); err != nil {
		return err
	}

	if name == "" {
		name = filepath.Base(path)
	}

	fmt.Printf("Adding: %s, %s\n", name, path)

	newProj := *data.NewProject(
		name,
		path,
		make(map[string]any, 0),
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
