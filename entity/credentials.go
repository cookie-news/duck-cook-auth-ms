package entity

type Credentials struct {
	User string `json:"user" format:"string"`
	Pass string `json:"pass" format:"string"`
}
