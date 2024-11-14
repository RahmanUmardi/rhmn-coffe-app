package usecase

import (
	"fmt"
	"rhmn-coffe/entity"
	"rhmn-coffe/repository"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Register(user entity.User) (entity.User, error)
	FindAll() ([]entity.User, error)
	FindById(id string) (entity.User, error)
	FindByUsername(username string) (entity.User, error)
	FindByUsernamePassword(username string, password string) (entity.User, error)
	Update(id string, input entity.UpdateUser) (entity.User, error)
	Delete(id string) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func (u *userUsecase) Register(user entity.User) (entity.User, error) {
	if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
		return entity.User{}, fmt.Errorf("username, password and role can't be empty")
	}

	exitUser, _ := u.userRepository.FindByUserName(user.Username)
	if exitUser.Username != "" {
		return entity.User{}, fmt.Errorf("username already exist")
	}

	user.Role = "employee"

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to hash password: %v", err)
	}
	user.Password = string(hash)

	return u.userRepository.Register(user)
}

func (u *userUsecase) FindAll() ([]entity.User, error) {
	return u.userRepository.FindAll()
}

func (u *userUsecase) FindById(id string) (entity.User, error) {
	return u.userRepository.FindById(id)
}

func (u *userUsecase) FindByUsername(username string) (entity.User, error) {
	return u.userRepository.FindByUserName(username)
}

func (u *userUsecase) FindByUsernamePassword(username string, password string) (entity.User, error) {
	userExist, err := u.userRepository.FindByUserName(username)
	if err != nil {
		return entity.User{}, fmt.Errorf("user doesn't exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(password))
	if err != nil {
		return entity.User{}, fmt.Errorf("password doesn't match")
	}

	return userExist, nil
}

func (u *userUsecase) Update(id string, input entity.UpdateUser) (entity.User, error) {
	user, err := u.userRepository.FindById(id)
	if err != nil {
		return entity.User{}, fmt.Errorf("user not found: %v", err)
	}

	if strings.TrimSpace(input.Username) != "" {
		user.Username = input.Username
	}
	if strings.TrimSpace(input.Password) != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return entity.User{}, fmt.Errorf("failed to hash password: %v", err)
		}
		user.Password = string(hashedPassword)
	}
	if strings.TrimSpace(input.Role) != "" {
		user.Role = input.Role
	}

	updatedUser, err := u.userRepository.Update(user)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to update user: %v", err)
	}

	return updatedUser, nil
}

func (u *userUsecase) Delete(id string) error {
	_, err := u.FindById(id)
	if err != nil {
		return err
	}

	err = u.userRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete user : %v", err)
	}

	return nil
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{userRepository: userRepository}
}
