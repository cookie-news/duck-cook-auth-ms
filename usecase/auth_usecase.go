package usecase

import (
	"duck-cook-auth/entity"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface {
	Login(pass, hash string) (isValid bool, err error)
	GenerateJWT(customerInfo entity.Customer) (jwt string, err error)
	ValidateJWT(jwtstr string) (err error)
}

type authUseCaseImpl struct {
}

var (
	ErrTokenExpire       = errors.New("token expirado")
	ErrTokenTokenInvalid = errors.New("token expirado")
)

func (usecase *authUseCaseImpl) ValidateJWT(jwtstr string) (err error) {
	var secretKey = []byte(os.Getenv("SECRET_KEY_JWT"))
	token, err := jwt.ParseWithClaims(jwtstr, &entity.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inválido")
		}
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(*entity.Claims); ok && token.Valid {
		now := time.Now()
		if claims.ExpiresAt.Time.Before(now) {
			return ErrTokenExpire
		}
		return nil
	}

	return ErrTokenTokenInvalid
}

func (usecase *authUseCaseImpl) Login(pass, hash string) (isValid bool, err error) {
	isValid = CheckPasswordHash(pass, hash)
	return
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (usecase *authUseCaseImpl) GenerateJWT(customerInfo entity.Customer) (jwtstr string, err error) {
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
