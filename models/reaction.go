package models

import (
  "time"
)

type Reaction struct {
  ID        uint      `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
  AccountId uint      `gorm:"NOT NULL"`
  PhotoId   uint      `gorm:"NOT NULL"`
  Reaction  string    `gorm:"type:VARCHAR(64);NOT NULL"`
  CreatedAt time.Time `gorm:"NOT NULL"`
  UpdatedAt time.Time `gorm:"NOT NULL"`
}
