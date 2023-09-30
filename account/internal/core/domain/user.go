package domain

import (
	"github.com/google/uuid"
)

// User defines domain model and its json and db representations
type User struct {
	UID      uuid.UUID `json:"uid" db:"uid" `
	Email    string    `json:"email" db:"email" `
	Password string    `json:"-" db:"password" `
	Name     string    `json:"name" db:"name" `
	ImageURL string    `json:"imageUrl" db:"image_url" `
	Website  string    `json:"website" db:"website" `
}
