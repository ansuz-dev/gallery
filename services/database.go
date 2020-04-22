package services

import (
  "gallery/models"
  _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDB(connection string) (err error) {

  // Create connection
  DB, err = gorm.Open("mysql", connection)
  if err != nil {
    return
  }

  err = DB.AutoMigrate(
    &models.Account{},
    &models.Gallery{},
    &models.Photo{},
    &models.Reaction{},
  ).Error
  if err != nil {
    return
  }
  DB.Model(&models.Gallery{}).AddForeignKey(
    "account_id", "accounts(id)", "CASCADE", "CASCADE",
  )
  DB.Model(&models.Photo{}).AddForeignKey(
    "account_id", "accounts(id)", "CASCADE", "CASCADE",
  )
  DB.Model(&models.Photo{}).AddForeignKey(
    "gallery_id", "galleries(id)", "CASCADE", "CASCADE",
  )
  DB.Model(&models.Reaction{}).AddForeignKey(
    "account_id", "accounts(id)", "CASCADE", "CASCADE",
  )
  DB.Model(&models.Reaction{}).AddForeignKey(
    "photo_id", "photos(id)", "CASCADE", "CASCADE",
  )

  return
}

func CloseDB() {
  if DB != nil {
    DB.Close()
  }
}

func GetAccountByID(id uint) (account *models.Account, err error) {
  Logger.Debugf("Get account by ID=[%d]", id)

  account = &models.Account{}
  err = DB.First(account, id).Error
  return
}

func GetAccountByEmail(email string) (account *models.Account, err error) {
  account = &models.Account{}
  err = DB.Where("email = ?", email).First(account).Error
  return
}

func UpdateAccount(id uint, newAccount *models.Account) (err error) {
  Logger.Debugf("Update account by ID=[%d]", id)

  account := &models.Account{}
  err = DB.First(account, id).Error
  if err != nil {
    return
  }

  if newAccount.Email != account.Email {
    account.Email = newAccount.Email

    err = DB.Save(account).Error
  }

  return
}

func DeleteAccount(id uint) (err error) {
  Logger.Debugf("Delete account by ID=[%d]", id)

  account := &models.Account{}
  err = DB.First(account, id).Error
  if err != nil {
    return
  }

  err = DB.Delete(account).Error
  return
}
