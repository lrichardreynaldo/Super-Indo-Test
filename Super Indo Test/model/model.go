package models

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	UserId   uint64 `json:"userId"`
	Username string `json:"username"`
	Token    string `json:"token"`
	jwt.StandardClaims
}

type Category struct {
	ID           uint64 `json:"id"`
	CategoryName string `json:"categoryName"`
}

type Product struct {
	ID          uint64 `json:"id"`
	CategoryID  uint64 `json:"categoryId"`
	ProductName string `json:"productName"`
}

type ProductDetail struct {
	ProductID   uint64  `json:"productId"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       uint64  `json:"stock"`
}

type Cart struct {
	UserId      uint64       `json:"userId"`
	CartDetails []CartDetail `json:"cartDetails"`
}

type CartDetail struct {
	ProductId uint64 `json:"productId"`
	Quantity  uint64 `json:"quantity"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
