package routes

import (
  "gallery/services"
  "github.com/gin-gonic/gin"
  "os"
  "testing"
)

var ts *gin.Engine

func TestMain(m *testing.M) {
  setup()
  code := m.Run()
  clear()

  os.Exit(code)
}

func setup() {
  _ = services.CreateLogger("..")
  _ = services.ConnectDB("dev:root@tcp(127.0.0.1:3306)/galleries?parseTime=true")

  ts = Create()
}

func clear() {
  services.DB.Close()
}
