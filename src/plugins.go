package src

import (
	"encoding/json"
	"io"
	"net/http"
)

const pluginGist = "https://gist.githubusercontent.com/TonimatasDEV/1e47502a243f938200a6354bdde65c4b/raw"

var Plugins map[string]PluginData

type PluginData struct {
	Spigot string `json:"spigot"`
	GitHub string `json:"github"`
}

func InitPlugins() {
	Info("Loading plugins...")
	pluginListRequest, err := http.Get(pluginGist)

	if err != nil {
		Fatal("Error retrieving the plugin list.")
	}

	pluginListJson, err := io.ReadAll(pluginListRequest.Body)
	if err != nil {
		Fatal("Error reading the plugin list.")
	}

	err = json.Unmarshal(pluginListJson, &Plugins)
	if err != nil {
		Fatal("Error parsing the plugin list." + err.Error())
	}

	Info("Plugins loaded successfully.")
}

func Download() {

}
