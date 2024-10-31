package usecases

import (
	"errors"
	models "superIndo/model"
	"superIndo/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var Users = make(map[string]models.User)
var Id int = 0
var Carts = make(map[int]models.Cart)

func Register(username string, password string) (resp models.User, err error) {
	if username == "" || password == "" {
		return resp, errors.New("username or password is empty")
	}

	resp.HashPassword(password)
	Id += 1
	resp.ID = uint64(Id)
	resp.Username = username

	Users[resp.Username] = resp
	return resp, err
}

func Login(username string, password string) (resp *models.Claims, err error) {
	if username == "" || password == "" {
		return resp, errors.New("username or password is empty")
	}

	user, ok := Users[username]
	if !ok {
		return resp, errors.New("user does not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return resp, errors.New("wrong password")
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		UserId:   user.ID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(utils.JwtKey)
	if err != nil {
		return resp, errors.New("token error")
	}
	claims.Token = tokenString

	return claims, err

}

func GetListCategories() (resp []models.Category, err error) {
	return utils.CategoriesProduct, nil
}

func GetProductByCategoryId(categoryId uint64) (resp []models.Product, err error) {
	mappingProductAndCategoryId := make(map[uint64][]models.Product)
	for _, product := range utils.Products {
		mappingProductAndCategoryId[product.CategoryID] = append(mappingProductAndCategoryId[product.CategoryID], product)
	}

	resp, ok := mappingProductAndCategoryId[categoryId]
	if !ok {
		return resp, errors.New("category does not exist")
	}

	return resp, nil
}

func GetProductDetailByProductId(productId uint64) (resp models.ProductDetail, err error) {
	mappingProductIdAndDetail := make(map[uint64]models.ProductDetail)
	for _, productDetail := range utils.ProductDetails {
		mappingProductIdAndDetail[productDetail.ProductID] = productDetail
	}

	resp, ok := mappingProductIdAndDetail[productId]
	if !ok {
		return resp, errors.New("product does not exist")
	}

	return resp, nil
}

func AddItemToCart(userId uint64, newItems []models.CartDetail) (err error) {
	mappingProductIdAndDetail := make(map[uint64]*models.ProductDetail)
	for _, productDetail := range utils.ProductDetails {
		mappingProductIdAndDetail[productDetail.ProductID] = &productDetail
	}
	for _, newItem := range newItems {
		if mappingProductIdAndDetail[newItem.ProductId] == nil {
			return errors.New("product does not exist")
		}
	}

	cart, ok := Carts[int(userId)]
	if !ok {
		cart = models.Cart{UserId: userId, CartDetails: []models.CartDetail{}}
	}

	for _, newItem := range newItems {
		for i, cartDetail := range cart.CartDetails {
			if cartDetail.ProductId == newItem.ProductId {
				cart.CartDetails[i].Quantity += newItem.Quantity
				Carts[int(userId)] = cart
				return
			}
		}

		cart.CartDetails = append(cart.CartDetails, newItem)
		Carts[int(userId)] = cart
	}

	return

}

func GetCurrentUserCartDetail(userId uint64) (resp models.Cart, err error) {
	resp, ok := Carts[int(userId)]
	if !ok {
		return resp, errors.New("user doesnt have cart yet")
	}
	return resp, nil
}
