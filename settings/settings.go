package settings

import (
	"os"
	"log"
	"bufio"
	"strings"
	//"reflect"
)

type Validation struct {}

// checks if current key and value are valid settings - cannot be empty strings.
func (validation *Validation) isValid(key string, value string) bool {
	if key == "" || value == "" {
		return false
	}
	return true
}

type Comments struct {}

// checks if given string is a comment - starts with: /, //, #.
func (comments *Comments) isComment(text string) bool {
	commentsPattern := []string{"/", "//", "#"}

	for _, element := range commentsPattern {
		if strings.HasPrefix(text, element) == true {
			return true
		}
	}

	return false
}

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

// method is reading config and returns go lang object in pattern map[string]string.
func (settings *Settings) read(path string) {
	result := make(map[string]string)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		configLine := scanner.Text()

		if settings.comments.isComment(configLine) == false {
			key, value := settings.prepareConfigItem(configLine)

			// create config configLine only when key and value are valid.
			if settings.validation.isValid(key, value) {
				result[key] = value
			}
		}
	}
	settings.settings = result
}

//Load config
func (settings *Settings) Load(path string) map[string]string {
	settings.read(path)
	return settings.settings
}

// returns config value for given key. Returns only string values.
func (settings *Settings) Get(key string) string {
	return settings.settings[key]
}

// main function to use settings.
func SetSettings(path string) *Settings {
	settings := new(Settings)
	settings.Load(path)
	return settings
}