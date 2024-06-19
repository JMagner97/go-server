package response

type EnrollmentResponse struct {
	Name             string `json:"name"`
	Surname          string `json:"surname"`
	Email            string `json:"email"`
	LectureName      string `json:"lecture_name"`
	StartYear        int    `json:"startyear"`
	EndYear          int    `json:"endyear"`
	ProfessorSurname string `json:"professor_surname"`
	DepartmentName   string `json:"department_name"`
}
