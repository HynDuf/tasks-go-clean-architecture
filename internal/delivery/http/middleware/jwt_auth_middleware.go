package middleware

import (
	"net/http"
	"strings"

	"github.com/HynDuf/tasks-go-clean-architecture/pkg/tokenutil"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"error"`
}

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := tokenutil.IsAuthorized(authToken, secret)
			if authorized {
				userID, err := tokenutil.ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, ErrorResponse{Message: err.Error()})
					c.Abort()
					return
				}
				c.Set("x-user-id", userID)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Not authorized"})
		c.Abort()
	}
}
