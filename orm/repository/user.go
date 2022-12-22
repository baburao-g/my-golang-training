package repository

import (
	"orm/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	AddUser(model.User) (model.User, error)
	GetUser(int) (model.User, error)
	GetByEmail(string) (model.User, error)
	GetAllUser() ([]model.User, error)
	UpdateUser(model.User) (model.User, error)
	DeleteUser(model.User) (model.User, error)
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		connection: DB(),
	}

}

// GetUser implements UserRepository
func (db *userRepository) GetUser(id int) (user model.User, err error) {
	return user, db.connection.First(&user, id).Error
}

// GetByEmail implements UserRepository
func (db *userRepository) GetByEmail(email string) (user model.User, err error) {
	return user, db.connection.First(&user, "email=?", email).Error
}

// GetAllUser implements UserRepository
func (db *userRepository) GetAllUser() (users []model.User, err error) {
	return users, db.connection.First(&users).Error
}

// AddUser implements UserRepository
func (db *userRepository) AddUser(user model.User) (model.User, error) {
	return user, db.connection.Create(&user).Error
}

// DeleteUser implements UserRepository
func (db *userRepository) DeleteUser(user model.User) (model.User, error) {
	if err := db.connection.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	return user, db.connection.First(&user).Updates(&user).Error
}

// UpdateUser implements UserRepository
func (db *userRepository) UpdateUser(user model.User) (model.User, error) {
	if err := db.connection.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	return user, db.connection.Delete(&user).Error
}
