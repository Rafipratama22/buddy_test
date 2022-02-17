package middleware

import (
	"fmt"
	"github.com/Rafipratama22/mnc_test.git/entity"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type AuthMiddleware interface {
	ValidateTokenUser(ctx *gin.Context)
	CreateToken(user_id uuid.UUID) (string, error)
	ValidateTokenCompany(ctx *gin.Context)
}

type authMiddleware struct {
	db *gorm.DB
}

func NewAuthMiddleware(db *gorm.DB) AuthMiddleware {
	return &authMiddleware{
		db: db,
	}
}

func (c *authMiddleware) CreateToken(user_id uuid.UUID) (string, error){
	godotenv.Load()
	secret_jwt := os.Getenv("JWT_KEY")
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret_jwt))
}

func (c *authMiddleware) ValidateTokenUser(ctx *gin.Context) { //Middleware function
	godotenv.Load()
	secret_jwt := os.Getenv("JWT_KEY")
	tokenString := ctx.GetHeader("Authorization") //Retrive token from header
	user := entity.User{}
	tokenMap := map[string]string{
		"user_id": "",
	}
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Error token doesnt exist",
		})
		ctx.Abort()
	}
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) { //Parse token
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", "secret")
		}
		return []byte(secret_jwt), nil
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Youre not authorized",
		})
		ctx.Abort()
	}
	if claims, err := token.Claims.(jwt.MapClaims); err && token.Valid {
		//Validate token
		fmt.Println("claims", claims)
		for key, val := range claims {
			if s, ok := val.(string); ok {
				tokenMap[key] = s
			}
		}
	}
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Youre not authorized",
		})
		ctx.Abort()
	}
	result := c.db.Model(&user).Where("id = ?", tokenMap["user_id"]).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Youre not authorized",
		})
		ctx.Abort()
	}
	ctx.Set("user_id", tokenMap["user_id"])
	ctx.Next()
}

func (c *authMiddleware) ValidateTokenCompany(ctx *gin.Context) {
	godotenv.Load()
	secret_jwt := os.Getenv("JWT_KEY")
	tokenString := ctx.GetHeader("Authorization")
	user := entity.Company{}
	tokenMap := map[string]string{
		"user_id": "",
	}
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Error token doesnt exist",
		})
		ctx.Abort()
	}
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", "secret")
		}
		return []byte(secret_jwt), nil
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Youre not authorized",
		})
		ctx.Abort()
	}
	if claims, err := token.Claims.(jwt.MapClaims); err && token.Valid {
		fmt.Println("claims", claims)
		for key, val := range claims {
			if s, ok := val.(string); ok {
				tokenMap[key] = s
			}
		}
	}
	c.db.Model(&user).Where("id = ?", tokenMap["user_id"]).First(&user)
	ctx.Set("user_id", tokenMap["user_id"])
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Youre not authorized",
		})
		ctx.Abort()
	}
	ctx.Next()
}