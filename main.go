package main

import (
    "log"
    "net/http"
)


func main() {
    initDb()

    http.HandleFunc("/", handleRoot)
    http.HandleFunc("/posts/add", handlePostsAdd)
    http.HandleFunc("/posts/edit", handlePostsEdit)
    http.HandleFunc("/posts", handlePosts)
    http.HandleFunc("/login", handleLogin)
    http.HandleFunc("/mineswaper", handleMineswaper)

    log.Println("started")
    panic(http.ListenAndServe(":8000", nil))
}
