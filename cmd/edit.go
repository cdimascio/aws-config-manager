package cmd

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
)

func Edit(args cli.Args, fileType string) error {
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

	return openFileInEditor(file)
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
