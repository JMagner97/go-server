package service

import (
	"fmt"
	En "go-server/EnumRole"
	repository "go-server/Repository"
	"go-server/utility"
	"net/http"
	"strings"
)

func CheckToken(r *http.Request) bool {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) > 1 {
		reqToken = splitToken[1]

		session, _ := utility.Store.Get(r, "idsession")
		if storedToken, ok := session.Values["token"].(string); ok && storedToken == reqToken {
			if storedRole, ok := session.Values["role"].(int); ok { // Assicurati che il tipo sia corretto
				switch storedRole {
				case En.Admin:
					fmt.Println("Ciao Admin")
				case En.Professor:
					fmt.Println("Ciao Professor")
				case En.Student:
					fmt.Println("Ciao Student")
				default:
					fmt.Println("Unknown role")
					return false
				}

				userRepo := repository.NewUserRepo(repository.GetDatabaseInstance())
				return repository.UserRepo.VerifyIsAuthenticated(userRepo, reqToken)
			}
		}
	}
	return false
}
