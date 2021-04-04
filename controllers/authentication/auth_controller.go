package authentication

import (
	"authentication/dto/errors"
	"authentication/interfaces/authentication"
	authenticationdto "authentication/dto/authentication"
	"authentication/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authenticationController struct {
	authenticationService authentication.AuthenticationService
}

func (authenticationController *authenticationController) CreateUser(c *gin.Context) {
	var registrationBody authenticationdto.RegistrationBody
	err := c.ShouldBindJSON(&registrationBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Message: "Data not formatted",
			Code: http.StatusBadRequest,
		})
		return
	}
	user, e := authenticationController.authenticationService.CreateUser(registrationBody)
	if e != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, user.ToUserJSON())
}

func (authenticationController *authenticationController) LoginUser(c *gin.Context) {
	var loginBody authenticationdto.LoginBody
	err := c.ShouldBindJSON(&loginBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Message: "Bad request",
			Code:    http.StatusBadRequest,
		})
		return
	}
	user, er := authenticationController.authenticationService.FindUserByUsername(loginBody.Username)
	if er != nil {
		c.JSON(http.StatusNotFound, errors.ErrorResponse{
			Message: fmt.Sprintf("User with username %v not found", loginBody.Username),
			Code:    http.StatusNotFound,
		})
		return
	}
	passwordCheck := user.CheckPassword(loginBody.Password)
	if !passwordCheck {
		c.JSON(http.StatusForbidden, errors.ErrorResponse{
			Message: "Password not matched",
			Code:    http.StatusForbidden,
		})
		return
	}
	jwtUtil := utils.GetJwtUtil()
	token := jwtUtil.CreateToken(user)
	if token == nil {
		c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Message: "Error in token generation",
			Code:    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, authenticationdto.LoginResponse{Jwt: *token})
}

func GetAuthenticationController(service authentication.AuthenticationService) authentication.AuthenticationController{
	return &authenticationController{
		authenticationService: service,
	}
}
