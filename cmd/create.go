package cmd

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"path"
)

func Create(args cli.Args) error {
	setting := args.First()
	if !isValidSettingName(setting) {
		return errors.New("invalid setting. setting cannot be 'config' or 'credentials'")
	}

	if hasSetting(setting) {
		return errors.New(fmt.Sprintf("setting, '%s' already exists", setting))
	}

	err := createEmptyFile(path.Join(configDir(), setting+ExtCredentials))
	if err != nil {
		return err
	}

	err = createEmptyFile(path.Join(configDir(), setting+ExtConfig))
	return err
}


func createEmptyFile(p string) error {
	_, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE, 0600)
	return err
}