/*
Settings module is responsible for create Settings from config file.
In this module all things named as config means file.cfg - text data. All things named settings mean golang object.

* all values fetched from config are stored in Settings as string object.

Example of use:

mySettings := SetSettings("mySettings.cfg") // config include DEBUG = true

fmt.println(mySettings.Get("DEBUG")) // print "true" string value

mySettings.Set(DEBUG, false) // set DEBUG to false boolean value.

 */
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

// Method transforms text to golang objects.
// We need key and value to create settings item.
func (settings *Settings) prepareConfigItem(configLine string) (string, string) {
	configData := strings.Split(configLine, "=")

	key := strings.Trim(configData[0], " ")
	value := strings.Join(configData[1:len(configData)], "=")
	value = strings.Trim(value, " ")

	return key, value
}

// Method takes configLine and transforms it to golang object.
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

// This method run all operations needed to transform text config to golang object - settings.
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

// Method returns settings item with given key.
func (settings *Settings) Get(key string) string {
	return settings.settings[key]
}

// Method gives opportunity to set settings item directly from code.
func (settings *Settings) Set(key string, value string) {
	settings.settings[key] = value
}

// This is the main function to setup our config file to Settings object.
func SetSettings(configPath string) *Settings {
	settings := new(Settings)
	settings.ProcessingConfigFile(configPath)
	return settings
}