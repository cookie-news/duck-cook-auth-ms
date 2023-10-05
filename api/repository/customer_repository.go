package repository

import "duck-cook-auth/entity"

type CustomerRepository interface {
	CreateCustomer(customer entity.CustomerInput) (entity.Customer, error)
	GetCustomerByUser(user string) (entity.CustomerInput, error)
}
