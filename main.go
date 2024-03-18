package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	initDb()

	var action string
	flag.StringVar(&action, "action", "", "what should be runing? 'server' or 'newuser'")
	flag.Parse()

	switch action {
	case "server":
		runServer()
	case "newuser":
		runAddNewUser()
	default:
		fmt.Printf("unknown action: %s\n", action)
		os.Exit(1)
	}
}
