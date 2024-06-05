package request

type LectureCreateRequest struct {
	LectureId   int
	LectureName string `validate:"required,min=1,max=100" json:"name"`
	Description string `validate:"required,min=1,max=100" json:"description"`
	StartYear   int
	EndYear     int
	ProfessorId int
}
