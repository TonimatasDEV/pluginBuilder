package src

import (
	"os"
	"os/exec"
	"pluginBuilder/src/utils"
)

func BuildMaven() error {
	// TODO: Maven build logic
	return nil
}

func BuildGradle(dir string) error { // Testing: build spigot 105298
	cmd := exec.Command(dir+"/gradlew", "build")
	cmd.Dir = dir
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	utils.Info("Build finish.")
	return nil
}
