package entity

type Customer struct {
	ID    string `json:"id" example:"1" format:"string"`
	Email string `json:"email" example:"usuario@host.com"`
	User  string `json:"user" example:"paulo"`
}
