package models

import "time"

//Author...
type Author struct{
	Id 		  string `json:"id"`
	Fullname string  `json:"fullname" binding:"required" minLength:"2" maxLength:"255" example:"John Doe"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt *time.Time `json:"updated_at"`
	DeleteAt *time.Time `json:"-"`
}
type GetAuthor struct {
	Id 		  string `json:"id"`
	Fullname string  `json:"fullname" binding:"required" minLength:"2" maxLength:"255" example:"John Doe"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt *time.Time `json:"updated_at"`
	DeleteAt *time.Time `json:"-"`
}

type CreateAuthorModel struct{
	 Fullname string  `json:"fullname" binding:"required" minLength:"2" maxLength:"50" example:"John Doe Oe"`

	// Firstname string  `json:"firstname" binding:"required" minLength:"2" maxLength:"50" example:"John"`
	// Lastname  string  `json:"lastname" binding:"required" minLength:"2" maxLength:"50" example:"Doe"`
	// Middlename  string  `json:"middlename" example:"O"`
	
}

type UpdateAuthorModel struct{
	Id 		  string `json:"id"`
	Fullname string  `json:"fullname" binding:"required" minLength:"2" maxLength:"50" example:"John Doe Oe"`

	// Firstname string  `json:"firstname" binding:"required" minLength:"2" maxLength:"50" example:"John"`
	// Lastname  string  `json:"lastname" binding:"required" minLength:"2" maxLength:"50" example:"Doe"`
	// Middlename  string  `json:"middlename" example:"O"`

}
