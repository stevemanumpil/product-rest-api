package model

import "time"

type Product struct {
	Name        string    `db:"name" json:"name"`
	Price       int       `db:"price" json:"price"`
	Description string    `db:"description" json:"description"`
	Quantity    int       `db:"quantity" json:"quantity"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}