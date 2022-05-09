package entity

import "github.com/google/uuid"

type StatusEventType string

const (
	StatusEventActive  StatusEventType = "active"
	StatusEventDecline StatusEventType = "decline"
	StatusEventDone    StatusEventType = "done"
)

type Event struct {
	Id           uuid.UUID
	CreatedAt    string
	UpdatedAt    string
	UserId       string
	CategoryId   string
	EventName    string
	IsTeam       bool
	ReminderDate string
	IsRepeated   bool
	RepeatedType string
	Status       StatusEventType // active, decline, done
	SendEmail    bool
	SendWa       bool
	SsetCalendar bool
	MeetingLink  string
}

type EventInput struct {
}
