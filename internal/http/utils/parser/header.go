package parser

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/domain/entity"
)

func ParseBearerToken(c *gin.Context) string {
	tokenHeader := c.GetHeader("Authorization")
	return tokenHeader
}

func ParseUser(c *gin.Context) *entity.User {
	user, exists := c.Get("user")
	if !exists {
		return nil
	}

	userValue, ok := user.(*entity.User)
	if ok {
		return userValue
	}

	return nil
}

func ParseMunicipality(c *gin.Context) *entity.User {
	user, exists := c.Get("user")
	if !exists {
		return nil
	}

	userValue, ok := user.(*entity.User)
	if ok {
		return userValue
	}

	return nil
}

func ParseUserAuthToken(c *gin.Context) *entity.UserAuthToken {
	userToken, exists := c.Get("user_auth_token")
	if !exists {
		return nil
	}

	userTokenValue, ok := userToken.(*entity.UserAuthToken)
	if ok {
		return userTokenValue
	}

	return nil
}
