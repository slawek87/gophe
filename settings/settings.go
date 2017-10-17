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

// method prepares config item like key and value fot that key.
// configLine is a string with pattern {key} = {value}
func (settings *Settings) prepareConfigItem(configLine string) (string, string) {
	configData := strings.Split(configLine, "=")

	key := strings.Trim(configData[0], " ")
	value := strings.Join(configData[1:len(configData)], "=")
	value = strings.Trim(value, " ")

	return key, value
}

func (settings *Settings) mapSettings(configLine string) {
	if settings.comments.isComment(configLine) == false {
		key, value := settings.prepareConfigItem(configLine)

		// create config configLine only when key and value are valid.
		if settings.validation.isValid(key, value) {
			settings.settings[key] = value
		}
	}
}

// reads settings data from given path.
func (settings *Settings) read(path string) *os.File {
	settingsFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer settingsFile.Close()

	return settingsFile
}

// Process config
func (settings *Settings) Process(path string) map[string]string {
	settings.settings = make(map[string]string)
	settingsFile := settings.read(path)

	scanner := bufio.NewScanner(settingsFile)
	for scanner.Scan() {
		configLine := scanner.Text()
		settings.mapSettings(configLine)
	}

	return settings.settings
}

// returns config value for given key. Returns only string values.
func (settings *Settings) Get(key string) string {
	return settings.settings[key]
}

// main function to use settings.
func SetSettings(path string) *Settings {
	settings := new(Settings)
	settings.Process(path)
	return settings
}