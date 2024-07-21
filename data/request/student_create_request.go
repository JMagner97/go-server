package request

import "time"

type StudentCreateRequest struct {
	//Id           int
	Username     string `validate:"required,min=5,max=20" json:"username"`
	Password     string
	Name         string `validate:"required,min=1,max=100" json:"name"`
	Surname      string `validate:"required,min=1,max=100" json:"surname"`
	Data         time.Time
	Address      string
	Email        string
	DepartmentId int
}
