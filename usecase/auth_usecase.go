package usecase

import (
	"duck-cook-auth/entity"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface {
	Login(pass, hash string) (isValid bool, err error)
	GenerateJWT(customerInfo entity.CustomerInput) (jwt string, err error)
	ValidateJWT(jwtstc string) (isValid bool, err error)
}

type authUseCaseImpl struct {
}

func (usecase *authUseCaseImpl) ValidateJWT(jwtstc string) (isValid bool, err error) {
	return
}

func (usecase *authUseCaseImpl) Login(pass, hash string) (isValid bool, err error) {
	isValid = CheckPasswordHash(pass, hash)
	return
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (usecase *authUseCaseImpl) GenerateJWT(customerInfo entity.CustomerInput) (jwtstr string, err error) {
	var secretKey = []byte(os.Getenv("SECRET_KEY_JWT"))
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &entity.Claims{
		UserId: customerInfo.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtstr, err = token.SignedString(secretKey)
	return
}

func NewAuthUserCase() AuthUseCase {
	return &authUseCaseImpl{}
}
