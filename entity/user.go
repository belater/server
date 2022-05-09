package entity

import "github.com/google/uuid"

type User struct {
	Id          uuid.UUID
	CreatedAt   string
	UpdatedAt   string
	FullName    string
	Surname     string
	Email       string
	PhoneNumber string
	Password    string
	Active      bool   // true if user is active
	Role        string // admin, user
	Slug        string // genereate unique slug : fullname: afista.pratama
}

type UserOutput struct {
	Id          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	FullName    string `json:"full_name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Active      bool   `json:"active"`
	Status      string `json:"status"`
	Slug        string `json:"slug"`
}

type UserLogin struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type UserRegister struct {
	FullName    string `json:"full_name" binding:"required"`
	Surname     string `json:"surname" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}
