package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"spm/data"
)

type LoadCmd struct {
	*defaultCmd
}

func NewLoadCmd() *LoadCmd {

	flagSet := flag.NewFlagSet("load", flag.ContinueOnError)
	flagSet.String("path", ".", "path to the project data file")

	return &LoadCmd{
		defaultCmd: newDefaultCmd(flagSet),
	}
}

func (l *LoadCmd) Execute(args []string, projData *data.ProjectData) error {
	err := l.flagSet.Parse(args[2:])
	if err != nil {
		return err
	}

	path := l.flagSet.Lookup("path").Value.String()

	fmt.Println("Reading path")

	// if cwd
	if path == "." {
		if len(l.flagSet.Args()) > 0 {
			path = filepath.Join(l.flagSet.Arg(0), "tpmproj.json")
		} else {
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}
			path = filepath.Join(cwd, "tpmproj.json")
		}
	}

	fmt.Println("ensuring file exists")

	// ensure the path exists
	if _, err := os.Stat(path); err != nil {
		return err
	}

	fmt.Println("reading file contents")

	fmt.Printf("Reading: %s\n", path)
	fileContents, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	fmt.Println("contents read")

	var newProj data.Project
	err = json.Unmarshal(fileContents, &newProj)
	if err != nil {
		return err
	}

	// if cwd
	if newProj.Path == "." {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		newProj.Path = cwd
	}

	// ensure the path exists
	if _, err := os.Stat(newProj.Path); err != nil {
		return err
	}

	fmt.Println("Adding project")

	// if project does not exist, append new project
	if err = projData.UpdateProject(&newProj); err != nil {
		projData.Projects = append(projData.Projects, &newProj)
	}

	return nil
}

func (l *LoadCmd) Serialize(args []string, projData *data.ProjectData) error {
	return projData.Serialize()
}
