package main

import (
    "net/http"
    "html/template"
    "strconv"
    "fmt"
    "log"
)

var templatesDir = "templates/"

type Page struct {
    User *User
    Content any
}

func renderTemplate(w http.ResponseWriter, r *http.Request, templateName string, data any) {
    t, err := template.ParseFiles(templatesDir + "base.html", templatesDir + templateName)
    if err != nil {
        log.Println(err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    cookies := r.Cookies()
    token := ""
    for _, c := range cookies {
        if c.Name == "token" {
            token = c.Value
            break
        }
    }
    user, _ := AuthUser(token)

    p := Page {
        User: user,
        Content: data,
    }

    err = t.Execute(w, p)

    if err != nil {
        log.Println(err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func handlePosts(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        renderTemplate(w, r, "posts.html", map[string][]Post{
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
    
    renderTemplate(w, r, "posts_add.html", nil)
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
    renderTemplate(w, r, "posts_edit.html", p)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, r, "index.html", nil)
}

func handleMineswaper(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, r, "mineswaper.html", nil)
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

        user, err := loginUser(username, password)
        if err != nil {
            renderTemplate(w, r, "login.html", map[string]string{
                "Username": username,
                "Error": fmt.Sprintf("can't login: %v", err),
            })
        } else {
            cookie := http.Cookie{
                Name: "token",
                Value: user.Token,
            }
            http.SetCookie(w, &cookie)
            http.Redirect(w, r, "/", http.StatusSeeOther)
        }
    } else {
        renderTemplate(w, r, "login.html", nil)
    }
}
