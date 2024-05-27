package response

type CourseResponse struct {
	CourseId     int    `json:"courseid"`
	CourseName   string `json:"coursename"`
	Description  string `json:"desciption"`
	StartYear    int    `json:"startyear"`
	EndYear      int    `json:"endyear"`
	DepartmentId int    `json:"departmentid"`
	ProfessorId  int    `json:"professorid"`
}
