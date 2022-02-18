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
	AllPost(ctx *gin.Context)
	DetailPost(ctx *gin.Context)
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
// @Tags Company
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} entity.User
// @Failure 400 {object} dto.ErrorResponse
// @Router /company/user/register [get]
func (c *companyController) AllUserRegister(ctx *gin.Context) {
	companyed, err := c.companyRepo.AllUserRegister()
	if err != nil {
		var errResponse dto.ErrorResponse
		errResponse.Message = "Failed to Fetch"
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusOK, companyed)
	}
}

// All User Login
// @Summary Retrieves the list of users who has to login in the app
// @Description Retrieves the list of users who has to login in the app
// @Tags Company
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} entity.User
// @Failure 400 {object} dto.ErrorResponse
// @Router /company/user/login [get]
func (c *companyController) AllUserLogin(ctx *gin.Context) {
	companyed, err := c.companyRepo.AllUserLogin()
	if err != nil {
		var errResponse dto.ErrorResponse
		errResponse.Message = "Failed to Fetch"
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusOK, companyed)
	}
}

// All User Point
// @Summary Retrieves all users point
// @Description Retrieves all users point
// @Tags Company
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {integer} point
// @Failure 400 {object} dto.ErrorResponse
// @Router /company/user/point [get]
func (c *companyController) AllUserPoint(ctx *gin.Context) {
	companyed, err := c.companyRepo.AllUserPoint()
	if err != nil {
		var errResponse dto.ErrorResponse
		errResponse.Message = "Failed to Fetch"
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusOK, companyed)
	}
}

// Detail User
// @Summary Retrieves detail one user
// @Description Retrieves detail one user
// @Tags Company
// @Accept  */*
// @Produce  json
// @Param id path string true "User Id"
// @Security ApiKeyAuth
// @Success 200 {object} entity.User
// @Failure 400 {object} dto.ErrorResponse
// @Router /company/user/:id [get]
func (c *companyController) DetailUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userId := uuid.MustParse(id)
	companyed, err := c.companyRepo.DetailUser(userId)
	if err != nil {
		var errResponse dto.ErrorResponse
		errResponse.Message = "Failed to Fetch"
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusOK, companyed)
	}
}

// All User Point
// @Summary Retrieves detail point from one user
// @Description Retrieves detail point from one user
// @Tags Company
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Param id path integer true "Post Id"
// @Param id path string true "User Id"
// @Success 200 {integer} point
// @Failure 400 {object} dto.ErrorResponse
// @Router /company/user/point/:id [get]
func (c *companyController) DetailUserPoint(ctx *gin.Context) {
	id := ctx.Param("id")
	userId := uuid.MustParse(id)
	companyed, err := c.companyRepo.DetailUserPoint(userId)
	if err != nil {
		var errResponse dto.ErrorResponse
		errResponse.Message = "Failed to Fetch"
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusOK, companyed)
	}
}

// All User Point
// @Summary Retrieves update point on one post
// @Description Retrieves update point on one post
// @Tags Company
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Param id path integer true "Post Id"
// @Param data body dto.PostPointDTO true "Point"
// @Success 200 {object} entity.Post 
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /company/user/:id [patch]
func (c *companyController) PostPoint(ctx *gin.Context) {
	var point dto.PostPointDTO
	id := ctx.Param("id")
	postId, err := strconv.Atoi(id)
	var errResponse dto.ErrorResponse
	if err != nil {
		errResponse.Message = "Failed to Fetch"
		ctx.JSON(http.StatusBadRequest, errResponse)
		panic(err)
	}
	err = ctx.ShouldBindJSON(&point)
	if err != nil {
		errResponse.Message = "Failed to Bind"
		ctx.JSON(http.StatusInternalServerError, errResponse)
		panic(err)
	}
	pointed, _ := c.companyRepo.PostPoint(postId, point.Point)
	ctx.JSON(http.StatusOK, pointed)
}

// All Post
// @Summary All Post from the app that has been posted
// @Description All Post from the app that has been posted
// @Tags Company
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} entity.Post
// @Failure 400 {object} dto.ErrorResponse
// @Router /company/post [get]
func (c *companyController) AllPost(ctx *gin.Context) {
	posted, err := c.companyRepo.AllPost()
	if err != nil {
		var errResponse dto.ErrorResponse
		errResponse.Message = "Failed to Fetch"
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusOK, posted)
	}
}

// Detail Post
// @Summary All Post from the app that has been posted
// @Description All Post from the app that has been posted
// @Tags Company
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Post ID"
// @Success 200 {object} entity.Post
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /company/post/:id [get]
func (c *companyController) DetailPost(ctx *gin.Context) {
	id := ctx.Param("id")
	var errResponse dto.ErrorResponse
	postId, err := strconv.Atoi(id)
	if err != nil {
		errResponse.Message = "Failed to Convert to Param"
		ctx.JSON(http.StatusInternalServerError, errResponse)
		panic(err)
	}
	posted, err := c.companyRepo.DetailPost(postId)
	if err != nil {
		errResponse.Message = "Failed to Fetch"
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusOK, posted)
	}
}