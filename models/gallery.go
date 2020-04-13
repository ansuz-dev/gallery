package models

import (
  "time"
)

type Gallery struct {
  ID        uint      `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
  AccountId uint      `gorm:"NOT NULL"`
  Name      string    `gorm:"type:VARCHAR(256);NOT NULL"`
  CreatedAt time.Time `gorm:"NOT NULL"`
  UpdatedAt time.Time `gorm:"NOT NULL"`
  Photos    []Photo
}
