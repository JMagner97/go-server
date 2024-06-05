package request

import "time"

type StudentUpdateRequest struct {
	Id           int
	Name         string `validate:"required,min=1,max=100" json:"name"`
	Surname      string `validate:"required,min=1,max=100" json:"surname"`
	Data         time.Time
	Address      string
	Email        string
	DepartmentId int
}
