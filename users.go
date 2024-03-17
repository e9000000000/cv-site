package main

import (
    "log"
    "fmt"
    "crypto/sha1"
    "encoding/hex"
)

func hashString(password string) string {
    h := sha1.New()
    h.Write([]byte(password))
    return hex.EncodeToString(h.Sum(nil))
}

type User struct {
    Id int
    Token string
    Name string
    PwHash string
    IsAdmin bool
}


func (u *User) SetTokenFromNameAndPwHash() {
    u.Token = hashString(fmt.Sprintf("%s:%s", u.Name, u.PwHash))
}

func (u *User) Save() error {
    if u.Id == 0 {
        res, err := db.Exec("INSERT INTO users (username, password_hash, token, is_admin) VALUES (?, ?, ?, ?)", u.Name, u.PwHash, u.Token, u.IsAdmin)
        if err != nil {
            return err
        }

        userId, err := res.LastInsertId()
        u.Id = int(userId)
        if err != nil {
            return err
        }
    } else {
        _, err := db.Exec("UPDATE users SET username=?, password_hash=?, token=?, is_admin=?", u.Name, u.PwHash, u.Token, u.IsAdmin)
        return err
    }
    return nil
}

// Token should be assigned before call
func (u *User) LoadByToken() error {
    rows, err := db.Query("SELECT id, username, password_hash, is_admin FROM users WHERE token=?", u.Token)
    defer rows.Close()
    if err != nil {
        return err
    }

    if rows.Next() {
        rows.Scan(&u.Id, &u.Name, &u.PwHash, &u.IsAdmin)
    } else {
        return fmt.Errorf("can't find user by token: %s", u.Token)
    }

    return nil
}

// Id should be assigned before call
func (u *User) LoadById() error {
    rows, err := db.Query("SELECT token, username, password_hash, is_admin FROM users WHERE id=?", u.Id)
    defer rows.Close()
    if err != nil {
        return err
    }

    if rows.Next() {
        rows.Scan(&u.Token, &u.Name, &u.PwHash, &u.IsAdmin)
    } else {
        return fmt.Errorf("can't find user by id: %s", u.Id)
    }

    return nil
}

func AuthUser(token string) (*User, error) {
    if token == "" {
        return nil, fmt.Errorf("no token")
    }

    u := User {Token: token}
    err := u.LoadByToken()
    if err != nil {
        return nil, err
    }
    return &u, nil
}

func loginUser(username string, password string) (*User, error) {
    log.Printf("attempt to login with username=%s\n", username)

    pwHash := hashString(password)
    rows, err := db.Query("SELECT id FROM users WHERE username=? AND password_hash=?", username, pwHash)
    defer rows.Close()
    if err != nil {
        return nil, err
    }

    var u User
    if rows.Next() {
        rows.Scan(&u.Id)
        err = u.LoadById()
        return &u, err
    } else {
        return nil, fmt.Errorf("wrong username or password")
    }
}

func addNewUser(username string, password string, isAdmin bool) (*User, error) {
    pwHash := hashString(password)
    u := User {
        Name: username,
        PwHash: pwHash,
        IsAdmin: isAdmin,
    }
    u.SetTokenFromNameAndPwHash()
    err := u.Save()
    return &u, err
}
