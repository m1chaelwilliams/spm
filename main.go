package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"tpm/cmd"
	"tpm/utils"

	"github.com/fatih/color"
)

func redErr(err error) string {
	errorStr := fmt.Sprintf("%s", err)
	return color.RedString(errorStr)
}

func main() {

	args := os.Args
	if len(args) < 2 {
		log.Fatal(redErr(errors.New("not enough arguments")))
	}

	// get the command
	command, err := cmd.GetCommand(args)
	if err != nil {
		log.Fatal(redErr(err))
	}

	projData, err := utils.GetProjectData()
	if err != nil {
		log.Fatal(redErr(err))
	}

	err = command.Execute(args, projData)
	if err != nil {
		log.Fatal(redErr(err))
	}
	command.Serialize(args, projData)
}
