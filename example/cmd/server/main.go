package main

import (
	"fmt"
	"os"

	spud "github.com/idreaminteractive/spud"
)

func main() {
	// setup our db
	// setup our logger
	// get our env config
	app := spud.NewApp("admin", "db.sqlite")
	app.Port = 8080
	app.Host = "0.0.0.0"

	// setup an errgroup to run and ensure we can close
	if err := app.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
