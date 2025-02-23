package src

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InitCLI() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		command := scanner.Text()

		commands := strings.Split(command, " ")
		args := commands[1:]

		switch commands[0] {
		case "exit":
			handleExit()
		case "help", "?":
			handleHelp()
		case "build":
			handleBuild(args)
		case "plugins":
			handlePlugins(args)
		default:
			Error("Unknown command. Use \"?\" or \"help\" for more information")
		}
	}
}

func handleExit() {
	Info("Exiting...")
	os.Exit(0)
}

func handleHelp() {
	Info("Commands: \n" +
		"exit    -> Exits the CLI.\n" +
		"help    -> Prints this help message.\n" +
		"build   -> Build a premium plugin. Ex: build <page|ex:spigot> <id|ex:1234>\n" +
		"plugins -> Prints all available plugins. Ex: plugins <page|ex:spigot>")
}

func handleBuild(args []string) {
	if len(args) < 2 {
		Error("Missing arguments. Should be \"build <page|ex:spigot> <id|ex:1234>\"")
		return
	}

	if args[0] != "spigot" {
		Error("For now, only exist: spigot")
		return
	}

	_, exist := Plugins[args[1]]
	if !exist {
		Error("Plugin \"" + args[1] + "\" not found")
		return
	}

	// TODO: Build logic
}

func handlePlugins(args []string) {
	if len(args) < 1 {
		Error("Missing arguments. Should be \"plugins <page|ex:spigot>\"")
		return
	}

	if args[0] != "spigot" {
		Error("For now, only exist: spigot")
		return
	}

	Info("Current spigot plugin list:")
	for id, data := range Plugins {
		fmt.Println(fmt.Sprintf("Plugin \"%s\":\n - Spigot: %s\n - GitHub: %s", id, data.Spigot, data.GitHub))
	}
}
