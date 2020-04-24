package routes

import (
  "errors"
  "gallery/models"
  "gallery/services"
  "github.com/gin-gonic/gin"
)

type Credential struct {
  Name     string `json:"name"`
  Email    string `json:"email"`
  Password string `json:"password"`
}

func Registration(ctx *gin.Context) {
  services.Logger.Info("Register with email and password")

  cred := &Credential{}

  if err := ctx.BindJSON(cred); err != nil {
    ctx.AbortWithError(400, errors.New("Invalid email or password"))
    return
  }

  account, err := services.Register(cred.Name, cred.Email, cred.Password)
  if err != nil {
    ctx.AbortWithError(400, errors.New("Cannot register the account"))
    return
  }

  ctx.JSON(200, account)
}

func Authentication(ctx *gin.Context) {
  services.Logger.Info("Authenticate by email and password")

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

  services.Logger.Infof("Get account information by id=[%d]", accountId)

  account, err := services.GetAccountByID(accountId.(uint))
  if err != nil {
    ctx.AbortWithError(404, errors.New("Account not found"))
    return
  }

  ctx.JSON(200, account)
}

func UpdateAccount(ctx *gin.Context) {
  services.Logger.Info("Update account information")

  accountId, exists := ctx.Get("account_id")
  if !exists {
    ctx.AbortWithError(401, errors.New("Unauthorized"))
    return
  }

  newAccount := &models.Account{}
  if err := ctx.BindJSON(newAccount); err != nil {
    ctx.AbortWithError(400, errors.New("Invalid data"))
    return
  }

  err := services.UpdateAccount(accountId.(uint), newAccount)
  if err != nil {
    ctx.AbortWithError(400, errors.New("Cannot update account information"))
    return
  }

  ctx.Status(200)
}

func DeleteAccount(ctx *gin.Context) {
  accountId, exists := ctx.Get("account_id")
  if !exists {
    ctx.AbortWithError(401, errors.New("Unauthorized"))
    return
  }

  services.Logger.Infof("Delete account by id=[%d]", accountId)

  err := services.DeleteAccount(accountId.(uint))
  if err != nil {
    ctx.AbortWithError(404, errors.New("Account not found"))
    return
  }

  ctx.Status(200)
}
