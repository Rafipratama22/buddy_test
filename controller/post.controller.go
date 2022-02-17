package controller

import (
	
	"github.com/Rafipratama22/mnc_test.git/entity"
	"github.com/Rafipratama22/mnc_test.git/repository"
	"net/http"
	"strconv"

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

func (c *postController) CreatePost(ctx *gin.Context) {
	var post entity.Post
	authorId := ctx.MustGet("user_id")
	userId := uuid.MustParse(authorId.(string))
	post.AuthorID = userId
	err := ctx.ShouldBindJSON(&post)
	if err != nil {
		panic(err)
	}
	posted, err := c.postRepo.CreatePost(post)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Create"})
	} else {
		ctx.JSON(http.StatusCreated, posted)
	}
}

func (c *postController) DetailPost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert id to int"})
	}
	result, err := c.postRepo.DetailPost(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Failed to Found Post"})
		} else {
		ctx.JSON(http.StatusOK, result)
	}
}

func (c *postController) UpdatePost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert id to int"})
	}
	var post entity.Post
	err = ctx.ShouldBindJSON(&post)
	if err != nil {
		panic(err)
	}
	result, err := c.postRepo.UpdatePost(id, post)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Update"})
		} else {
		ctx.JSON(http.StatusOK, result)
	}
}