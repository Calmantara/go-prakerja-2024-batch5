package model

import "time"

const (
	GENDER_MALE   GenderType = "MALE"
	GENDER_FEMALE GenderType = "FEMALE"
)

type (
	GenderType string

	// Domain akan diproyeksikan kedalam bentuk stuct
	// https://gorm.io/docs/models.html
	Student struct {
		ID        uint64     `json:"id" gorm:"column:id;autoIncrement"`
		Name      string     `json:"name" gorm:"column:name"`
		Email     string     `json:"email" gorm:"column:email"`
		Gender    GenderType `json:"gender" gorm:"column:gender"`
		DoB       time.Time  `json:"dob" gorm:"column:dob;type:date"`
		DeletedAt *time.Time `json:"-" gorm:"-"`
	}

	StudentUpdate struct {
		Name   string     `json:"name"`
		Gender GenderType `json:"gender"`
		DoB    time.Time  `json:"dob"`
	}

	StudentCreate struct {
		Name   string     `json:"name"`
		Email  string     `json:"email"`
		Gender GenderType `json:"gender"`
		DoB    time.Time  `json:"dob"`
	}
)
