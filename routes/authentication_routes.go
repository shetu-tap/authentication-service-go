package routes

import (
	"authentication/controllers/authentication"
	authentication_repository "authentication/repository/authentication"
	authentication_service "authentication/services"
	"github.com/gin-gonic/gin"
)

func AddAuthenticationRoutes(r *gin.Engine) *gin.Engine {
	authenticationRepository := authentication_repository.GetAuthenticationRepository()
	authenticationService := authentication_service.GetAuthenticationService(authenticationRepository)
	authenticationController := authentication.GetAuthenticationController(authenticationService)
	authenticationRoutes := r.Group("/auth")
	authenticationRoutes.POST("/register", authenticationController.CreateUser)
	authenticationRoutes.POST("/login", authenticationController.LoginUser)
	return r
}
