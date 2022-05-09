package entity

import "github.com/google/uuid"

type Category struct {
	Id        uuid.UUID
	CreatedAt string
	UpdatedAt string
	Name      string
}
type CategoryInput struct {
	Name string
}
