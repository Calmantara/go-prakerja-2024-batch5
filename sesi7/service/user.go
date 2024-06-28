package service

import (
	"github.com/Calmantara/go-prakerja-2024-batch5/sesi7/model"
	"github.com/Calmantara/go-prakerja-2024-batch5/sesi7/repository"
)

// semua POINTER STRUCT yang mau masuk ke golongan UserServiceInterface
// harus mengimplementasikan method yang ada di INTERFACE
type UserServiceInterface interface {
	Get() ([]*model.User, error)
	Create(student *model.User) error
	Update(id uint64, student *model.UserUpdate) error
	Delete(id uint64) error

	GetByEmail(email string) (*model.User, error)
}

type UserService struct {
	UserLocalRepo *repository.UserLocalRepo
	UserPgRepo    *repository.UserPgRepo
}

func (u *UserService) Get() ([]*model.User, error) {
	return u.UserPgRepo.Get()
}

func (u *UserService) Create(student *model.User) error {
	return u.UserPgRepo.Create(student)
}

func (u *UserService) Update(id uint64, student *model.UserUpdate) error {
	return u.UserPgRepo.Update(id, student)
}

func (u *UserService) Delete(id uint64) error {
	return u.UserPgRepo.Delete(id)
}

func (u *UserService) GetByEmail(email string) (*model.User, error) {
	return u.UserPgRepo.GetByEmail(email)
}
