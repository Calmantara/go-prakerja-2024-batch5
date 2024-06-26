package service

import (
	"github.com/Calmantara/go-prakerja-2024-batch5/sesi6/model"
	"github.com/Calmantara/go-prakerja-2024-batch5/sesi6/repository"
)

type UserService struct {
	UserLocalRepo *repository.StudentLocalRepo
	UserPgRepo    *repository.StudentPgRepo
}

func (u *UserService) Get() ([]*model.Student, error) {
	return u.UserPgRepo.Get()
}

func (u *UserService) Create(student *model.Student) error {
	return u.UserPgRepo.Create(student)
}

func (u *UserService) Update(id uint64, student *model.StudentUpdate) error {
	return u.UserPgRepo.Update(id, student)
}

func (u *UserService) Delete(id uint64) error {
	return u.UserPgRepo.Delete(id)
}
