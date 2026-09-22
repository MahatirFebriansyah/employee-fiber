package model

type Employee struct {
	ID		int    `json:"id"`
	Name	string `json:"name"`
	Email	string `json:"email"`
	Position string `json:"position"`
}