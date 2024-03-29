package controllers

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/garylow2001/GossipGo-Backend/configs"
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
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{Username: body.Username}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Failed to create user, please use another username"})
		return
	}

	auth := models.Auth{UserID: user.ID, Password: string(hash)}

	result = initializers.DB.Create(&auth)

	if result.Error != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Failed to create auth for user"})
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
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	var auth models.Auth
	var user models.User

	result := initializers.DB.Where("username = ?", body.Username).First(&user)

	if result.Error != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Failed to find user"})
		return
	}

	result = initializers.DB.Where("user_id = ?", user.ID).First(&auth)

	if result.Error != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Failed to find auth for user"})
		return
	}

	userID := auth.UserID

	err := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(body.Password))

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Wrong password"})
		return
	}

	var (
		key         *ecdsa.PrivateKey
		token       *jwt.Token
		tokenString string
	)

	key = GetPrivateKey()
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JWT key"})
		return
	}

	token = jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"sub": userID,
			"exp": time.Now().Add(configs.JWTExpirationTime).Unix(),
		})

	tokenString, err = token.SignedString(key)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	if result.Error != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update auth token"})
		return
	}

	context.SetSameSite(http.SameSiteNoneMode)
	context.SetCookie(
		"Authorization",
		tokenString,
		configs.JWTExpirationTimeInSeconds,
		"",
		"gossipgo-backend.onrender.com",
		true,
		true,
	)
	context.IndentedJSON(http.StatusOK, gin.H{"user": user})
}

func Logout(context *gin.Context) {
	context.SetSameSite(http.SameSiteNoneMode)
	context.SetCookie(
		"Authorization",
		"",
		-1,
		"",
		"gossipgo-backend.onrender.com",
		true,
		true,
	)
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func GetPrivateKey() *ecdsa.PrivateKey {
	keyBytes, err := os.ReadFile("ecdsa_private_key.pem")
	if err != nil {
		log.Fatalf("Unable to read private key: %v", err)
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		log.Fatalf("Unable to decode PEM block containing private key")
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		log.Fatalf("Unable to parse ECDSA private key: %v", err)
	}

	return key
}

/*
* Validate validates the JWT token and returns the user if the token is valid.
* This is used by the frontend to retrieve the currently logged in user based on
* the authentication token stored in cookies.
  - @param context The context of the request.
*/
func Validate(context *gin.Context) {
	user := context.MustGet("user").(models.User)

	if err := initializers.DB.
		Preload("Threads").
		Preload("Comments").
		Preload("ThreadLikes").
		Preload("CommentLikes").
		First(&user, user.ID).
		Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"user": user})
}
