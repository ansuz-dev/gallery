package services

import (
  "errors"
  "gallery/models"
)

func Authenticate(email string, password string) (tokenStr string, err error) {
  account, err := GetAccountByEmail(email)
  if err != nil {
    return
  }

  // check matched password
  // for example, by bcrypt library
  if account.Password != password {
    return "", errors.New("Invalid email or password")
  }

  // create a new token
  tokenStr, err = CreateToken(account.ID)

  return
}

func Register(email string, password string) (account *models.Account, err error) {
  account = &models.Account{
    Email:    email,
    Password: password,
  }

  err = DB.Create(account).Error
  return
}
