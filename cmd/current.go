package cmd

import (
	"fmt"
	"github.com/cdimascio/aws-config-manager/cmd/color"
	"log"
	"os"
	"path"
)

func Current() error {
	fmt.Printf("%s %s\n", string(color.ColorGreen), current())
	return nil
}

func current() string {
	link, err := os.Readlink(awsCredentials())
	if err != nil {
		log.Fatal(err)
	}
	return stripTypeExt(path.Base(link))
}