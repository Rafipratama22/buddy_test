package controller

import (
	"github.com/Rafipratama22/mnc_test.git/entity"
	"github.com/Rafipratama22/mnc_test.git/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetAllPoint(ctx *gin.Context)
}

type userController struct {
	userRepo repository.UserRepository	
}

func NewUserController(userRepo repository.UserRepository) UserController {
	return &userController{
		userRepo: userRepo,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}
	usered, err := c.userRepo.RegisterUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Create"})
	} else {
		ctx.JSON(http.StatusCreated, usered)
	}
}

func (c *userController) Login(ctx *gin.Context) {
	var user entity.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}
	token, err := c.userRepo.LoginUser(user.Email, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func (c *userController) GetAllPoint(ctx *gin.Context) {
	user_id := uuid.MustParse(ctx.Param("user_id"))
	point, err := c.userRepo.GetAllPoint(user_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"point": point})
	}
}

