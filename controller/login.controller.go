package controller

import (
	"github.com/Rafipratama22/mnc_test.git/dto"
	"github.com/Rafipratama22/mnc_test.git/entity"
	"github.com/Rafipratama22/mnc_test.git/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	RegisterUser(ctx *gin.Context)
	RegisterCompany(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type loginController struct {
	loginRepo repository.LoginRepository
}

func NewLoginController(loginRepo repository.LoginRepository) LoginController {
	return &loginController{
		loginRepo: loginRepo,
	}
}

func (c *loginController) RegisterUser(ctx *gin.Context) {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}
	usered, err := c.loginRepo.RegisterUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Create"})
	} else {
		ctx.JSON(http.StatusCreated, usered)
	}
}

func (c *loginController) RegisterCompany(ctx *gin.Context) {
	var company entity.Company
	err := ctx.ShouldBindJSON(&company)
	if err != nil {
		panic(err)
	}
	companyed, err := c.loginRepo.RegisterCompany(company)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Create"})
	} else {
		ctx.JSON(http.StatusCreated, companyed)
	}
}


func (c *loginController) Login(ctx *gin.Context) {
	var loginDto dto.LoginDto
	err := ctx.ShouldBindJSON(&loginDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	role := loginDto.Role
	if role == "user" {
		token, err := c.loginRepo.LoginUser(loginDto.Email, loginDto.Password)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"token": token})
		}
	} else if role == "company" {
		token, err := c.loginRepo.LoginCompany(loginDto.Email, loginDto.Password)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"token": token})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "role must be user or company"})
	}
}

// func (c *loginController) LoginCompany(ctx *gin.Context) {
// 	var company entity.Company
// 	err := ctx.ShouldBindJSON(&company)
// 	if err != nil {
// 		panic(err)
// 	}
// 	token, err := c.loginRepo.LoginCompany(company.Email, company.Password)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
// 	} else {
// 		ctx.JSON(http.StatusOK, gin.H{"token": token})
// 	}
// }