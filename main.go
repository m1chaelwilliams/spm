package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"

	"github.com/vimichael/spm/cmd"
	"github.com/vimichael/spm/utils"
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

	projData, err := utils.GetProjectData(args[1])
	if err != nil {
		log.Fatal(redErr(err))
	}

	err = command.Execute(args, projData)
	if err != nil {
		log.Fatal(redErr(err))
	}
	command.Serialize(args, projData)
}
