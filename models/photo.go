package models

import (
  "time"
)

type Photo struct {
  ID          uint       `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
  AccountId   uint       `gorm:"NOT NULL" json:"account_id"`
  GalleryId   uint       `gorm:"NOT NULL" json:"gallery_id"`
  Name        string     `gorm:"type:VARCHAR(256);NOT NULL" json:"name"`
  Description string     `gorm:"type:TEXT" json:"description,omitempty"`
  Path        string     `gorm:"NOT NULL" json:"path"`
  Size        uint       `gorm:"NOT NULL" json:"size"`
  CreatedAt   time.Time  `gorm:"NOT NULL" json:"created_at"`
  UpdatedAt   time.Time  `gorm:"NOT NULL" json:"updated_at"`
  Reactions   []Reaction `json:"reactions,omitempty"`
}
