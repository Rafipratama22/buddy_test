package repository

import (
	"github.com/Rafipratama22/mnc_test.git/entity"
	"github.com/Rafipratama22/mnc_test.git/middleware"
	// "github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRepository interface {
	RegisterUser(user entity.User) (entity.User, error)
	RegisterCompany(company entity.Company) (entity.Company, error)
	LoginUser(email string, password string) (string, error)
	LoginCompany(email string, password string) (string, error)
	LogOutUser(id string) error
	LogOutCompany(id string) error
}

type loginRepository struct {
	db *gorm.DB
}

func NewLoginRepo(db *gorm.DB) LoginRepository {
	return &loginRepository{
		db: db,
	}
}


func (c *loginRepository) RegisterUser(user entity.User) (entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hashedPassword)
	result := c.db.Create(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (c *loginRepository) RegisterCompany(company entity.Company) (entity.Company, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(company.Password), bcrypt.DefaultCost)
	if err != nil {
		return company, err
	}
	company.Password = string(hashedPassword)
	result := c.db.Create(&company)
	if result.Error != nil {
		return company, result.Error
	}
	return company, nil
}

func (c *loginRepository) LoginUser(email string, password string) (string, error) {
	var user entity.User
	var middleware middleware.AuthMiddleware = middleware.NewAuthMiddleware(c.db)
	var userToken string
	c.db.Where("email = ?", email).First(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "" , err
	}
	userToken, err = middleware.CreateToken(user.ID)
	if err != nil {
		return "", err
	}
	result := c.db.Model(&user).Update("is_active", true)
	if result.Error != nil {
		return "", result.Error
	}
	return userToken, nil
}

func (c *loginRepository) LoginCompany(email string, password string) (string, error) {
	var company entity.Company
	var middleware middleware.AuthMiddleware = middleware.NewAuthMiddleware(c.db)
	var companyToken string
	c.db.Where("email = ?", email).First(&company)
	err := bcrypt.CompareHashAndPassword([]byte(company.Password), []byte(password))
	if err != nil {
		return "" , err
	}
	
	companyToken, err = middleware.CreateToken(company.ID)
	if err != nil {
		return "", err
	}
	return companyToken, nil
}

func (c *loginRepository) LogOutUser(id string) error {
	var user entity.User
	c.db.Where("id = ?", id).First(&user)
	result := c.db.Model(&user).Update("is_active", false)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *loginRepository) LogOutCompany(id string) error {
	var company entity.Company
	c.db.Where("id = ?", id).First(&company)
	result := c.db.Model(&company).Update("is_active", false)
	if result.Error != nil {
		return result.Error
	}
	return nil
}