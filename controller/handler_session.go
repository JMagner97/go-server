package controller

import (
	"fmt"
	"go-server/helper"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/sessions"
)

var users = map[string]string{
	"Mac":   "admin",
	"admin": "password",
}

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	helper.PanicIfError(err)
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	fmt.Println("Username:", username)
	fmt.Println("Password:", password)

	if originalPassword, ok := users[username]; ok {
		fmt.Println("Password inserita dall'utente:", password)
		fmt.Println("Password memorizzata nell'oggetto users:", originalPassword)
		session, _ := store.Get(r, "session.id")
		if password == originalPassword {
			session.Values["authenticated"] = true
			session.Save(r, w)

		} else {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "user is not found", http.StatusNotFound)
		return
	}
	w.Write([]byte("Logged in Successfully"))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	session.Values["authenticated"] = false
	session.Save(r, w)
	w.Write([]byte("Logout"))
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	if (session.Values["authenticated"] != nil) && session.Values["authenticated"] != false {
		w.Write(([]byte(time.Now().String())))
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}
