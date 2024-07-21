package request

type ProfessorRequest struct {
	Username     string `validate:"required,min=5,max=20" json:"username"`
	Password     string
	Name         string `validate:"required,min=1,max=100" json:"name"`
	Surname      string `validate:"required,min=1,max=100" json:"surname"`
	Email        string
	DepartmentId int
}
