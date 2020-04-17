package routes

import (
  "errors"
  "gallery/services"
  "github.com/gin-gonic/gin"
)

type Credential struct {
  Email    string `json:"email"`
  Password string `json:"password"`
}

func Authentication(ctx *gin.Context) {
  cred := &Credential{}

  if err := ctx.BindJSON(cred); err != nil {
    ctx.AbortWithError(401, errors.New("Invalid email or password"))
    return
  }

  token, err := services.Authenticate(cred.Email, cred.Password)
  if err != nil {
    ctx.AbortWithError(401, errors.New("Invalid email or password"))
    return
  }

  ctx.String(200, token)
}

func GetAccount(ctx *gin.Context) {
  accountId, exists := ctx.Get("account_id")
  if !exists {
    ctx.AbortWithError(401, errors.New("Unauthorized"))
    return
  }

  account, err := services.GetAccountByID(accountId.(uint))
  if err != nil {
    ctx.AbortWithError(404, errors.New("Account not found"))
    return
  }

  ctx.JSON(200, account)
}

func UpdateAccount(ctx *gin.Context) {

}

func DeleteAccount(ctx *gin.Context) {

}
