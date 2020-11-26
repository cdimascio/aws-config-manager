package cmd

import (
	"fmt"
	"github.com/cdimascio/aws-config-manager/cmd/color"
	"io/ioutil"
)

func List() error {
	files, err := list()
	if err != nil {
		return err
	}

	current := current()
	for _, setting := range files {
		if setting == current {
			fmt.Printf("%s* %s\n", color.ColorGreen, setting)
		} else {
			fmt.Printf("%s  %s\n", color.ColorReset, setting)
		}
	}
	return nil
}

func list() ([]string, error) {
	files, err := ioutil.ReadDir(configDir())
	if err != nil {
		return nil, err
	}

	var settings []string
	var m = make(map[string]string)
	for _, f := range files {
		fileName := f.Name()
		setting := stripTypeExt(fileName)
		_, ok := m[setting]
		if !ok {
			settings = append(settings, setting)
		}
		m[setting] = fileName
	}
	return settings, nil
}
