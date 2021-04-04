package routes

import (
	user_controller "authentication/controllers/users"
	"authentication/middlewares/authentication"
	authentication_repository "authentication/repository/authentication"
	authentication_service "authentication/services"
	"github.com/gin-gonic/gin"
)

func AddUserRoutes(r *gin.Engine) *gin.Engine {
	authenticationRepository := authentication_repository.GetAuthenticationRepository()
	authenticationService := authentication_service.GetAuthenticationService(authenticationRepository)
	userController := user_controller.GetUserController(authenticationService)
	userRoutes := r.Group("/users").Use(authentication.AuthenticationMiddleware())
	userRoutes.GET("/me", userController.Me)
	return r
}
