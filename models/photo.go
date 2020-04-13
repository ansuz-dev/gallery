package models

import (
  "time"
)

type Photo struct {
  ID          uint      `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
  AccountId   uint      `gorm:"NOT NULL"`
  GalleryId   uint      `gorm:"NOT NULL"`
  Name        string    `gorm:"type:VARCHAR(256);NOT NULL"`
  Description string    `gorm:"type:TEXT"`
  Path        string    `gorm:"NOT NULL"`
  Size        uint      `gorm:"NOT NULL"`
  CreatedAt   time.Time `gorm:"NOT NULL"`
  UpdatedAt   time.Time `gorm:"NOT NULL"`
  Reactions   []Reaction
}
