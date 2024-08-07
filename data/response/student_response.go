package response

import "time"

type StudentResponse struct {
	//Id           int       `json:"id"`
	Username     string    `json:"username,omitempty"`
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	Data         time.Time `json:"data"`
	Address      string    `json:"address"`
	Email        string    `json:"email"`
	DepartmentId int       `json:"departmentId"`
}
