package main

import (
	"fmt"
	"os"
	"os/signal"
	"pluginBuilder/src"
	"syscall"
)

func main() {
	stopSignal()

	src.InitDirectories()
	src.InitLogger()
	src.InitPlugins()

	src.Info("Plugin Builder initialized correctly.\nType \"?\" or \"help\" for more information.")

	src.InitCLI()
}

func stopSignal() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Print("\n")
		src.Info("Stopping...")
		os.Exit(0)
	}()
}
