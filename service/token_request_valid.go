package service

import (
	repository "go-server/Repository"
	"net/http"
	"strings"
)

func CheckToken(token *http.Request) bool {
	reqtoken := token.Header.Get("Authorization")
	splitToken := strings.Split(reqtoken, "Bearer ")
	lentoken := len(splitToken)
	if lentoken > 1 {
		reqtoken = splitToken[1]
		userRepo := repository.NewUserRepo(repository.GetDatabaseInstance())
		isAuth := repository.UserRepo.VerifyIsAuthenticated(userRepo, reqtoken)
		return isAuth
	}
	return false
}
