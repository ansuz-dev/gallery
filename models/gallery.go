package models

import (
  "time"
)

type Gallery struct {
  ID        uint      `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
  AccountId uint      `gorm:"NOT NULL" json:"account_id"`
  Name      string    `gorm:"type:VARCHAR(256);NOT NULL" json:"name"`
  CreatedAt time.Time `gorm:"NOT NULL" json:"created_at"`
  UpdatedAt time.Time `gorm:"NOT NULL" json:"updated_at"`
  Photos    []Photo   `json:"photos,omitempty"`
  Account   Account   `json:"account,omitempty"`
}
