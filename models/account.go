package models

import (
  "time"
)

type Account struct {
  ID        uint       `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
  Email     string     `gorm:"type:VARCHAR(256);NOT NULL;UNIQUE" json:"email"`
  Password  string     `gorm:"type:VARCHAR(64);NOT NULL" json:"-"`
  CreatedAt time.Time  `gorm:"NOT NULL" json:"created_at"`
  UpdatedAt time.Time  `gorm:"NOT NULL" json:"updated_at"`
  Galleries []Gallery  `json:"galleries,omitempty"`
  Photos    []Photo    `json:"photos,omitempty"`
  Reactions []Reaction `json:"reactions,omitempty"`
}
