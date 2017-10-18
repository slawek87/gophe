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

// Method to transform text to key and value to golang object.
func (settings *Settings) prepareConfigItem(configLine string) (string, string) {
	configData := strings.Split(configLine, "=")

	key := strings.Trim(configData[0], " ")
	value := strings.Join(configData[1:len(configData)], "=")
	value = strings.Trim(value, " ")

	return key, value
}

// Method takes configLine and transform it to golang object.
// Returns only valid objects - with valid key and value.
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

	return settingsFile
}

// This method run all operations needed to transform text config to golang object.
func (settings *Settings) ProcessingConfigFile(configPath string) map[string]string {
	settings.settings = make(map[string]string)
	settingsFile := settings.read(configPath)
	defer settingsFile.Close()

	scanner := bufio.NewScanner(settingsFile)
	for scanner.Scan() {
		configLine := scanner.Text()
		settings.mapSettings(configLine)
	}

	return settings.settings
}

// Method returns config items with given key.
func (settings *Settings) Get(key string) string {
	return settings.settings[key]
}


// This is the main function to setup our config file to Settings object.
func SetSettings(configPath string) *Settings {
	settings := new(Settings)
	settings.ProcessingConfigFile(configPath)
	return settings
}