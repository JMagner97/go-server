package response

type EnrollmentResponse struct {
	Name             string `json:"name"`
	Surname          string `json:"surname"`
	Email            string `json:"email"`
	LectureName      string `json:"lectureName"`
	StartYear        int    `json:"startyear"`
	EndYear          int    `json:"endyear"`
	ProfessorSurname string `json:"professorSurname"`
	DepartmentName   string `json:"departmentName"`
}
