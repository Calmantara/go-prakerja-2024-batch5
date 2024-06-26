package repository

import (
	"errors"
	"time"

	"github.com/Calmantara/go-prakerja-2024-batch5/sesi6/model"
	"gorm.io/gorm"
)

type StudentLocalRepo struct {
	students []*model.Student
}

func (s *StudentLocalRepo) Get() ([]*model.Student, error) {
	if len(s.students) <= 0 {
		return []*model.Student{}, nil
	}
	students := []*model.Student{}
	for _, student := range s.students {
		if student.DeletedAt != nil {
			continue
		}
		students = append(students, student)
	}
	return students, nil
}

func (s *StudentLocalRepo) Create(student *model.Student) error {
	student.ID = uint64(len(s.students) + 1)
	s.students = append(s.students, student)
	return nil
}

func (s *StudentLocalRepo) Update(id uint64, studentUpdate *model.StudentUpdate) error {
	for _, student := range s.students {
		if student.ID == id && student.DeletedAt == nil {
			student.Name = studentUpdate.Name
			student.Gender = studentUpdate.Gender
			student.DoB = studentUpdate.DoB
			return nil
		}
	}
	return errors.New("student not found")
}

func (s *StudentLocalRepo) Delete(id uint64) error {
	for _, student := range s.students {
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

type StudentPgRepo struct {
	DB *gorm.DB
}

func (s *StudentPgRepo) Get() ([]*model.Student, error) {
	students := []*model.Student{}
	err := s.DB.Debug().Find(&students).Error
	return students, err
}

func (s *StudentPgRepo) Create(student *model.Student) error {
	err := s.DB.Debug().Create(&student).Error
	return err
}

func (s *StudentPgRepo) Update(id uint64, studentUpdate *model.StudentUpdate) error {
	err := s.DB.Debug().
		Where("id = ?", id).
		Updates(&model.Student{
			Name:   studentUpdate.Name,
			DoB:    studentUpdate.DoB,
			Gender: studentUpdate.Gender,
		}).Error
	return err
}

func (s *StudentPgRepo) Delete(id uint64) error {
	err := s.DB.Debug().
		Where("id = ?", id).
		Delete(&model.Student{}).Error
	return err
}
