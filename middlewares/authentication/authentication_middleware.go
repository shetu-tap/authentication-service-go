package authentication

import (
	"authentication/dto/errors"
	"authentication/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")
		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, errors.ErrorResponse{
				Message: "Unauthorized",
				Code:    http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		extractedToken := strings.Split(authorizationHeader, "Bearer ")
		if len(extractedToken) == 2 {
			authorizationHeader = strings.TrimSpace(extractedToken[1])
		} else {
			c.JSON(http.StatusUnauthorized, errors.ErrorResponse{
				Message: "Header not in right format",
				Code:    http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		jwtUtil := utils.GetJwtUtil()

		token, e := jwtUtil.ValidateToken(authorizationHeader)
		if e != nil {
			c.JSON(http.StatusForbidden, errors.ErrorResponse{
				Message: "Unable to validate token",
				Code:    http.StatusForbidden,
			})
			c.Abort()
			return
		}

		claims := token.Claims.(*jwt.StandardClaims)

		if claims.ExpiresAt < time.Now().Unix() {
			c.JSON(http.StatusForbidden, errors.ErrorResponse{
				Message: "Token expired",
				Code:    http.StatusForbidden,
			})
			c.Abort()
			return
		}
		c.Set("username", claims.Subject)
		c.Next()
	}
}
