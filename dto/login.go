package dto

type Login struct {
	Login string `json:"login"`
	Pass  string `json:"password"`
}

type AuthUser struct {
	ID      string
	Level   string
	Session string
}
