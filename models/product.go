package models

import "time"

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name" gorm:"type: varchar(255)"`
	Price       int    `json:"price" gorm:"type: int"`
	Description string `json:"description" gorm:"type: varchar(255)"`
	Stock       int    `json:"stock" gorm:"type: int"`
	Photo       string `json:"photo" gorm:"type: varchar(255)"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}
