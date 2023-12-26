package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/garylow2001/GossipGo-Backend/initializers"
	"github.com/garylow2001/GossipGo-Backend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(context *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	if context.Bind(&body) != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to read body"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to hash password"})
		return
	}

	user := models.User{Username: body.Username}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to create user, please use another username"})
		return
	}

	auth := models.Auth{UserID: user.ID, Password: string(hash)}

	result = initializers.DB.Create(&auth)

	if result.Error != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to create auth for user"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "User created"})
}

func Login(context *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	if context.Bind(&body) != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to read body"})
		return
	}

	var auth models.Auth
	var user models.User

	result := initializers.DB.Where("username = ?", body.Username).First(&user)

	if result.Error != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to find user"})
		return
	}

	result = initializers.DB.Where("user_id = ?", user.ID).First(&auth)

	if result.Error != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to find auth for user"})
		return
	}

	userID := auth.UserID

	err := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(body.Password))

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong password"})
		return
	}

	var (
		token       *jwt.Token
		tokenString string
	)

	token = jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"sub": userID,
			"exp": time.Now().Add(time.Hour * 72).Unix(),
		})

	tokenString, err = token.SignedString([]byte(os.Getenv("JWT_KEY")))

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
		return
	}

	auth.Token = tokenString

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Login success", "token": tokenString})
}
