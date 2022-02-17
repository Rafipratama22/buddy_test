package controller

import (
	"net/http"

	"github.com/Rafipratama22/mnc_test.git/dto"
	"github.com/Rafipratama22/mnc_test.git/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController interface {
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

// All User Point
// @Summary Retrieves all users point
// @Description Retrieves all users point
// @Tags Users
// @Accept  */*
// @Produce  json
// @Success 200 {integer} point
// @Failure 400 {object} dto.ErrorResponse
// @Router /user/point [get] 
func (c *userController) GetAllPoint(ctx *gin.Context) {
	var errResponse dto.ErrorResponse
	user_id := uuid.MustParse(ctx.MustGet("user_id").(string))
	point, err := c.userRepo.GetAllPoint(user_id)
	if err != nil {
		errResponse.Message = "Failed to get all point"
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"point": point})
	}
}

