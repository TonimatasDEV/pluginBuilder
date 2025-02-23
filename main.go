package main

import (
	"fmt"
	"os"
	"os/signal"
	"pluginBuilder/src"
	"pluginBuilder/src/utils"
	"syscall"
)

func main() {
	stopSignal()

	src.InitDirectories()
	utils.InitLogger()
	src.InitPlugins()

	utils.Info("Plugin Builder initialized correctly.\nType \"?\" or \"help\" for more information.")

	src.InitCLI()
}

func stopSignal() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Print("\n")
		utils.Info("Stopping...")
		os.Exit(0)
	}()
}
