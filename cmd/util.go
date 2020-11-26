package cmd

import (
	"log"
	"os"
	"os/user"
	"regexp"
)

const ExtCredentials = ".credentials"
const ExtConfig = ".config"

func configDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}
	return usr.HomeDir+"/.aws_cred_man"
}

func awsCredentials() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}
	return usr.HomeDir+"/.aws/credentials"
}

func awsConfig() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}
	return usr.HomeDir+"/.aws/config"
}

func isSymlink(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		log.Fatal(err)
	}
	return fi.Mode() & os.ModeSymlink != 0
}

// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func hasSetting(name string) bool {
	settings, _ := list()
	for _, setting := range settings {
		if setting == name {
			return true
		}
	}
	return false
}

func isValidSettingName(name string) bool {
	reserved := []string{"config", "credentials"}
	for _, res := range reserved {
		if name == res {
			return false
		}
	}
	return true
}

func isValidType(name string) bool {
	reserved := []string{"config", "credentials"}
	for _, res := range reserved {
		if name == res {
			return true
		}
	}
	return false
}

var re = regexp.MustCompile(`(^.*)(\.config|\.credentials)$`)
func stripTypeExt(name string) string {
	return re.ReplaceAllString(name, "$1")
}