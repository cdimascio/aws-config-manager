package cmd

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"path"
	"strings"
)

func Edit(args cli.Args, fileType string) error {
	var setting string
	if args.Len() >= 1 {
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
	if strings.HasPrefix(fileType, "conf") {
		file = path.Join(configDir(), setting+ExtConfig)
	} else if strings.HasPrefix(fileType, "cred"){
		file = path.Join(configDir(), setting+ExtCredentials)
	} else {
		return errors.New("type must be one of conf[ig] or cred[entials]")
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