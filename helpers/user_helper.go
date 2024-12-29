package helpers

import (
	"nganterin-go/models/dto"

	"github.com/gin-gonic/gin"
)

func GetUserData(c *gin.Context) dto.User {
	user_data, _ := c.Get("user")

	userDTO := user_data.(dto.User)

	return userDTO
}