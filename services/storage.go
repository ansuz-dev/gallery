package services

import (
  "path"
)

const (
  PhotosDir = "photos"
)

var RootDir = "."

func GetPhotoDir(accountId uint) string {
  photoDir := path.Join(RootDir, string(accountId))

  return photoDir
}
