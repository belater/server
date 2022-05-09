package entity

import "github.com/google/uuid"

type MembershipPayment struct {
	Id            uuid.UUID
	CreatedAt     string
	UpdatedAt     string
	MembershipId  string
	PaymentMethod string
	Paid          bool
}
type MembershipPaymentInput struct {
	MembershipId  string
	PaymentMethod string
	Paid          bool
}
