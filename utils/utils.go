package utils

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"

	"spm/data"
	"spm/shared"
)

func openProjectsFile(exePath string) ([]byte, error) {
	return os.ReadFile(filepath.Join(exePath, shared.PROJECT_DATA_FILEPATH))
}

func GetProjectPath() (string, error) {
	switch runtime.GOOS {
	case "linux":
		xdgHome, exists := os.LookupEnv("XDG_DATA_HOME")
		if !exists {
			home, exists := os.LookupEnv("HOME")
			if !exists {
				return "", errors.New(
					"$HOME is not set. Unable to locate proper data storage directory",
				)
			}
			return fmt.Sprintf("%s/.local/share/spm", home), nil
		}
		return fmt.Sprintf("%s/spm", xdgHome), nil
	default:
		return "", fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
}

func GetProjectData(arg string) (*data.ProjectData, error) {
	projectDataPath, err := GetProjectPath()
	if err != nil {
		return nil, err
	}
	os.MkdirAll(projectDataPath, 0o777)

	if arg == "spinup" {
		return &data.ProjectData{
			Projects: make([]*data.Project, 0),
			DataPath: projectDataPath,
		}, nil
	}

	fileContents, err := openProjectsFile(projectDataPath)
	if err != nil {
		return nil, err
	}

	var projectData data.ProjectData
	err = json.Unmarshal(fileContents, &projectData)
	projectData.DataPath = projectDataPath

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

func GetDateStr() string {
	dateAdded := time.Now()
	month, day, year := dateAdded.Month(), dateAdded.Day(), dateAdded.Year()
	return fmt.Sprintf("%d/%d/%d", month, day, year)
}

// checks if date1 < date2. if so, return true
func IsDateStrLess(date1, date2 string) bool {
	date1Split := strings.Split(date1, "/")
	date2Split := strings.Split(date2, "/")

	month1, day1, year1 := date1Split[0], date1Split[1], date1Split[2]
	month2, day2, year2 := date2Split[0], date2Split[1], date2Split[2]

	if year1 < year2 {
		return true
	}

	if month1 < month2 {
		return true
	}

	if day1 < day2 {
		return true
	}

	return false
}

func IsSupportedSortStrategy(metaTarget string) bool {
	switch metaTarget {
	case "date_added", "last_queried", "alphabet":
		return true
	default:
		return false
	}
}
