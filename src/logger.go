package src

import (
	"log"
	"os"
)

const (
	LightGray   = "\033[0;37m"
	Red         = "\033[1;31m"
	StrongGreen = "\033[1;32m"
	Yellow      = "\033[1;33m"
	Blue        = "\033[1;34m"
	Purple      = "\033[1;35m"
	Reset       = "\033[0;57m"
)

func InitLogger() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lmsgprefix | log.Ltime)
}

func Info(message string) {
	setPrefix(Blue, "INFO")
	log.Println(message + Reset)
}

func Warn(message string) {
	setPrefix(Yellow, "WARN")
	log.Println(message + Reset)
}

func Error(message string) {
	setPrefix(Red, "ERROR")
	log.Println(message + Reset)
}

func Fatal(message string) {
	setPrefix(Purple, "CRASH")
	log.Fatalln(message + Reset)
}

func setPrefix(color, level string) {
	log.SetPrefix(StrongGreen + "[PluginBuilder] " + color + "[" + level + "]: " + LightGray)
}
