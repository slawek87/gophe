package settings

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSuccessSetSettings(t *testing.T) {
	mySettings := SetSettings("./settings_test.cfg")
	assert.Equal(t, mySettings.Get("DATABASE_DRIVER"), "mysql", "Database driver should be set to mysql.")
	assert.Equal(t, mySettings.Get("DATABASE_ARGS"), "root:abcdefg123456@tcp(localhost:3306)/mydb?charset=utf8&parseTime=True&loc=Local", "Should be equal.")
}

func TestFailSetSettings(t *testing.T) {
	mySettings := SetSettings("./settings_test.cfg")
	assert.NotEqual(t, mySettings.Get("DATABASE_DRIVER"), "mysql2", "Database driver should be different then mysql.")
	assert.NotEqual(t, mySettings.Get("DATABASE_ARGS"), "f8&parseTime=True&loc=Local", "Should not be equal.")
}