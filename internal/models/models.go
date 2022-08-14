package models

import (
	"database/sql"
	"time"
)

type Paste struct {
	ID        uint `gorm:"primaryKey"`
	Text      string
	Expiry    sql.NullTime
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
