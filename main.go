package main

import (
	"fmt"
	"github.com/slawek87/gophe/settings"
)

func main() {
	mySettings := settings.SetSettings("./settings/default.cfg")
	test := mySettings.Get("DATABASE_DRIVER")
	fmt.Println(test)
}