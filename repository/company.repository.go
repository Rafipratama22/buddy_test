package repository

import (
	"fmt"
	"github.com/Rafipratama22/mnc_test.git/entity"
	"github.com/Rafipratama22/mnc_test.git/middleware"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CompanyRepository interface {
	AllUserRegister() ([]entity.User, error)
	AllUserLogin() ([]entity.User, error)
	AllUserPoint() (int, error)
	DetailUser(user_id uuid.UUID) (entity.User, error)
	DetailUserPoint(user_id uuid.UUID) (int, error)
	PostPoint(id int, point int) (entity.Post, error)
	AllPost() ([]entity.Post, error)
	DetailPost(id int) (entity.Post, error)
}

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepo(db *gorm.DB) CompanyRepository {
	return &companyRepository{
		db: db,
	}
}

func (c *companyRepository) LoginCompany(email string, password string) (string, error) {
	var company entity.Company
	var middleware middleware.AuthMiddleware = middleware.NewAuthMiddleware(c.db)
	var companyToken string
	c.db.Where("email = ?", email).First(&company)
	err := bcrypt.CompareHashAndPassword([]byte(company.Password), []byte(password))
	if err != nil {
		return "", err
	}
	companyToken, err = middleware.CreateToken(company.ID)
	if err != nil {
		return "", err
	}
	return companyToken, nil
}

func (c *companyRepository) AllUserRegister() ([]entity.User, error){
	var users []entity.User
	result := c.db.Model(&users).Find(&users)
	if result.Error != nil {
		return users, result.Error
	} else {
		return users, nil
	}
}

func (c *companyRepository) AllUserLogin() ([]entity.User, error) {
	var users []entity.User
	result := c.db.Model(&users).Where("is_active = ?", true).Find(&users)
	if result != nil {
		return users, result.Error
	} else {
		return users, nil
	}
}

func (c *companyRepository) AllUserPoint() (int, error) {
	var post []entity.Post
	result := c.db.Find(&post)
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

func (c *companyRepository) DetailUser(user_id uuid.UUID) (entity.User, error){
	var user entity.User
	result := c.db.Model(&user).Where("id = ?", user_id).First(&user)
	if result.Error != nil {
		return user, result.Error
	} else {
		return user, nil
	}
}

func (c *companyRepository) DetailUserPoint(user_id uuid.UUID) (int, error) {
	var posts []entity.Post
	result := c.db.Model(&posts).Where("author_id = ?", user_id).Find(&posts)
	if result.Error != nil {
		return 0, result.Error
	} else {
		var total int
		for _, value := range posts {
			total += value.Point
		}
		return total, nil
	}
}

func (c *companyRepository) PostPoint(id int, point int) (entity.Post, error) {
	var post entity.Post
	result := c.db.Model(&post).Where("id = ?", id).Find(&post)
	if result.Error != nil {
		return post, result.Error
	}
	post.Point += point
	resulted := c.db.Model(&post).Where("id = ?", id).Update("point", post.Point)
	if resulted.Error != nil {
		fmt.Println(resulted.Error)
		return post, resulted.Error
	}
	return post, nil
}

func (c *companyRepository) AllPost() ([]entity.Post, error) {
	var posts []entity.Post
	result := c.db.Model(&posts).Find(&posts)
	if result.Error != nil {
		return posts, result.Error
	} else {
		return posts, nil
	}
}

func (c *companyRepository) DetailPost(id int) (entity.Post, error) {
	var post entity.Post
	result := c.db.Model(&post).Where("id = ?", id).First(&post)
	if result.Error != nil {
		return post, result.Error
	} else {
		return post, nil
	}
} 