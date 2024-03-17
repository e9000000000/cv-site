package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func runAddNewUser() {
    fmt.Println("add new user ran")

    reader := bufio.NewReader(os.Stdin)

    fmt.Print("username: ")
    username, _ := reader.ReadString('\n')
    username = strings.Trim(username, " \n")

    fmt.Print("password: ")
    password, _ := reader.ReadString('\n')
    password = strings.Trim(password, " \n")

    fmt.Print("is admin? (t/f): ")
    isAdminString, _ := reader.ReadString('\n')
    isAdmin := strings.ToLower(strings.Trim(isAdminString, " \n")) == "t"

    _, err := addNewUser(username, password, isAdmin)
    if err != nil {
        fmt.Printf("ERROR: %s\n", err)
        os.Exit(1)
    } else {
        fmt.Println("user added")
    }
}
