package main

import (
	"log"
	"net/http"
)

func runServer() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/posts/add", handlePostsAdd)
	http.HandleFunc("/posts/edit", handlePostsEdit)
	http.HandleFunc("/posts", handlePosts)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/logout", handleLogout)
	http.HandleFunc("/mineswaper", handleMineswaper)

	log.Println("server started")
	panic(http.ListenAndServe(":8000", nil))
}
