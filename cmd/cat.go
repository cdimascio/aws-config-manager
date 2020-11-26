package cmd

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
)

func Cat(args cli.Args, fileType string) error {
	var setting string
	if args.Len() >= 1 {
		setting = args.First()
		if !hasSetting(setting) {
			return errors.New(fmt.Sprintf("setting '%s' does not exist", setting))
		}
	} else {
		setting = current()
	}

	file, err := pathFromType(setting, fileType)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	fmt.Print(string(data))
	return nil
}
