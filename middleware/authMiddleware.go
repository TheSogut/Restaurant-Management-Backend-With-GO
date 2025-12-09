package middleware

import (
	"fmt"
	helper "golang-restaurant-management/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("token")
		if clientToken == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No authorization header provided")})
			ctx.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "nil" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			ctx.Abort()
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Set("first_name", claims.First_name)
		ctx.Set("last_name", claims.Last_name)
		ctx.Set("uid", claims.UId)

		ctx.Next()
	}

}
