package customer_repository

import (
	"duck-cook-auth/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	Email     string             `bson:"email"`
	User      string             `bson:"user"`
	Pass      string             `bson:"pass"`
}

func (customer Customer) ToEntityCustomer() entity.Customer {
	return entity.Customer{
		ID:    customer.ID.Hex(),
		Email: customer.Email,
		User:  customer.User,
	}
}

func (customer Customer) ToEntityCustomerInput() entity.CustomerInput {
	helper := customer.ToEntityCustomer()
	return entity.CustomerInput{
		Customer: helper,
		Pass:     customer.Pass,
	}
}

func (Customer) FromEntity(customer entity.CustomerInput) *Customer {
	id, _ := primitive.ObjectIDFromHex(customer.ID)
	return &Customer{
		ID:    id,
		Email: customer.Email,
		Pass:  customer.Pass,
		User:  customer.User,
	}
}
