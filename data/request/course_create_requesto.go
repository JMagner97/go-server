package request

type CourseCreateRequest struct {
	CourseId     int
	CourseName   string `validate:"required,min=1,max=100" json:"coursename"`
	Description  string `validate:"required,min=1,max=100" json:"desciption"`
	StartYear    int
	EndYear      int
	DepartmentId int
	ProfessorId  int
}
