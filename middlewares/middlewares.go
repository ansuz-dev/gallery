package middlewares

import (
  "errors"
  "gallery/services"
  "github.com/gin-gonic/gin"
  "strings"
)

// Check authentication
func RequireAuthentication() gin.HandlerFunc {
  return func(ctx *gin.Context) {
    // get token from Authorization header
    token := ctx.GetHeader("Authorization")

    arr := strings.Split(token, " ")
    if len(arr) != 2 {
      ctx.AbortWithError(401, errors.New("Unauthorized"))
      return
    }

    if arr[0] != "Bearer" {
      ctx.AbortWithError(401, errors.New("Unauthorized"))
      return
    }

    token = arr[1]

    claims, err := services.VerifyToken(token)
    if err != nil {
      ctx.AbortWithError(401, errors.New("Unauthorized"))
      return
    }

    // after got claims, set it as `account_id` in context
    ctx.Set("account_id", claims.AccountID)

    ctx.Next()
  }
}
