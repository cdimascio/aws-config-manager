package cmd

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"path"
)

func Edit(args cli.Args) error {
	if args == nil || args.Len() < 1 {
		return errors.New("edit [config|credentials] [<setting>]")
	}

	fileType := args.First()
	if !isValidType(args.First()) {
		return errors.New("edit [config|credentials] [<setting>]")
	}

	var setting string
	if args.Len() > 1 {
		setting = args.Get(1)
		if !hasSetting(setting) {
			return errors.New(fmt.Sprintf("setting '%s' does not exist", setting))
		}
	} else {
		setting = current()
	}

	if !isValidSettingName(args.Get(1)) {
		return errors.New("invalid setting name. cannot be 'config' or 'credential'")
	}

	var file string
	if fileType == "config" {
		file = path.Join(configDir(), setting+ExtConfig)
	} else {
		file = path.Join(configDir(), setting+ExtCredentials)
	}
	err := openFileInEditor(file)
	return err
}

const DefaultEditor = "vi"

func openFileInEditor(filename string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = DefaultEditor
	}

	executable, err := exec.LookPath(editor)
	if err != nil {
		return err
	}

	cmd := exec.Command(executable, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}