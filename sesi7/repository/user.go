package repository

import (
	"errors"
	"time"

	"github.com/Calmantara/go-prakerja-2024-batch5/sesi7/model"
	"gorm.io/gorm"
)

type UserLocalRepo struct {
	users []*model.User
}

func (s *UserLocalRepo) Get() ([]*model.User, error) {
	if len(s.users) <= 0 {
		return []*model.User{}, nil
	}
	users := []*model.User{}
	for _, student := range s.users {
		if student.DeletedAt != nil {
			continue
		}
		users = append(users, student)
	}
	return users, nil
}

func (s *UserLocalRepo) Create(student *model.User) error {
	student.ID = uint64(len(s.users) + 1)
	s.users = append(s.users, student)
	return nil
}

func (s *UserLocalRepo) Update(id uint64, studentUpdate *model.UserUpdate) error {
	for _, student := range s.users {
		if student.ID == id && student.DeletedAt == nil {
			student.Name = studentUpdate.Name
			student.Gender = studentUpdate.Gender
			student.DoB = studentUpdate.DoB
			return nil
		}
	}
	return errors.New("student not found")
}

func (s *UserLocalRepo) Delete(id uint64) error {
	for _, student := range s.users {
		if student.ID == id && student.DeletedAt == nil {
			tn := time.Now()
			student.DeletedAt = &tn
			return nil
		}
	}
	return errors.New("student not found")
}

// BELAJAR DATABASE: MySQL, PostgreSQL
// BELAJAR DDL/DML (membuat database, schema, table, dan memasukkan data)
// BELAJAR SQL Golang

type UserPgRepo struct {
	DB *gorm.DB
}

func (s *UserPgRepo) Get() ([]*model.User, error) {
	users := []*model.User{}
	err := s.DB.Debug().Find(&users).Error
	return users, err
}

func (s *UserPgRepo) Create(student *model.User) error {
	err := s.DB.Debug().Create(&student).Error
	return err
}

func (s *UserPgRepo) Update(id uint64, studentUpdate *model.UserUpdate) error {
	err := s.DB.Debug().
		Where("id = ?", id).
		Updates(&model.User{
			Name:   studentUpdate.Name,
			DoB:    studentUpdate.DoB,
			Gender: studentUpdate.Gender,
		}).Error
	return err
}

func (s *UserPgRepo) Delete(id uint64) error {
	err := s.DB.Debug().
		Where("id = ?", id).
		Delete(&model.User{}).Error
	return err
}

func (s *UserPgRepo) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := s.DB.Debug().Where("email = ?", email).Find(&user).Error
	return user, err
}
