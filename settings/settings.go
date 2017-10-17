package settings

import (
	"os"
	"log"
	"bufio"
	"strings"
)

type Settings struct {
	settings   	 	map[string]string
	comments 		Comments
	validation  	Validation
}

// We use that method to transform text to key and value objects.
func (settings *Settings) prepareConfigItem(configLine string) (string, string) {
	configData := strings.Split(configLine, "=")

	key := strings.Trim(configData[0], " ")
	value := strings.Join(configData[1:len(configData)], "=")
	value = strings.Trim(value, " ")

	return key, value
}

// We use that method to map settings from given configLine.
func (settings *Settings) mapSettings(configLine string) {
	if settings.comments.isComment(configLine) == false {
		key, value := settings.prepareConfigItem(configLine)

		// create settings item only when key and value are valid.
		if settings.validation.isValid(key, value) {
			settings.settings[key] = value
		}
	}
}

// Simple read file method.
func (settings *Settings) read(path string) *os.File {
	settingsFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer settingsFile.Close()

	return settingsFile
}

// We use this method to transform config text data to settings object.
func (settings *Settings) ProcessingConfigFile(configPath string) map[string]string {
	settings.settings = make(map[string]string)
	settingsFile := settings.read(configPath)

	scanner := bufio.NewScanner(settingsFile)
	for scanner.Scan() {
		configLine := scanner.Text()
		settings.mapSettings(configLine)
	}

	return settings.settings
}

// Use this method to get value fo given settings item.
func (settings *Settings) Get(key string) string {
	return settings.settings[key]
}


// main function to initialize settings object.
func SetSettings(configPath string) *Settings {
	settings := new(Settings)
	settings.ProcessingConfigFile(configPath)
	return settings
}