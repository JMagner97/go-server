package request

type CourseUpdateRequest struct {
	CourseId     int    `json:"courseid"`
	CourseName   string `validate:"required,min=1,max=100" json:"coursename"`
	Description  string `validate:"required,min=1,max=100" json:"desciption"`
	StartYear    int
	EndYear      int
	DepartmentId int
	ProfessorId  int
}
