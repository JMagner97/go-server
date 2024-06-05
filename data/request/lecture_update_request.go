package request

type LectureUpdateRequest struct {
	LectureId   int    `json:"lectureid"`
	LectureName string `validate:"required,min=1,max=100" json:"name"`
	Description string `validate:"required,min=1,max=100" json:"description"`
	StartYear   int
	EndYear     int
	ProfessorId int
}
