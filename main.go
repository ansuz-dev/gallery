package main

import (
  "fmt"
  "gallery/services"
)

func main() {
  _ = services.ConnectDB(
    "dev:root@tcp(127.0.0.1:3306)/galleries?parseTime=true",
  )
  fmt.Println("Connected !")

  // account := models.Account{
  //   ID:        1,
  //   Email:     "john.doe@example.com",
  //   Password:  "123456",
  //   CreatedAt: time.Now(),
  // }

  // fmt.Printf("account: %#v\n", account)

  // var imageSrc = "./cat.jpg"
  // sizes, err := services.ResizeAll(imageSrc)
  // if err != nil {
  //   fmt.Println(err)
  // } else {
  //   fmt.Println(sizes)
  // }

}
