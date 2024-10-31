package controllers

import (
	"net/http"
	"strconv"
	models "superIndo/model"
	usecases "superIndo/usecase"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	user, err := usecases.Register(user.Username, user.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "username": user.Username})
}

func Login(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	resp, err := usecases.Login(user.Username, user.Password)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": resp})
}

func GetCurrentUserCartDetail(context *gin.Context) {
	userId, ok := context.Get("userId")
	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	resp, err := usecases.GetCurrentUserCartDetail(uint64(userId.(int)))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": resp})
}

func ListCategory(context *gin.Context) {
	resp, err := usecases.GetListCategories()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": resp})
}

func GetProductByCategoryId(context *gin.Context) {
	param := context.Param("id")
	categoryID, err := strconv.Atoi(param)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	resp, err := usecases.GetProductByCategoryId(uint64(categoryID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": resp})
}

func GetProductDetailByProductId(context *gin.Context) {
	id := context.Param("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	resp, err := usecases.GetProductDetailByProductId(uint64(productId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": resp})
}

func AddProductToCart(context *gin.Context) {
	userId, ok := context.Get("userId")
	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var newItem []models.CartDetail
	if err := context.ShouldBindJSON(&newItem); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := usecases.AddItemToCart(uint64(userId.(int)), newItem); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": "product added"})
}
