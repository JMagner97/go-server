package controller

import (
	"fmt"
	En "go-server/EnumRole"
	repository "go-server/Repository"
	utility "go-server/Utility"
	helper "go-server/helper"
	"net/http"
	"strings"
)

var users = map[string]string{
	"Mac":   "admin",
	"admin": "password",
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
	switch verifiedRole {
	case En.Admin, En.Professor, En.Student:
		fmt.Println("Logged in as", verifiedRole)
		tokenString, _ := utility.GenerateJWTToken(username)
		repository.UserRepo.UpdateToken(user_repo, username, tokenString)

		err := utility.RedisClient.Set(utility.Ctx, tokenString, verifiedRole, 0).Err()
		if err != nil {
			http.Error(w, "Error saving session", http.StatusInternalServerError)
			return
		}

		webResponse := helper.WebResponse{
			Status: "ok",
			Data: map[string]string{
				"token": tokenString,
			},
		}
		helper.WriteResponse(w, webResponse, http.StatusOK)

	default:
		fmt.Println("Unknown role")
		http.Error(w, "user is not found or invalid password", http.StatusUnauthorized)
		return
	}
}
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")

	if len(splitToken) > 1 {
		token := splitToken[1]

		// Delete the token from Redis
		err := utility.RedisClient.Del(utility.Ctx, token).Err()
		if err != nil {
			fmt.Println("Error deleting token from Redis:", err)
			webRepo := helper.WebResponse{
				Status: "Error",
			}
			helper.WriteResponse(w, webRepo, http.StatusInternalServerError)
			return
		}

		webRepo := helper.WebResponse{
			Status: "ok",
		}
		helper.WriteResponse(w, webRepo, http.StatusOK)
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
