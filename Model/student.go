package Model

import (
	"time"
)

type Student struct {
	Id        int
	Name      string
	Surname   string
	Birthdate time.Time
	Address   string
	Email     string
}
