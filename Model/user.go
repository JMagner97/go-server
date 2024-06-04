package Model

type User struct {
	Id       int
	Username string
	Password string
	Token    string
	Valid    bool
}
