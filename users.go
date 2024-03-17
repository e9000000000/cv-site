package main

import (
    "log"
    "fmt"
)

func loginUser(username string, password string) error {
    log.Printf("attempt to login with username=%s\n", username)
    return fmt.Errorf("any logins are not allowed for now")

}
