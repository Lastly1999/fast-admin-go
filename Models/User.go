package Models

import "time"

type User struct {
	ID           uint
	UserName         string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
