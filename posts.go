package main

import (
    "log"
)

type Post struct {
    Id int
    Title string
    Text string
}


func getPosts() []Post {
    rows, err := db.Query("SELECT * FROM posts ORDER BY id DESC")
    if err != nil {
        panic(err)
    }

    var posts []Post
    for rows.Next() {
        var p Post
        rows.Scan(&p.Id, &p.Title, &p.Text)
        posts = append(posts, p)
    }

    rows.Close()
    return posts
}

func getPost(id int) *Post {
    rows, err := db.Query("SELECT * FROM posts WHERE id=?", id)
    defer rows.Close()
    if err != nil {
        panic(err)
    }

    var p Post
    if rows.Next() {
        rows.Scan(&p.Id, &p.Title, &p.Text)
    } else {
        log.Printf("can't get post, no post with id %d\n", id)
    }

    return &p
}

func addPost(title string, text string) {
    q, err := db.Prepare("INSERT INTO posts (title, text) VALUES (?, ?)")
    if err != nil {
        panic(err)
    } else {
        q.Exec(title, text)
    }
}

func deletePost(id int) {
    q, err := db.Prepare("DELETE FROM posts WHERE id=?")
    if err != nil {
        panic(err)
    } else {
        q.Exec(id)
    }
}

func editPost(id int, title string, text string) {
    q, err := db.Prepare("UPDATE posts SET title=?, text=? WHERE id=?")
    if err != nil {
        panic(err)
    } else {
        _, err := q.Exec(title, text, id)
        if err != nil {
            panic(err)
        }
    }
}
