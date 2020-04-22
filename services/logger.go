package services

import (
  "github.com/gin-gonic/gin"
  "github.com/sirupsen/logrus"
  "github.com/sirupsen/logrus/hooks/writer"
  "io/ioutil"
  "os"
  "path"
  "time"
)

const (
  ErrorLog = "error.log"
  DebugLog = "debug.log"
)

var Logger *logrus.Logger
var errorFile *os.File
var debugFile *os.File

func CreateLogger(rootDir string) (err error) {
  dir := rootDir
  if dir == "" {
    dir, err = os.Getwd()
    if err != nil {
      return
    }
  }

  errorfile, err := os.OpenFile(path.Join(dir, ErrorLog), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
  if err != nil {
    return
  }

  debugfile, err := os.OpenFile(path.Join(dir, DebugLog), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
  if err != nil {
    return
  }

  Logger = logrus.New()
  Logger.SetOutput(ioutil.Discard)
  Logger.SetLevel(logrus.DebugLevel)

  Logger.AddHook(&writer.Hook{
    Writer: errorfile,
    LogLevels: []logrus.Level{
      logrus.WarnLevel,
      logrus.ErrorLevel,
      logrus.FatalLevel,
      logrus.PanicLevel,
    },
  })

  Logger.AddHook(&writer.Hook{
    Writer: debugfile,
    LogLevels: []logrus.Level{
      logrus.DebugLevel,
      logrus.InfoLevel,
    },
  })

  return
}

func LogToFile() gin.HandlerFunc {
  return func(ctx *gin.Context) {
    startTime := time.Now()

    // Process request
    ctx.Next()

    endTime := time.Now()
    latencyTime := endTime.Sub(startTime)

    Logger.Infof(
      "| %3d | %5v | %10s | %s | %s |",
      ctx.Writer.Status(),
      latencyTime,
      ctx.ClientIP(),
      ctx.Request.Method,
      ctx.Request.RequestURI,
    )
  }
}

func CloseLogger() {
  if errorFile != nil {
    errorFile.Close()
  }

  if debugFile != nil {
    debugFile.Close()
  }
}
