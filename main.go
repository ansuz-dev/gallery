package main

import (
  "gallery/routes"
  "gallery/services"
)

func main() {
  _ = services.ConnectDB(
    "dev:root@tcp(127.0.0.1:3306)/galleries?parseTime=true",
  )
  defer services.CloseDB()

  err := services.CreateLogger(".")
  if err != nil {
    panic(err)
  }
  defer services.CloseLogger()

  g := routes.Create()
  g.Run("127.0.0.1:3000")
}
