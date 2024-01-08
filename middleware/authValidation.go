package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/garylow2001/GossipGo-Backend/controllers"
	"github.com/garylow2001/GossipGo-Backend/initializers"
	"github.com/garylow2001/GossipGo-Backend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware(context *gin.Context) {
	// Get token from header
	tokenString, err := context.Cookie("Authorization")

	if err != nil || tokenString == "" {
		context.AbortWithStatusJSON(401, gin.H{"error": "No token provided"})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check if signing method is valid
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return controllers.GetPrivateKey().Public(), nil
	})

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Error parsing token"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := claims["exp"].(float64)
		user_id := claims["sub"]

		if time.Now().After(time.Unix(int64(exp), 0)) {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
		}

		var user models.User
		initializers.DB.First(&user, user_id)

		if user.ID == 0 {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found, please log in"})
			return
		}

		context.Set("user", user)
		context.Next()
	} else {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
}
