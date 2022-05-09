package entity

import "github.com/google/uuid"

type StatusMembershipType string

const (
	StatusMembershipPaid    StatusMembershipType = "paid"
	StatusMembershipFree    StatusMembershipType = "free"
	StatusMembershipExpired StatusMembershipType = "expired"
)

type Membership struct {
	Id          uuid.UUID
	CreatedAt   string
	UpdatedAt   string
	UserId      string
	Status      string // paid, free, expired
	ExpiredDate string
}
type MembershipInput struct {
	UserId      string
	Status      string
	ExpiredDate string
}
