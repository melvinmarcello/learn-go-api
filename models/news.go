package models

import "time"

type News struct {
	id          int64     `db:"id"`
	title       string    `db:"title"`
	excerpt     string    `db:"excerpt"`
	description string    `db:"description"`
	publishedAt time.Time `db:"published_at"`
	createdAt   time.Time `db:"created_at"`
	updatedAt   time.Time `db:"updated_at"`
}
