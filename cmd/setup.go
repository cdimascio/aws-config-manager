package cmd

import (
	"fmt"
	"log"
	"os"
	"path"
)

func Initialize() {
	ensureDir(configDir())
	cfSteps := setup(awsConfig(), ExtConfig)
	crSteps := setup(awsCredentials(), ExtCredentials)
	steps := append(cfSteps, crSteps...)

	if len(steps) > 0 {
		fmt.Println("You're all set up and good to go!")
		fmt.Println("Here's what changed:")
		for _, step := range steps {
			fmt.Println(step)
		}
	}
}

func ensureDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0600)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func setup(awsFile, suffix string) []string{
	var steps []string
	if fileExists(awsFile) {
		if !isSymlink(awsFile) {
			file := path.Join(configDir(), "default"+suffix)
			err := os.Rename(awsFile, file)
			if err != nil {
				log.Fatal(err)
			}
			steps = append(steps, fmt.Sprintf("moved %s to %s", awsFile, file))

			err = os.Symlink(file, awsFile)
			if err != nil {
				log.Fatal(err)
			}
			steps = append(steps, fmt.Sprintf("symlinked %s to %s", awsFile, file))
		}
	}
	return steps
}
