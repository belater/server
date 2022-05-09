package entity

import "github.com/google/uuid"

type Contact struct {
	Id            uuid.UUID
	CreatedAt     string
	UpdatedAt     string
	EventId       string
	Surname       string
	ContactNumber string
	Email         string
}
type ContactInput struct {
	EventId       string
	Surname       string
	ContactNumber string
	Email         string
}
