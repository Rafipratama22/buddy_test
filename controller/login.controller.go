package controller

import (
	"net/http"

	"github.com/Rafipratama22/mnc_test.git/dto"
	"github.com/Rafipratama22/mnc_test.git/entity"
	"github.com/Rafipratama22/mnc_test.git/repository"
	// "github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	RegisterUser(ctx *gin.Context)
	RegisterCompany(ctx *gin.Context)
	Login(ctx *gin.Context)
	LogOutCompany(ctx *gin.Context)
	LogOutUser(ctx *gin.Context)
}

type loginController struct {
	loginRepo repository.LoginRepository
}

func NewLoginController(loginRepo repository.LoginRepository) LoginController {
	return &loginController{
		loginRepo: loginRepo,
	}
}

// Register User
// @Summary Register user in the app
// @Description Register user in the app
// @Tags Form
// @Accept  */*
// @Produce  json
// @Param data body dto.RegisterUser true "User"
// @Success 201 {object} entity.User
// @Failure 400 {object} dto.ErrorResponse
// @Router /form/register/user [post]
func (c *loginController) RegisterUser(ctx *gin.Context) {
	var user entity.User
	var errResponse dto.ErrorResponse
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}
	usered, err := c.loginRepo.RegisterUser(user)
	if err != nil {
		errResponse.Message = "Failed to Create"
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusCreated, usered)
	}
}

// Register Company
// @Summary Register company in the app
// @Description Register company in the app
// @Tags Form
// @Accept  */*
// @Produce  json
// @Param data body dto.RegisterCompany true "Company"
// @Success 201 {object} entity.Company
// @Failure 400 {object} dto.ErrorResponse
// @Router /form/register/company [post]
func (c *loginController) RegisterCompany(ctx *gin.Context) {
	var company entity.Company
	var errResponse dto.ErrorResponse
	err := ctx.ShouldBindJSON(&company)
	if err != nil {
		panic(err)
	}
	companyed, err := c.loginRepo.RegisterCompany(company)
	if err != nil {
		errResponse.Message = "Failed to Create"
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusCreated, companyed)
	}
}

// Login
// @Summary Login for user and company in the app
// @Description Login for user and company company in the app
// @Tags Form
// @Accept  */*
// @Produce  json
// @Param data body dto.LoginDto true "Company"
// @Success 200 {string} {token: "Bearer {token}"}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /form/login [post]
func (c *loginController) Login(ctx *gin.Context) {
	var loginDto dto.LoginDto
	var errResponse dto.ErrorResponse
	err := ctx.ShouldBindJSON(&loginDto)
	if err != nil {
		errResponse.Message = "Failed to Bind"
		ctx.JSON(http.StatusInternalServerError, errResponse)
	}
	role := loginDto.Role
	if role == "user" {
		token, err := c.loginRepo.LoginUser(loginDto.Email, loginDto.Password)
		if err != nil {
			errResponse.Message = "Failed to Login User"
			ctx.JSON(http.StatusBadRequest, errResponse)
		} else {
			ctx.JSON(http.StatusOK, gin.H{"token": token})
		}
	} else if role == "company" {
		token, err := c.loginRepo.LoginCompany(loginDto.Email, loginDto.Password)
		if err != nil {
			errResponse.Message = "Failed to Login Company"
			ctx.JSON(http.StatusBadRequest, errResponse)
		} else {
			ctx.JSON(http.StatusOK, gin.H{"token": token})
		}
	} else {
		errResponse.Message = "role must be user or company"
		ctx.JSON(http.StatusBadRequest, errResponse)
	}
}

// Log Out Company
// @Summary Log Out Company in the app
// @Description Log Out Company in the app
// @Tags Form
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} dto.LogOutDto
// @Failure 400 {object} dto.ErrorResponse
// @Router /logut/company [post]
func (c *loginController) LogOutCompany(ctx *gin.Context) {
	var errResponse dto.ErrorResponse
	var logOutDto dto.LogOutDto
	authorId := ctx.MustGet("user_id")
	err := c.loginRepo.LogOutCompany(authorId.(string))
	if err != nil {
		errResponse.Message = "Failed to Logout Company"
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		logOutDto.Message = "Logout Company Success"
		ctx.JSON(http.StatusOK, logOutDto)
	}
}

// Log Out Company
// @Summary Log Out User in the app
// @Description Log Out User in the app
// @Tags Form
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} dto.LogOutDto
// @Failure 400 {object} dto.ErrorResponse
// @Router /logut/user [post]
func (c *loginController) LogOutUser(ctx *gin.Context) {
	var errResponse dto.ErrorResponse
	var logOutDto dto.LogOutDto
	authorId := ctx.MustGet("user_id")
	err := c.loginRepo.LogOutUser(authorId.(string))
	if err != nil {
		errResponse.Message = "Failed to Logout User"
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		logOutDto.Message = "Logout User Success"
		ctx.JSON(http.StatusOK, gin.H{"message": "Logout User Success"})
	}
}
