package repository

import (
	"fmt"
	"github.com/Rafipratama22/mnc_test.git/entity"
	"github.com/Rafipratama22/mnc_test.git/middleware"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(newuser entity.User) (entity.User, error)
	LoginUser(email string, password string) (string, error)
	GetAllPoint(user_id uuid.UUID) (int, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (c *userRepository) RegisterUser(newuser entity.User) (entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newuser.Password), bcrypt.DefaultCost)
	if err != nil {
		return newuser, err
	}
	fmt.Println(hashedPassword)
	newuser.Password = string(hashedPassword)
	result := c.db.Create(&newuser)
	if result.Error != nil {
		return newuser, result.Error
	} else {
		return newuser, nil
	}
}

func (c *userRepository) LoginUser(email string, password string) (string, error) {
	var user entity.User
	var middleware middleware.AuthMiddleware = middleware.NewAuthMiddleware(c.db)
	var userToken string
	c.db.Where("email = ?", email).First(&user)
	fmt.Println(user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "" , err
	}
	fmt.Println("Login Success")

	userToken, err = middleware.CreateToken(user.ID)
	if err != nil {
		return "", err
	}
	fmt.Println(userToken)
	return userToken, nil
}

func (c *userRepository) GetAllPoint(user_id uuid.UUID) (int, error){
	var post []entity.Post
	result := c.db.Model(&post).Where("author_id = ?", user_id).Find(&post)
	if result.Error != nil {
		return 0, result.Error
	} else {
		var total int
		for _, value := range post {
			total += value.Point
		}
		return total, nil
	}
}