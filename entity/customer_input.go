package entity

type CustomerInput struct {
	Customer
	Pass string `json:"pass" format:"string"`
}

func (c *CustomerInput) ToCustomer() Customer {
	return Customer{
		ID:    c.ID,
		Email: c.Email,
		User:  c.User,
	}
}
