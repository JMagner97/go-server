package main

import (
	"fmt"
	repository "go-server/Repository"
	"go-server/controller"
	"go-server/helper"
	"go-server/route"
	"go-server/service"
	"go-server/service/enrollment"
	"go-server/service/lectures"
	professorservice "go-server/service/professor_service"
	"log"
	"net/http"

	_ "github.com/gorilla/sessions"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost" //"host.docker.internal"
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

	//connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	//	os.Getenv("DB_HOST"),
	//	os.Getenv("DB_PORT"),
	//	os.Getenv("DB_USER"),
	//	os.Getenv("DB_PASSWORD"),
	//	os.Getenv("DB_NAME"))

	psqlInfo := fmt.Sprintf("host=%s port=%d  dbname=%s user=%s password=%s sslmode=disable",
		host, port, dbname, user, password)

	fmt.Println("Connection string:", psqlInfo)

	db, err := repository.NewDatabase("postgres", psqlInfo)

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
	lectureRepo := repository.NewLectureRepo(db)
	enrollmentRepo := repository.NewStudentLectureRepo(db)
	professorRepo := repository.NewProfessorRepo(db)
	//servizio
	studentService := service.NewStudentServiceImpl(studentRepo)
	lecturesService := lectures.NewLectureServiceImpl(lectureRepo)
	enrollmentService := enrollment.NewStudentLectureServiceImpl(enrollmentRepo)
	professorService := professorservice.NewProfessorService(professorRepo)
	//controller
	studentController := controller.NewStudentController(studentService)
	lecturesController := controller.NewLecturesController(lecturesService)
	enrollmentController := controller.NewStudentLectureController(enrollmentService)
	professorController := controller.NewProfessorsController(professorService)

	routes := route.NewRouter(studentController, lecturesController, enrollmentController, professorController)
	fmt.Println("Sto qui per il server")
	server := http.Server{Addr: "0.0.0.0:8080", Handler: routes}

	errx := server.ListenAndServe()
	if errx != nil {
		fmt.Println("Sto in errore", errx)
		log.Fatal("ListenAndServe: ", errx)
	}
	helper.PanicIfError(errx)
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	//fmt.Println("Starting server at port 8080")
	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	log.Fatal(err)
	//}
}
