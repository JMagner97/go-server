package enumrole

import (
	"net/http"

	st "go-server/utility"
)

type Role int

const (
	Professor = iota + 1
	Student
	Admin
	Unknown
)

func IsAdmin(r *http.Request) bool {
	s, _ := st.Store.Get(r, "idsession")
	return s.Values["role"] == Admin
}

func IsProfessor(r *http.Request) bool {
	s, _ := st.Store.Get(r, "idsession")
	return s.Values["role"] == Professor
}

func IsStudent(r *http.Request) bool {
	s, _ := st.Store.Get(r, "idsession")
	return s.Values["role"] == Student
}
