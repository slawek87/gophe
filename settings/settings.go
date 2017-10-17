package settings

import (
	"os"
	"log"
	"bufio"
	"strings"
	//"reflect"
)

type Validation struct {}

// checks if current key and value are valid settings.
func (validation *Validation) isValid(key string, value string) bool {
	if key == "" || value == "" {
		return false
	}
	return true
}

type Comments struct {}

// checks if given string is a comment.
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

// method prepares config item to key and value.
// item is a string pattern {key} = {value} and it should be returned as key (string) and value (interface).
func (settings *Settings) prepareConfigItem(item string) (string, string) {
	data := strings.Split(item, "=")

	key := strings.Trim(data[0], " ")
	value := strings.Join(data[1:len(data)], "=")
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
		item := scanner.Text()

		if settings.comments.isComment(item) == false {
			key, value := settings.prepareConfigItem(scanner.Text())
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