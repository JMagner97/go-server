package controller

import (
	"fmt"
	En "go-server/EnumRole"
	repository "go-server/Repository"
	helper "go-server/helper"
	utility "go-server/utility"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
)

var users = map[string]string{
	"Mac":   "admin",
	"admin": "password",
}

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
		session, _ := utility.Store.Get(r, "session.id")
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
	verifiedRole := repository.UserRepo.VerifyCredentials(user_repo, username, password)
	fmt.Println("Username:", username)
	fmt.Println("Password:", password)
	if verifiedRole != En.Unknown {

		fmt.Println("Password inserita dall'utente:", password)
		tokenString, _ := utility.GenerateJWTToken(username)
		repository.UserRepo.UpdateToken(user_repo, username, tokenString)
		session, _ := utility.Store.New(r, "idsession")
		session.Options = &sessions.Options{
			MaxAge:   3600, // 1 ora
			HttpOnly: true, // Più sicuro
		}
		session.Values = map[interface{}]interface{}{} // Pulisce i vecchi valori
		session.Values["role"] = verifiedRole          // Salva solo il ruolo attuale
		session.Values["token"] = tokenString          // Salva anche il token
		errx := session.Save(r, w)
		if errx != nil {
			http.Error(w, errx.Error(), http.StatusInternalServerError)
			return
		}

		webResponse := helper.WebResponse{
			Status: "ok",
			Data: map[string]string{
				"token": tokenString,
			},
		}
		helper.WriteResponse(w, webResponse, http.StatusOK)

	} else {
		http.Error(w, "user is not found or invalid password", http.StatusUnauthorized)
		return
	}

}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	//db := repository.GetDatabaseInstance()
	//user_repo := repository.NewUserRepo(db)
	reqtoken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqtoken, "Bearer ")

	lentoken := len(splitToken)
	if lentoken > 1 {
		session, _ := utility.Store.Get(r, "idsession")
		session.Values = map[interface{}]interface{}{}
		session.Save(r, w)
		//verified := repository.UserRepo.Logout(user_repo, splitToken[1])
		// if verified {
		// 	webRepo := helper.WebResponse{
		// 		Status: "ok",
		// 	}
		// 	helper.WriteResponse(w, webRepo, http.StatusOK)
		// } else {
		// 	webRepo := helper.WebResponse{
		// 		Status: "Error",
		// 	}

		// 	helper.WriteResponse(w, webRepo, http.StatusBadRequest)
		// }
	} else {
		webRepo := helper.WebResponse{
			Status: "Error",
		}

		helper.WriteResponse(w, webRepo, http.StatusBadRequest)
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
			webRepo := helper.WebResponse{

				Status: "ok",
			}
			helper.WriteResponse(w, webRepo, http.StatusCreated)
		} else {
			webRepo := helper.WebResponse{
				Status: "Internal Server Error",
				Data:   error,
			}
			helper.WriteResponse(w, webRepo, http.StatusBadRequest)
		}
	} else {
		webRepo := helper.WebResponse{
			Status: "Username already exists",
		}
		helper.WriteResponse(w, webRepo, http.StatusConflict)
	}
}
