package services

import (
  "fmt"
  "github.com/dgrijalva/jwt-go"
  "time"
)

const (
  hmacSecret = "iQasdq4MMdY2wxZCpAm1SxpbkGQopM4wx9QLgtVaHfjGCavuMLcuAZG6CvFxJaMd"
)

type TokenClaims struct {
  jwt.StandardClaims
  AccountID uint `json:"account_id"`
}

func CreateToken(accountId uint) (tokenStr string, err error) {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
    AccountID: accountId,
    StandardClaims: jwt.StandardClaims{
      Issuer:    "Gallery",
      ExpiresAt: time.Now().Unix() + 24*60*60, // will be expired in 1 day
      NotBefore: time.Now().Unix(),
    },
  })

  tokenStr, err = token.SignedString([]byte(hmacSecret))
  return
}

func VerifyToken(tokenStr string) (claims *TokenClaims, err error) {
  token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      fmt.Println("Invalid signing method")
      return nil, fmt.Errorf("Invalid token")
    }

    return []byte(hmacSecret), nil
  })

  if err != nil {
    fmt.Println("err:", err)
    return
  }

  fmt.Println("token:", token)

  if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
    return claims, nil
  } else {
    fmt.Println("claims:", claims)

    return nil, fmt.Errorf("Invalid token")
  }
}