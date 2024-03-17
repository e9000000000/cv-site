package main

import (
    "net/http"
    "html/template"
    "strconv"
    "fmt"
    "log"
)

var templatesDir = "templates/"

func renderTemplate(w http.ResponseWriter, templateName string, data any) {
    t, err := template.ParseFiles(templatesDir + "base.html", templatesDir + templateName)
    if err != nil {
        log.Println(err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    err = t.Execute(w, data)

    if err != nil {
        log.Println(err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func handlePosts(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        renderTemplate(w, "posts.html", map[string][]Post{
            "Posts": getPosts(),
        })
    } else if r.Method == http.MethodDelete {
        err := r.ParseForm()
        if err != nil {
            http.Error(w, "can't parse get params", http.StatusBadRequest)
            return
        }
        id, _ := strconv.Atoi(r.Form["id"][0])

        deletePost(id)
        fmt.Fprint(w, "deleted")
    }
}

func handlePostsAdd(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        err := r.ParseForm()
        if err != nil {
            http.Error(w, "can't parse form", http.StatusBadRequest)
            return
        }
        data := r.PostForm
        
        addPost(data["title"][0], data["text"][0])
    }
    
    renderTemplate(w, "posts_add.html", nil)
}

func handlePostsEdit(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "can't parse form", http.StatusBadRequest)
        return
    }

    data := r.Form
    id, _ := strconv.Atoi(data["id"][0])

    if r.Method == http.MethodPost {
        editPost(id, data["title"][0], data["text"][0])
    }
    
    p := getPost(id)
    renderTemplate(w, "posts_edit.html", p)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "index.html", nil)
}

func handleMineswaper(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "mineswaper.html", nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        err := r.ParseForm()
        if err != nil {
            http.Error(w, fmt.Sprintf("can't parse form: %v", err), http.StatusBadRequest)
            return
        }

        data := r.PostForm
        username := data["username"][0]
        password := data["password"][0]

        err = loginUser(username, password)

        if err != nil {
            renderTemplate(w, "login.html", map[string]string{
                "Username": username,
                "Error": fmt.Sprintf("can't login: %v", err),
            })
            return
        }
    }

    renderTemplate(w, "login.html", nil)
}
