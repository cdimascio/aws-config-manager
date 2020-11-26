package cmd

import (
	"errors"
	"log"
	"os"
	"os/user"
	"path"
	"regexp"
	"strings"
)

const ExtCredentials = ".credentials"
const ExtConfig = ".config"

func configDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir + "/.aws_cred_man"
}

func awsCredentials() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir + "/.aws/credentials"
}

func awsConfig() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir + "/.aws/config"
}

func isSymlink(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		log.Fatal(err)
	}
	return fi.Mode()&os.ModeSymlink != 0
}

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

var re = regexp.MustCompile(`(^.*)(\.config|\.credentials)$`)

func stripTypeExt(name string) string {
	return re.ReplaceAllString(name, "$1")
}

func pathFromType(setting, typ string) (string, error) {
	var file string
	if strings.HasPrefix(typ, "conf") {
		file = path.Join(configDir(), setting+ExtConfig)
	} else if strings.HasPrefix(typ, "cred") {
		file = path.Join(configDir(), setting+ExtCredentials)
	} else {
		return "", errors.New("type must be one of conf[ig] or cred[entials]")
	}
	return file, nil
}
