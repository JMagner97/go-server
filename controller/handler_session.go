package controller

import (
	"fmt"
	repository "go-server/Repository"
	utility "go-server/Utility"
	"go-server/data/response"
	helper "go-server/helper"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/sessions"
)

var users = map[string]string{
	"Mac":   "admin",
	"admin": "password",
}

var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

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
		session, _ := Store.Get(r, "session.id")
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

func LoginHandlerToken(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	helper.PanicIfError(err)
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	db := repository.GetDatabaseInstance()
	user_repo := repository.NewUserRepo(db)
	verified := repository.UserRepo.VerifyCredentials(user_repo, username, password)
	fmt.Println("Username:", username)
	fmt.Println("Password:", password)
	if verified {

		fmt.Println("Password inserita dall'utente:", password)
		tokenString, err := utility.GenerateJWTToken(username)
		repository.UserRepo.UpdateToken(user_repo, username, tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		var maps = make(map[string]string)
		maps["token"] = tokenString
		helper.WriteResponse(w, maps)
	} else {
		http.Error(w, "user is not found or invalid password", http.StatusNotFound)
		return
	}

}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	db := repository.GetDatabaseInstance()
	user_repo := repository.NewUserRepo(db)
	reqtoken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqtoken, "Bearer ")

	lentoken := len(splitToken)
	if lentoken > 1 {
		verified := repository.UserRepo.Logout(user_repo, splitToken[1])
		if verified {
			webRepo := response.WebResponse{
				Code:   200,
				Status: "ok",
			}
			helper.WriteResponse(w, webRepo)
		} else {
			webRepo := response.WebResponse{
				Code:   403,
				Status: "Error",
			}

			helper.WriteResponse(w, webRepo)
		}
	} else {
		webRepo := response.WebResponse{
			Code:   403,
			Status: "Error",
		}

		helper.WriteResponse(w, webRepo)
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}
	db := repository.GetDatabaseInstance()
	user_repo := repository.NewUserRepo(db)
	userValid := repository.UserRepo.VerifyUsername(user_repo, username)
	if !userValid {

		verified, error := repository.UserRepo.Signup(user_repo, username, password)
		if verified {
			webRepo := response.WebResponse{
				Code:   200,
				Status: "ok",
			}
			helper.WriteResponse(w, webRepo)
		} else {
			webRepo := response.WebResponse{
				Code:   403,
				Status: "Internal Server Error",
				Data:   error,
			}
			helper.WriteResponse(w, webRepo)
		}
	} else {
		webRepo := response.WebResponse{
			Code:   403,
			Status: "Username already exists",
		}
		helper.WriteResponse(w, webRepo)
	}
}
