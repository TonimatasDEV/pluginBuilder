package src

import (
	"log"
	"os"
	"os/user"
	"pluginBuilder/src/utils"
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

	err = os.MkdirAll(programDir, os.ModePerm)
	if err != nil {
		utils.Fatal("Error creating the program directory: " + err.Error())
	}
}

func getPluginDir(page, id string) string {
	dir := programDir + "\\plugins\\" + page + "\\" + id

	err := os.RemoveAll(dir)
	if err != nil {
		utils.Fatal("Error deleting old directory of: " + id + err.Error())
	}

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		utils.Fatal("Error creating the " + id + " directory: " + err.Error())
	}

	return dir
}
