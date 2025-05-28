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
	spud.HE()
	fmt.Println("Hello, World!")
	os.Exit(0)
}
