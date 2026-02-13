package model

import "time"

type Todo struct {
	ID        string    `db:"id"`
	Title     string    `db:"title"`
	Details   string    `db:"details"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}