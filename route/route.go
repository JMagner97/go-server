package route

import (
	"fmt"
	repository "go-server/Repository" // Renamed to lowercase
	"go-server/controller"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func wrapHandler(handler http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		handler(w, r)
	}
}

func RoleMiddleware(handler httprouter.Handle, allowedRole string, skipTokenCheck bool) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if skipTokenCheck {
			handler(w, r, ps)
			return
		}

		appRoleStr := os.Getenv("APP_ROLE")
		fmt.Println("App Role:", appRoleStr)
		appRole, err := strconv.Atoi(appRoleStr)
		if err != nil {
			http.Error(w, "Invalid APP_ROLE", http.StatusInternalServerError)
			return
		}
		reqToken := r.Header.Get("Authorization")
		if reqToken == "" {
			http.Error(w, "Required token", http.StatusForbidden)
			return
		}
		splitToken := strings.Split(reqToken, "Bearer ")
		if splitToken[1] == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		userRole, err := checkUserRole(splitToken[1])
		if err != nil {
			http.Error(w, "This token is not for your role", http.StatusInternalServerError)
			return
		}

		if userRole == appRole {
			handler(w, r, ps)
			return
		}

		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}

func checkUserRole(storedToken string) (int, error) {
	db := repository.GetDatabaseInstance()
	user_repo := repository.NewUserRepo(db)
	return user_repo.CheckUserRole(storedToken)
}

func NewRouter(studenController *controller.StudentController, lecturesController *controller.LectureController, enrollmentController *controller.LectureStudentController, professorController *controller.ProfessorController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome Home")
	})

	router.GET("/api/student", RoleMiddleware(studenController.FindAll, "professor", false))
	router.GET("/api/student/:email", RoleMiddleware(studenController.FindById, "student", false))
	router.POST("/api/student", RoleMiddleware(studenController.Create, "student", true))
	router.POST("/api/professor", RoleMiddleware(wrapHandler(professorController.Create), "professor", true))
	router.PATCH("/api/student/:email", RoleMiddleware(studenController.Update, "student", false))
	router.DELETE("/api/student/:email", RoleMiddleware(studenController.Delete, "student", false))
	router.GET("/api/lecture", RoleMiddleware(lecturesController.FindAll, "student", false))
	router.GET("/api/lecture/:name", lecturesController.FindById)
	router.POST("/api/lecture", RoleMiddleware(lecturesController.Create, "professor", false))
	router.PATCH("/api/lecture/:name", RoleMiddleware(lecturesController.Update, "professor", false))
	router.DELETE("/api/lecture/:name", RoleMiddleware(lecturesController.Delete, "professor", false))
	router.GET("/lecture/:departmentName/:lectureName", lecturesController.FindByIds)
	router.POST("/login/student", wrapHandler(controller.LoginHandlerToken))
	router.POST("/login/professor", wrapHandler(controller.LoginHandlerToken))
	router.POST("/SignUp", wrapHandler(controller.SignUp))
	router.POST("/logout", wrapHandler(controller.LogoutHandler))
	router.GET("/api/enrollment", RoleMiddleware(enrollmentController.FindAll, "professor", false))
	router.POST("/api/enrollment", RoleMiddleware(enrollmentController.Create, "student", false))
	router.DELETE("/api/enrollment/:email/:name", RoleMiddleware(enrollmentController.Delete, "student", false))

	return router
}
