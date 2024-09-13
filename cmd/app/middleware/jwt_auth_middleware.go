package middleware

//
//import (
//	"github.com/gin-gonic/gin"
//	"go-shift/cmd/app/domain/dto/system"
//	"go-shift/cmd/app/util"
//	"net/http"
//	"strings"
//)
//
//func JwtAuthMiddleware(secret string) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		authHeader := c.Request.Header.Get("Authorization")
//		t := strings.Split(authHeader, " ")
//		if len(t) == 2 {
//			authToken := t[1]
//			authorized, err := util.IsAuthorized(authToken, secret)
//			if authorized {
//				userID, err := util.ExtractIDFromToken(authToken, secret)
//				if err != nil {
//					c.JSON(http.StatusUnauthorized, system.AppError{Message: err.Error()})
//					c.Abort()
//					return
//				}
//				c.Set("x-user-id", userID)
//				c.Next()
//				return
//			}
//			c.JSON(http.StatusUnauthorized, system.AppError{Message: err.Error()})
//			c.Abort()
//			return
//		}
//		c.JSON(http.StatusUnauthorized, system.AppError{Message: "Not authorized"})
//		c.Abort()
//	}
//}
