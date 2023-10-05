package usecase

import (
	"duck-cook-auth/api/repository"
	"duck-cook-auth/entity"
	"errors"
	"regexp"
)

type CustomerUseCase interface {
	Validate(customer entity.CustomerInput) (err error)
	CreateCustomer(customer entity.CustomerInput) (entity.Customer, error)
	GetCustomerByUser(user string) (entity.CustomerInput, error)
}

type customerUseCaseImpl struct {
	customerRepository repository.CustomerRepository
}

func (usecase customerUseCaseImpl) GetCustomerByUser(user string) (entity.CustomerInput, error) {
	return usecase.customerRepository.GetCustomerByUser(user)
}

func (usecase customerUseCaseImpl) CreateCustomer(customer entity.CustomerInput) (customerResult entity.Customer, err error) {
	err = usecase.Validate(customer)
	if err != nil {
		return
	}
	return usecase.customerRepository.CreateCustomer(customer)
}

var (
	ErrInvalidEmail = errors.New("email is invalid")
)

func (*customerUseCaseImpl) Validate(customer entity.CustomerInput) (err error) {
	emailIsValid := checkEmail(customer.Email)
	if !emailIsValid {
		return ErrInvalidEmail
	}

	return nil
}

func checkEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}

func NewCustomerUseCase(customerRepository repository.CustomerRepository) CustomerUseCase {
	return &customerUseCaseImpl{
		customerRepository,
	}
}
