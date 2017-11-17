package dto

type Login struct {
	Login	string	`json:"login"`
	Pass	string	`json:"password"`
}

type AuthClient struct {
	ID		string
	Level	string
	Session	string
}