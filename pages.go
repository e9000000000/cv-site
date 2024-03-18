package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templatesDir = "templates/"

type Page struct {
	User    *User
	Content any
}

func makeBasePage(r *http.Request) *Page {
	cookies := r.Cookies()
	token := ""
	for _, c := range cookies {
		if c.Name == "token" {
			token = c.Value
			break
		}
	}
	user, _ := AuthUser(token)

	return &Page{
		User:    user,
		Content: nil,
	}
}

func renderTemplate(w http.ResponseWriter, templateName string, p *Page) {
	t, err := template.ParseFiles(templatesDir+"base.html", templatesDir+templateName)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = t.Execute(w, p)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func handlePosts(w http.ResponseWriter, r *http.Request) {
	p := makeBasePage(r)

	if r.Method == http.MethodGet {
		p.Content = getPosts()
		renderTemplate(w, "posts.html", p)
	} else if r.Method == http.MethodDelete {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "can't parse get params", http.StatusBadRequest)
			return
		}
		id, _ := strconv.Atoi(r.Form["id"][0])

		err = deletePost(id, p.User)
		if err != nil {
			http.Error(w, fmt.Sprintf("ERROR: %v", err), http.StatusInternalServerError)
		} else {
			fmt.Fprint(w, "deleted")
		}
	}
}

func handlePostsAdd(w http.ResponseWriter, r *http.Request) {
	p := makeBasePage(r)

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "can't parse form", http.StatusBadRequest)
			return
		}
		data := r.PostForm

		err = addPost(data["title"][0], data["text"][0], p.User)
		if err != nil {
			http.Error(w, fmt.Sprintf("ERROR: %v", err), http.StatusInternalServerError)
			return
		}
	}

	renderTemplate(w, "posts_add.html", p)
}

func handlePostsEdit(w http.ResponseWriter, r *http.Request) {
	p := makeBasePage(r)

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "can't parse form", http.StatusBadRequest)
		return
	}

	data := r.Form
	id, _ := strconv.Atoi(data["id"][0])

	if r.Method == http.MethodPost {
		err = editPost(id, data["title"][0], data["text"][0], p.User)
		if err != nil {
			http.Error(w, fmt.Sprintf("ERROR: %v", err), http.StatusInternalServerError)
			return
		}
	}

	p.Content = getPost(id)
	renderTemplate(w, "posts_edit.html", p)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	p := makeBasePage(r)
	renderTemplate(w, "index.html", p)
}

func handleMineswaper(w http.ResponseWriter, r *http.Request) {
	p := makeBasePage(r)
	renderTemplate(w, "mineswaper.html", p)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	p := makeBasePage(r)

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("can't parse form: %v", err), http.StatusBadRequest)
			return
		}

		data := r.PostForm
		username := data["username"][0]
		password := data["password"][0]

		user, err := loginUser(username, password)
		if err != nil {
			p.Content = map[string]string{
				"Username": username,
				"Error":    fmt.Sprintf("can't login: %v", err),
			}
		} else {
			cookie := http.Cookie{
				Name:  "token",
				Value: user.Token,
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	renderTemplate(w, "login.html", p)
}
