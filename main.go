package main

import (
  "fmt"
  "gallery/routes"
  "gallery/services"
)

func main() {
  _ = services.ConnectDB(
    "dev:root@tcp(127.0.0.1:3306)/galleries?parseTime=true",
  )
  fmt.Println("Connected !")

  g := routes.Create()
  g.Run("127.0.0.1:3000")
}
