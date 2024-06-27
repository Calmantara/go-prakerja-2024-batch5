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
	User struct {
		ID        uint64     `json:"id" gorm:"column:id;autoIncrement"`
		Password  string     `json:"-" gorm:"column:password"`
		Name      string     `json:"name" gorm:"column:name"`
		Email     string     `json:"email" gorm:"column:email"`
		Gender    GenderType `json:"gender" gorm:"column:gender"`
		DoB       time.Time  `json:"dob" gorm:"column:dob;type:date"`
		DeletedAt *time.Time `json:"-" gorm:"-"`
	}

	UserUpdate struct {
		Name   string     `json:"name"`
		Gender GenderType `json:"gender"`
		DoB    time.Time  `json:"dob"`
	}

	UserCreate struct {
		Name     string     `json:"name"`
		Password string     `json:"password"`
		Email    string     `json:"email"`
		Gender   GenderType `json:"gender"`
		DoB      time.Time  `json:"dob"`
	}

	UserLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
