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

func NewRouter(studenController *controller.StudentController, lecturesController *controller.LectureController, enrollmentController *controller.LectureStudentController, professorController *controller.ProfessorController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome Home")
	})

	router.GET("/api/student", studenController.FindAll)
	router.GET("/api/student/:email", studenController.FindById)
	router.POST("/api/student", studenController.Create)
	router.POST("/api/professor", wrapHandler(professorController.Create))
	router.PATCH("/api/student/:email", studenController.Update)
	router.DELETE("/api/student/:email", studenController.Delete)
	router.GET("/api/lecture", lecturesController.FindAll)
	router.GET("/api/lecture/:name", lecturesController.FindById)
	router.POST("/api/lecture", lecturesController.Create)
	router.PATCH("/api/lecture/:name", lecturesController.Update)
	router.DELETE("/api/lecture/:name", lecturesController.Delete)
	router.GET("/lecture/:departmentName/:lectureName", lecturesController.FindByIds)
	router.POST("/login", wrapHandler(controller.LoginHandlerToken))
	router.POST("/SignUp", wrapHandler((controller.SignUp)))
	router.POST("/logout", wrapHandler(controller.LogoutHandler))
	router.GET("/api/enrollment", enrollmentController.FindAll)
	router.POST("/api/enrollment", enrollmentController.Create)
	router.DELETE("/api/enrollment/:email/:name", enrollmentController.Delete)

	return router
}
