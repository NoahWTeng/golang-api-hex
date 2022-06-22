package main

import (
	app "api-golang-hex/src"
	"log"
)

func main() {

	application, err := app.CreateNewApp()

	if err != nil {
		log.Fatal(err)
	}

	_ = application.Init()
}
