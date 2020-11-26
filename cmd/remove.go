package cmd

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"path"
)

func Remove(args cli.Args) error {
	if args == nil || args.Len() < 1 {
		return errors.New("remove <setting>")
	}

	setting := args.First()

	current := current()
	if setting == current {
		return errors.New(fmt.Sprintf("'%s' is in use. cannot remove it.", setting))
	}

	var errors []error
	configFile := path.Join(configDir(), setting+ExtConfig)
	credsFile := path.Join(configDir(), setting+ExtCredentials)

	err := os.Remove(credsFile)
	if !os.IsNotExist(err) {
		errors = append(errors, err)
	}

	err = os.Remove(configFile)
	if !os.IsNotExist(err) {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors[0]
	}
	return nil
}
