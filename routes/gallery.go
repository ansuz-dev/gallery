package routes

import (
  "errors"
  "gallery/models"
  "gallery/services"
  "github.com/gin-gonic/gin"
)

func CreateGallery(ctx *gin.Context) {
  services.Logger.Info("Create new gallery")

  accountId, exists := ctx.Get("account_id")
  if !exists {
    ctx.AbortWithError(401, errors.New("Unauthorized"))
    return
  }

  gallery := &models.Gallery{}
  if err := ctx.BindJSON(gallery); err != nil {
    ctx.AbortWithError(400, errors.New("Invalid input data"))
    return
  }

  err := services.CreateGallery(accountId.(uint), gallery)
  if err != nil {
    ctx.AbortWithError(400, errors.New("Cannot create gallery"))
    return
  }

  ctx.JSON(200, gallery)
}

func GetGalleries(ctx *gin.Context) {
  services.Logger.Info("Get galleries")

  accountId, exists := ctx.Get("account_id")
  if !exists {
    ctx.AbortWithError(401, errors.New("Unauthorized"))
    return
  }

  galleries, err := services.GetGalleries(accountId.(uint))
  if err != nil {
    ctx.AbortWithError(400, errors.New("Cannot get galleries"))
    return
  }

  ctx.JSON(200, galleries)
}

func GetAllGalleries(ctx *gin.Context) {

}

func GetGallery(ctx *gin.Context) {

}

func UpdateGallery(ctx *gin.Context) {

}

func DeleteGallery(ctx *gin.Context) {

}
