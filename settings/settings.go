package settings

import (
	"os"
	"log"
	"bufio"
	"strings"
	//"reflect"
)

type Settings struct {
	settings    map[string]string
}

// method prepares config item to key and value.
// item is a string pattern {key} = {value} and it should be returned as key (string) and value (interface).
func (settings *Settings) prepareConfigItem(item string) (string, string) {
	data := strings.Split(item, "=")

	key := strings.Trim(data[0], " ")
	value := strings.Trim(data[1], " ")

	return key, value
}


// checks if given string is a comment.
func (settings *Settings) isComment(text string) bool {
	comments := []string{"/", "//", "#"}

	for _, element := range comments {
		if strings.HasPrefix(text, element) == true {
			return true
		}
	}

	return false
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
		item := scanner.Text()

		if settings.isComment(item) == false {
			key, value := settings.prepareConfigItem(scanner.Text())
			result[key] = value
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

// main function do use settings.
func SetSettings(path string) *Settings {
	settings := new(Settings)
	settings.Load(path)
	return settings
}