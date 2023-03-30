package model

import "time"

type Book struct {
	ID        int    `gorm:"primaryKey"`
	Name_Book string `gorm:"not null;type:varchar(50)"`
	Author    string `gorm:"not null;type:varchar(50)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BookResponse struct {
	ID        int       `json:"id"`
	Name_Book string    `json:"name_book"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type BookRequest struct {
	Name_Book string `json:"name_book"`
	Author    string `json:"author"`
}
