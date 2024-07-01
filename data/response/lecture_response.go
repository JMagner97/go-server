package response

type LectureResponse struct {
	//LectureId    int    `json:"lectureid"`
	LectureName  string `json:"name"`
	Description  string `json:"description"`
	StartYear    int    `json:"startyear"`
	EndYear      int    `json:"endyear"`
	ProfessorId  int    `json:"professorid"`
	DepartmentId int    `json:"departmentid"`
}
