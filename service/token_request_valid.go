package service

import (
	"fmt"
	En "go-server/EnumRole"
	repository "go-server/Repository"
	"go-server/Utility"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
)

func CheckToken(r *http.Request) bool {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) > 1 {
		reqToken = splitToken[1]

		// Retrieve the role from Redis using the token
		userRoleStr, err := Utility.RedisClient.Get(Utility.Ctx, reqToken).Result()
		if err == redis.Nil {
			fmt.Println("Token not found")
			return false
		} else if err != nil {
			fmt.Println("Error getting token from Redis:", err)
			return false
		}

		// Convert userRoleStr to an integer
		userRole, err := strconv.Atoi(userRoleStr)
		if err != nil {
			fmt.Println("Error converting role:", err)
			return false
		}

		// Check the user's role
		switch userRole {
		case En.Admin:
			fmt.Println("Hello Admin")
		case En.Professor:
			fmt.Println("Hello Professor")
		case En.Student:
			fmt.Println("Hello Student")
		default:
			fmt.Println("Unknown role")
			return false
		}

		// Verify the token is authenticated in the database
		userRepo := repository.NewUserRepo(repository.GetDatabaseInstance())
		return repository.UserRepo.VerifyIsAuthenticated(userRepo, reqToken)
	}
	return false
}
