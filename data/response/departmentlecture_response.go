package response

type DepartmentLectureResponse struct {
	DepartmentName string `json:"departmentName"`
	LectureName    string `json:"lectureName"`
	Description    string `json:"description"`
	StartYear      int    `json:"startyear"`
	EndYear        int    `json:"endyear"`
}
