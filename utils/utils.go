package utils

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"tpm/data"

	"github.com/fatih/color"
)

const projectFilePath = "projects.json"

func openProjectsFile() ([]byte, error) {
	return os.ReadFile(projectFilePath)
}

func GetProjectData() (*data.ProjectData, error) {

	fileContents, err := openProjectsFile()
	if err != nil {
		return nil, err
	}

	var projectData data.ProjectData
	err = json.Unmarshal(fileContents, &projectData)

	return &projectData, err
}

func StringifyFlagSet(flagSet *flag.FlagSet) string {
	var flags strings.Builder

	flagSet.VisitAll(func(f *flag.Flag) {
		flags.WriteString(
			fmt.Sprintf("\t- %s: %s\n", color.YellowString(f.Name), f.Usage),
		)
	})

	return fmt.Sprintf(
		"%s:\n%s",
		color.GreenString(flagSet.Name()),
		flags.String(),
	)
}
