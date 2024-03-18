package main

import (
	"fmt"
	"log"
)

type Post struct {
	Id     int
	Title  string
	Text   string
	Author *User
}

func getPosts() []Post {
	rows, err := db.Query("SELECT id, title, text, author_id FROM posts ORDER BY id DESC")
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for rows.Next() {
		var p Post
		var u User
		rows.Scan(&p.Id, &p.Title, &p.Text, &u.Id)
		err = u.LoadById()
		if err != nil {
			panic(err)
		} else {
			p.Author = &u
		}
		posts = append(posts, p)
	}

	return posts
}

func getPost(id int) *Post {
	rows, err := db.Query("SELECT id, title, text, author_id FROM posts WHERE id=?", id)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var p Post
	var u User
	if rows.Next() {
		rows.Scan(&p.Id, &p.Title, &p.Text, &u.Id)
		err = u.LoadById()
		if err != nil {
			panic(err)
		} else {
			p.Author = &u
		}
	} else {
		log.Printf("can't get post, no post with id %d\n", id)
	}

	return &p
}

func addPost(title string, text string, author *User) error {
	if author == nil {
		return fmt.Errorf("can't create post without author")
	}

	_, err := db.Exec("INSERT INTO posts (title, text, author_id) VALUES (?, ?, ?)", title, text, author.Id)
	if err != nil {
		panic(err)
	}

	return nil
}

func deletePost(id int, u *User) error {
	if u == nil || !u.IsAdmin {
		return fmt.Errorf("only admin can delete posts")
	}

	_, err := db.Exec("DELETE FROM posts WHERE id=?", id)
	if err != nil {
		panic(err)
	}

	return nil
}

func editPost(id int, title string, text string, u *User) error {
	if u == nil || !u.IsAdmin {
		return fmt.Errorf("only admin can edit posts")
	}

	_, err := db.Exec("UPDATE posts SET title=?, text=? WHERE id=?", title, text, id)
	if err != nil {
		panic(err)
	}

	return nil
}
