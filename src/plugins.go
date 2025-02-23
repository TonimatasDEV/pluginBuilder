package src

import (
	"encoding/json"
	"io"
	"net/http"
	"pluginBuilder/src/utils"
)

const pluginGist = "https://gist.githubusercontent.com/TonimatasDEV/1e47502a243f938200a6354bdde65c4b/raw"

var Plugins map[string]PluginData

type PluginData struct {
	Spigot string `json:"spigot"`
	GitHub string `json:"github"`
}

func InitPlugins() {
	utils.Info("Loading plugins...")
	pluginListRequest, err := http.Get(pluginGist)

	if err != nil {
		utils.Fatal("Error retrieving the plugin list.")
	}

	defer utils.CloseReadCloser(pluginListRequest.Body)

	pluginListJson, err := io.ReadAll(pluginListRequest.Body)
	if err != nil {
		utils.Fatal("Error reading the plugin list.")
	}

	err = json.Unmarshal(pluginListJson, &Plugins)
	if err != nil {
		utils.Fatal("Error parsing the plugin list." + err.Error())
	}

	utils.Info("Plugins loaded successfully.")
}

func build(page, id string, data PluginData) {
	dir := getPluginDir(page, id)

	err := utils.DownloadFile(data.GitHub+"/archive/refs/heads/main.zip", dir+"\\plugin.zip")
	if err != nil {
		utils.Error("Error downloading the plugin repository. " + err.Error())
		return
	}

	err = utils.UnzipFile(dir+"\\plugin.zip", dir)
	if err != nil {
		utils.Error("Error unzipping the plugin. " + err.Error())
		return
	}

	// TODO: Build logic
}
