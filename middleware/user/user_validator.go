package user_middlewares

import (
	"encoding/json"
	"fmt"
	"hotel-booking-golang-gin/initializers"
	user_interface "hotel-booking-golang-gin/interfaces/user"
	error_handler "hotel-booking-golang-gin/util/error"
	response_handler "hotel-booking-golang-gin/util/response"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

func UserValidator(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
		c.Abort()
		return
	}
	var bodyContent map[string]interface{}

	err = json.Unmarshal(bodyBytes, &bodyContent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error parsing JSON"})
		c.Abort()
		return
	}
	if bodyContent["role"] == nil {
		bodyContent["role"] = "USER"
	}
	if bodyContent["role"].(string) != "USER" && bodyContent["role"].(string) != "MERCHANT" {
		c.JSON(http.StatusBadRequest, response_handler.Error("VALIDATION_ERROR", &error_handler.ErrArg{
			Code:        "VALIDATION_ERROR",
			Description: "Role must be either MERCHANT or USER",
		}))
		c.Abort()
		return
	}

	if bodyContent["IsActive"] == nil {
		bodyContent["IsActive"] = true
	}

	if bodyContent["IsVerified"] == nil {
		bodyContent["IsVerified"] = false
	}

	user_obj := user_interface.User{
		Name:       bodyContent["name"].(string),
		Email:      bodyContent["email"].(string),
		Password:   bodyContent["password"].(string),
		Role:       bodyContent["role"].(string),
		IsActive:   bodyContent["IsActive"].(bool),
		IsVerified: bodyContent["IsVerified"].(bool),
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(&user_obj)

	if err != nil {
		c.JSON(http.StatusBadRequest, response_handler.Error("VALIDATION_ERROR", &error_handler.ErrArg{
			Code:        "VALIDATION_ERROR",
			Description: err.Error(),
		}))
		c.Abort()
		return
	}

	c.Set("User", user_obj)
	c.Next()
}

func UserLoginValidator(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
		c.Abort()
		return
	}
	var bodyContent map[string]interface{}

	err = json.Unmarshal(bodyBytes, &bodyContent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error parsing JSON"})
		c.Abort()
		return
	}

	user_obj := user_interface.UserLogin{
		Email:    bodyContent["email"].(string),
		Password: bodyContent["password"].(string),
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(&user_obj)

	if err != nil {
		c.JSON(http.StatusBadRequest, response_handler.Error("VALIDATION_ERROR", &error_handler.ErrArg{
			Code:        "VALIDATION_ERROR",
			Description: err.Error(),
		}))
		c.Abort()
		return
	}

	c.Set("User", user_obj)
	c.Next()
}

func ForRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("User").(user_interface.User)
		if user.Role != role {
			c.JSON(http.StatusUnauthorized, response_handler.Error("UNAUTHORIZED", &error_handler.ErrArg{
				Code:        "UNAUTHORIZED",
				Description: "Unauthorized",
			}))
			c.Abort()
			return
		}
		fmt.Println("Role is correct", user.Role)
		c.Next()
	}
}

func Authenticate(c *gin.Context) {
	// Get the token from the header
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, response_handler.Error("UNAUTHORIZED", &error_handler.ErrArg{
			Code:        "UNAUTHORIZED",
			Description: "Unauthorized",
		}))
		c.Abort()
		return
	}

	// Validate the token
	user, err := ValidateToken(accessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response_handler.Error("UNAUTHORIZED", &error_handler.ErrArg{
			Code:        "UNAUTHORIZED",
			Description: "Unauthorized",
		}))
		c.Abort()
		return
	}

	c.Set("User", user)
	c.Next()
}

func ValidateToken(token string) (user_interface.User, error) {
	// Validate the token
	claims := ParseAccessToken(token)

	if claims.StandardClaims.Valid() != nil {
		return user_interface.User{}, claims.StandardClaims.Valid()
	}

	userSession := user_interface.UserSession{}
	initializers.DB.Where(&user_interface.UserSession{
		AccessToken: token,
	}).First(&userSession)

	if userSession.ID == 0 {
		return user_interface.User{}, &error_handler.ErrArg{
			Code:        "SESSION_NOT_FOUND",
			Description: "User not found",
			Title:       "User not found",
		}
	}

	user := user_interface.User{}
	initializers.DB.First(&user, userSession.UserID)

	if user.ID == 0 {
		return user_interface.User{}, &error_handler.ErrArg{
			Code:        "USER_NOT_FOUND",
			Description: "User not found",
			Title:       "User not found",
		}
	}

	return user, nil
}

func ParseAccessToken(accessToken string) *user_interface.UserClaims {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &user_interface.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("secret")), nil
	})

	return parsedAccessToken.Claims.(*user_interface.UserClaims)
}

func ParseRefreshToken(refreshToken string) *jwt.StandardClaims {
	parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("refreshsecret")), nil
	})
	return parsedRefreshToken.Claims.(*jwt.StandardClaims)
}
