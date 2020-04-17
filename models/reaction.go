package models

import (
  "time"
)

type Reaction struct {
  ID        uint      `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
  AccountId uint      `gorm:"NOT NULL" json:"account_id"`
  PhotoId   uint      `gorm:"NOT NULL" json:"photo_id"`
  Reaction  string    `gorm:"type:VARCHAR(64);NOT NULL" json:"reaction"`
  CreatedAt time.Time `gorm:"NOT NULL" json:"created_at"`
  UpdatedAt time.Time `gorm:"NOT NULL" json:"updated_at"`
}
