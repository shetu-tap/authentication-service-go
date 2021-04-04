package users

import "github.com/gin-gonic/gin"

type UserController interface {
	Me(c *gin.Context)
}
