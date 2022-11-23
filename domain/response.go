package domain

import "time"

type Response struct {
	Data interface{}
	Code int
	Err  error
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type UserResponse struct {
	Id             string    `json:"id"`
	Email          string    `json:"email"`
	Address        *string   `json:"address"`
	FullName       string    `json:"full_name"`
	Phone          *string   `json:"phone"`
	ProfilePicture *string   `json:"profile_picture"`
	Role           int       `json:"role"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
