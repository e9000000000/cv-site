package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readLine(prompt string, r bufio.Reader) string {
	fmt.Print(prompt)
	result, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.Trim(result, " \n")
}

func runAddNewUser() {
	fmt.Println("add new user ran")

	reader := bufio.NewReader(os.Stdin)

	username := readLine("username: ")
	password := readLine("password: ")
	isAdmin := strings.ToLower(readLine("is admin? (t/f): ")) == "t"

	_, err := addNewUser(username, password, isAdmin)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	} else {
		fmt.Println("user added")
	}
}
