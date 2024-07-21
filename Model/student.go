package Model

import (
	"time"
)

type Student struct {
	Username     string
	Password     string
	Role         int
	Id           int
	Name         string
	Surname      string
	Birthdate    time.Time
	Address      string
	Email        string
	DepartmentId int
}
