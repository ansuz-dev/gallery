package models

import (
  "time"
)

type Account struct {
  ID        uint      `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
  Email     string    `gorm:"type:VARCHAR(256);NOT NULL;UNIQUE"`
  Password  string    `gorm:"type:VARCHAR(64);NOT NULL"`
  CreatedAt time.Time `gorm:"NOT NULL"`
  UpdatedAt time.Time `gorm:"NOT NULL"`
  Galleries []Gallery
  Photos    []Photo
  Reactions []Reaction
}
