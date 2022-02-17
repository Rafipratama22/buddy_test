package controller

import (
	"net/http"
	"strconv"

	"github.com/Rafipratama22/mnc_test.git/dto"
	"github.com/Rafipratama22/mnc_test.git/entity"
	"github.com/Rafipratama22/mnc_test.git/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PostController interface {
	CreatePost(ctx *gin.Context)
	DetailPost(ctx *gin.Context)
	UpdatePost(ctx *gin.Context)
}

type postController struct {
	postRepo repository.PostRepository
}

func NewPostController(postRepo repository.PostRepository) PostController {
	return &postController{
		postRepo: postRepo,
	}
}

// Post Create
// @Summary Retrieves the list of users who has to register in the app
// @Description Retrieves the list of users who has to register in the app
// @Tags Post
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Success 201 {object} entity.Post
// @Failure 400 {object} dto.ErrorResponse
// @Router /post [post]
func (c *postController) CreatePost(ctx *gin.Context) {
	var post entity.Post
	var errResponse dto.ErrorResponse
	authorId := ctx.MustGet("user_id")
	userId := uuid.MustParse(authorId.(string))
	post.AuthorID = userId
	err := ctx.ShouldBindJSON(&post)
	if err != nil {
		panic(err)
	}
	posted, err := c.postRepo.CreatePost(post)
	if err != nil {
		errResponse.Message = "Failed to create post"
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusCreated, posted)
	}
}

// Detail Post
// @Summary All Post from the app that has been posted
// @Description All Post from the app that has been posted
// @Tags Post
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} entity.Post
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /post/:id [get]
func (c *postController) DetailPost(ctx *gin.Context) {
	var errResponse dto.ErrorResponse
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		errResponse.Message = "Failed to convert id to int"
		ctx.JSON(http.StatusInternalServerError, errResponse)
	}
	result, err := c.postRepo.DetailPost(id)
	if err != nil {
		errResponse.Message = "Failed to Found Post"
		ctx.JSON(http.StatusNotFound, errResponse)
		} else {
		ctx.JSON(http.StatusOK, result)
	}
}

// Update Post
// @Summary All Post from the app that has been posted
// @Description All Post from the app that has been posted
// @Tags Post
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} entity.Post
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /post/:id [put]
func (c *postController) UpdatePost(ctx *gin.Context) {
	var errResponse dto.ErrorResponse
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		errResponse.Message = "Failed to convert id to int"
		ctx.JSON(http.StatusInternalServerError, errResponse)
	}
	var post entity.Post
	err = ctx.ShouldBindJSON(&post)
	if err != nil {
		panic(err)
	}
	result, err := c.postRepo.UpdatePost(id, post)
	if err != nil{
		errResponse.Message = "Failed to update post"
		ctx.JSON(http.StatusBadRequest, errResponse)
		} else {
		ctx.JSON(http.StatusOK, result)
	}
}