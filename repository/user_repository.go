package repository

import (
	"rhmn-coffe/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user entity.User) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	Delete(id string) error
	FindById(id string) (entity.User, error)
	FindByUserName(username string) (entity.User, error)
	FindAll() ([]entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) Register(user entity.User) (entity.User, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u *userRepository) Update(user entity.User) (entity.User, error) {
	if err := u.db.Save(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u *userRepository) Delete(id string) error {
	err := u.db.Where("user_id = ?", id).Delete(&entity.User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) FindById(id string) (entity.User, error) {
	var user entity.User
	err := u.db.Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u *userRepository) FindByUserName(username string) (entity.User, error) {
	var user entity.User
	err := u.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u *userRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
