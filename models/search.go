package models

import "time"

type SearchItem struct {
	Id   int64  `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time
}
