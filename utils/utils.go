package utils

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"spm/data"
	"spm/shared"
	"strings"

	"github.com/fatih/color"
)

func openProjectsFile(exePath string) ([]byte, error) {
	return os.ReadFile(filepath.Join(exePath, shared.PROJECT_DATA_FILEPATH))
}

func GetProjectData() (*data.ProjectData, error) {

	exePath, err := os.Executable()
	if err != nil {
		return nil, err
	}
	exePath = filepath.Dir(exePath)

	fileContents, err := openProjectsFile(exePath)
	if err != nil {
		return nil, err
	}

	var projectData data.ProjectData
	err = json.Unmarshal(fileContents, &projectData)
	projectData.ExePath = exePath

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
