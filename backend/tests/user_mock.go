package tests

import (
	"github.com/stretchr/testify/mock"
	"github.com/tf63/go_api/internal/entity"
)

type MockUserRepository struct {
	mock.Mock
}

func (mur *MockUserRepository) CreateUser(input entity.NewUser) (userId int, err error) {
	args := mur.Called(input)
	return args.Int(0), args.Error(1)
}

func (mur *MockUserRepository) ReadUser(userId int) (user entity.User, err error) {
	args := mur.Called(userId)
	return args.Get(0).(entity.User), args.Error(1)
}

func (mur *MockUserRepository) ReadUsers() (users []entity.User, err error) {
	args := mur.Called()
	return args.Get(0).([]entity.User), args.Error(1)
}

func (mur *MockUserRepository) UpdateUser(input entity.NewUser, userId int) (err error) {
	args := mur.Called(input, userId)
	return args.Error(0)
}

func (mur *MockUserRepository) DeleteUser(userId int) (err error) {
	args := mur.Called(userId)
	return args.Error(0)
}
