package route

import (
	"fmt"
	"go-server/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func wrapHandler(handler http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		handler(w, r)
	}
}

func NewRouter(studenController *controller.StudentController, courseController *controller.CourseController, enrollmentController *controller.EnrollmentController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome Home")
	})

	router.GET("/api/student", studenController.FindAll)
	router.GET("/api/student/:idstudente", studenController.FindById)
	router.POST("/api/student", studenController.Create)
	router.PATCH("/api/student/:idstudente", studenController.Update)
	router.DELETE("/api/student/:idstudente", studenController.Delete)
	router.GET("/api/course", courseController.FindAll)
	router.GET("/api/course/:courseid", courseController.FindById)
	router.POST("/api/course", courseController.Create)
	router.PATCH("/api/course/:courseid", courseController.Update)
	router.DELETE("/api/course/:courseid", courseController.Delete)
	router.POST("/login", wrapHandler(controller.LoginHandler))
	router.GET("/dashboard", wrapHandler((controller.DashboardHandler)))
	router.POST("/logout", wrapHandler(controller.LogoutHandler))
	router.GET("/api/enrollment", enrollmentController.FindAll)
	router.POST("/api/enrollment", enrollmentController.Create)
	router.DELETE("/api/enrollment/:studentid/:courseid", enrollmentController.Delete)
	return router
}
