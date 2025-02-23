package src

import (
	"log"
	"os"
	"os/user"
)

var (
	programDir string
)

func InitDirectories() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	programDir = usr.HomeDir + "\\AppData\\Local\\Programs\\PluginBuilder"

	err = os.MkdirAll(programDir, 0777)
	if err != nil {
		Fatal("Error creating the program directory: " + err.Error())
	}
}

func getPluginDir(name string) string {
	dir := programDir + "\\plugins\\" + name

	err := os.MkdirAll(dir, 0777)
	if err != nil {
		Fatal("Error creating the " + name + " directory: " + err.Error())
	}

	return dir
}
