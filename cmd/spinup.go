package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"spm/data"

	"github.com/fatih/color"
)

type SpinupCmd struct {
	*defaultCmd
}

func NewSpinupCmd() *SpinupCmd {

	flagSet := flag.NewFlagSet("spinup", flag.ContinueOnError)

	return &SpinupCmd{
		newDefaultCmd(flagSet),
	}
}

func (s *SpinupCmd) Execute(args []string, projData *data.ProjectData) error {
	path := filepath.Join(projData.ExePath, "projects.json")

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(path)
		if err != nil {
			return err
		}
		projData.Serialize()
		fmt.Println(color.GreenString("Created a new database"))
		return nil
	}
	return errors.New("database already exists")
}

func (s *SpinupCmd) Serialize(args []string, projData *data.ProjectData) error {
	return nil
}
