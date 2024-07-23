package Enumrole

import (
	"net/http"

	st "go-server/Utility"
)

type Role int

const (
	Student = iota + 1
	Professor
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
