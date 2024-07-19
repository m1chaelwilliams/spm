package cmd

import (
	"flag"
	"spm/data"
	"spm/utils"
)

type Command interface {
	Execute(args []string, projData *data.ProjectData) error
	Serialize(args []string, projData *data.ProjectData) error
	Help() string
}

type defaultCmd struct {
	flagSet *flag.FlagSet
}

func newDefaultCmd(flagSet *flag.FlagSet) *defaultCmd {
	return &defaultCmd{
		flagSet: flagSet,
	}
}

func (d *defaultCmd) Serialize(args []string, projData *data.ProjectData) error {
	return nil
}

func (d *defaultCmd) Help() string {
	return utils.StringifyFlagSet(d.flagSet)
}
