package main

import (
	"fmt"
	"os"

	spud "github.com/idreaminteractive/spud"
)

func main() {
	// setup our logger
	// get our env config

	app := spud.NewApp("/admin",
		spud.WithPort(8080),
		spud.WithHost("0.0.0.0"),
		spud.WithSQLite("./spud.db"),
		spud.WithMode(spud.Production),
	)

	// setup an errgroup to run and ensure we can close
	if err := app.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
