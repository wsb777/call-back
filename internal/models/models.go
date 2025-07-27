package models

import "time"

type User struct {
	ID     int 
	Login string
	Password string 
	CreateAt time.Time
	UpdateAt time.Time
	SystemRole int // ForeignKey к SystemRole
}

type SystemRole struct {
	ID int
	Name string
}

type Room struct {
	ID int
	Members []int
	Name string
	CreateAt time.Time
	UpdateAt time.Time
}