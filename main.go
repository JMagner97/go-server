package main

import (
	"database/sql"
	"fmt"
	repository "go-server/Repository"
	"go-server/controller"
	"go-server/helper"
	"go-server/route"
	"go-server/service"
	"go-server/service/course"
	"go-server/service/enrollment"
	"log"
	"net/http"

	_ "github.com/gorilla/sessions"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "lorenzomagnano"
	password = "admin"
	dbname   = "provadb"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "Not found 404", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	// address := r.FormValue("address") // Uncomment if you need to handle the address field
	fmt.Fprintf(w, "name = %s\n", name)
}

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d  dbname=%s user=%s password=%s sslmode=disable",
		host, port, dbname, user, password)

	fmt.Println("Connection string:", psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	fmt.Println("Successfully connected to the database!")

	//studentRepo := repository.StudentRepo.
	//repository
	studentRepo := repository.NewStudRepo(db)
	courseRepo := repository.NewCourseRepo(db)
	enrollmentRepo := repository.NewEnrollmentRepo(db)
	//servizio
	studentService := service.NewStudentServiceImpl(studentRepo)
	coursesService := course.NewCourseServiceImpl(courseRepo)
	enrollmentService := enrollment.NewEnrollmentServiceImpl(enrollmentRepo)
	//controller
	studentController := controller.NewStudentController(studentService)
	courseController := controller.NewCourseController(coursesService)
	enrollmentController := controller.NewEnrollmentController(enrollmentService)
	routes := route.NewRouter(studentController, courseController, enrollmentController)

	server := http.Server{Addr: "localhost:8888", Handler: routes}
	errx := server.ListenAndServe()
	helper.PanicIfError(errx)
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}