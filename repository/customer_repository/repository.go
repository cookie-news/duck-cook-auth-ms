package customer_repository

import (
	"context"
	"duck-cook-auth/api/repository"
	"duck-cook-auth/entity"
	"errors"
	"log"
	"regexp"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	_mongo "duck-cook-auth/pkg/mongo"
)

type repositoryImpl struct {
	customerCollection *mongo.Collection
}

func (repo repositoryImpl) GetCustomerByUser(user string) (customer entity.CustomerInput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var customerModel Customer

	err = repo.customerCollection.FindOne(ctx, bson.M{"user": user}).Decode(&customerModel)
	customer = customerModel.ToEntityCustomerInput()
	return
}

func (repo repositoryImpl) CreateCustomer(customer entity.CustomerInput) (entity.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	var customerModel Customer
	customerModel = *customerModel.FromEntity(customer)
	timeNow := time.Now()
	customerModel.CreatedAt = timeNow
	customerModel.UpdatedAt = timeNow
	passHash, err := HashPassword(customer.Pass)
	if err != nil {
		return customerModel.ToEntityCustomer(), err
	}
	customerModel.Pass = passHash
	res, err := repo.customerCollection.InsertOne(ctx, &customerModel)

	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, writeErr := range writeErr.WriteErrors {
				if writeErr.Code == 11000 {
					errorMsg := writeErr.Message
					startIdx := strings.Index(errorMsg, "{")
					endIdx := strings.Index(errorMsg, "}")
					if startIdx != -1 && endIdx != -1 {
						fieldInfo := errorMsg[startIdx+1 : endIdx]

						re := regexp.MustCompile(`(\w+):`)
						match := re.FindStringSubmatch(fieldInfo)
						if len(match) >= 2 {
							fieldName := match[1]
							return customerModel.ToEntityCustomer(), errors.New("duplicate " + fieldName)
						}
					}

				} else {
					log.Fatal(err)
				}
			}
		} else {
			log.Fatal(err)
		}
	}

	customerModel.ID = res.InsertedID.(primitive.ObjectID)

	return customerModel.ToEntityCustomer(), nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func New(mongoDb mongo.Database) repository.CustomerRepository {
	customerCollection := mongoDb.Collection(_mongo.COLLETCTION_CUSTOMER)
	return &repositoryImpl{
		customerCollection,
	}
}
