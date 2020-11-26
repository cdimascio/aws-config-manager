package cmd

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"path"
)

func Use(args cli.Args) error {
	setting := args.First()
	if !hasSetting(setting) {
		return errors.New(fmt.Sprintf("setting '%s' does not exist", setting))
	}

	_ = os.Remove(awsCredentials())
	_ = os.Remove(awsConfig())

	settingCredFile := path.Join(configDir(), setting+ExtCredentials)
	settingConfigFile := path.Join(configDir(), setting+ExtConfig)

	err := os.Symlink(settingCredFile, awsCredentials())
	if err != nil {
		return err
	}

	err = os.Symlink(settingConfigFile, awsConfig())
	if err != nil {
		return err
	}
	return nil
}
