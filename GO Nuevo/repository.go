package main

import (
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUserByID(id int) (*User, error) {
	var user User
	if r.db.First(&user, id).RecordNotFound() {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}

func (r *userRepository) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) UpdateUser(user *User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(id int) error {
	if err := r.db.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}
