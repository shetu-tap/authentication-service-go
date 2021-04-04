package authentication

import "github.com/gin-gonic/gin"

type AuthenticationController interface {
	CreateUser(c *gin.Context);
	LoginUser(c *gin.Context);
}
