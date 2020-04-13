package services

import (
  "gallery/models"
  _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
)

func ConnectDB(connection string) *gorm.DB {

  // Create connection
  db, err := gorm.Open("mysql", connection)
  if err != nil {
    panic(err)
  }

  err = db.AutoMigrate(
    &models.Account{},
    &models.Gallery{},
    &models.Photo{},
    &models.Reaction{},
  ).Error
  if err != nil {
    panic(err)
  }
  db.Model(&models.Gallery{}).AddForeignKey(
    "account_id", "accounts(id)", "CASCADE", "CASCADE",
  )
  db.Model(&models.Photo{}).AddForeignKey(
    "account_id", "accounts(id)", "CASCADE", "CASCADE",
  )
  db.Model(&models.Photo{}).AddForeignKey(
    "gallery_id", "galleries(id)", "CASCADE", "CASCADE",
  )
  db.Model(&models.Reaction{}).AddForeignKey(
    "account_id", "accounts(id)", "CASCADE", "CASCADE",
  )
  db.Model(&models.Reaction{}).AddForeignKey(
    "photo_id", "photos(id)", "CASCADE", "CASCADE",
  )

  return db
}
