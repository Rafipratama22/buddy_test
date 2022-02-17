package controller

import (
	"github.com/Rafipratama22/mnc_test.git/dto"
	"github.com/Rafipratama22/mnc_test.git/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CompanyController interface {
	AllUserRegister(ctx *gin.Context)
	AllUserLogin(ctx *gin.Context)
	AllUserPoint(ctx *gin.Context)
	DetailUser(ctx *gin.Context)
	DetailUserPoint(ctx *gin.Context)
	PostPoint(ctx *gin.Context)
}

type companyController struct {
	companyRepo repository.CompanyRepository
}

func NewCompanyController(companyRepo repository.CompanyRepository) CompanyController {
	return &companyController{
		companyRepo: companyRepo,
	}
}

// All User Register
// @Summary Retrieves the list of users who has to register in the app
// @Description Retrieves the list of users who has to register in the app
// @Tags root
// @Accept  */*
// @Produce  json
// @Success 200 {array} map[string]interface{}
// @Router /company/user/register [get]
func (c *companyController) AllUserRegister(ctx *gin.Context) {
	companyed, err := c.companyRepo.AllUserRegister()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Fetch"})
	} else {
		ctx.JSON(http.StatusOK, companyed)
	}
}

// All User Login
// @Summary Retrieves the list of users who has to login in the app
// @Description Retrieves the list of users who has to login in the app
// @Produce  json
// @Success 200 {array} map[string]interface{}
// @Router /company/user/login [get]
func (c *companyController) AllUserLogin(ctx *gin.Context) {
	companyed, err := c.companyRepo.AllUserLogin()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Fetch"})
	} else {
		ctx.JSON(http.StatusOK, companyed)
	}
}

// All User Point
// @Summary Retrieves all users point
// @Description Retrieves all users point
// @Produce  json
// @Success 200 {integer} point
// @Router /company/user/point [get]
func (c *companyController) AllUserPoint(ctx *gin.Context) {
	companyed, err := c.companyRepo.AllUserPoint()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Fetch"})
	} else {
		ctx.JSON(http.StatusOK, companyed)
	}
}

func (c *companyController) DetailUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userId := uuid.MustParse(id)
	companyed, err := c.companyRepo.DetailUser(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Fetch"})
	} else {
		ctx.JSON(http.StatusOK, companyed)
	}
}

func (c *companyController) DetailUserPoint(ctx *gin.Context) {
	id := ctx.Param("id")
	userId := uuid.MustParse(id)
	companyed, err := c.companyRepo.DetailUserPoint(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Fetch"})
	} else {
		ctx.JSON(http.StatusOK, companyed)
	}
}

func (c *companyController) PostPoint(ctx *gin.Context) {
	var point dto.PostPointDTO
	id := ctx.Param("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	err = ctx.ShouldBindJSON(&point)
	if err != nil {
		panic(err)
	}
	pointed, _ := c.companyRepo.PostPoint(postId, point.Point)
	ctx.JSON(http.StatusOK, pointed)
}