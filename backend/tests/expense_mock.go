package tests

import (
	"github.com/stretchr/testify/mock"
	"github.com/tf63/go_api/api/rest"
	"github.com/tf63/go_api/internal/entity"
)

type MockExpenseRepository struct {
	mock.Mock
}

func (mer *MockExpenseRepository) CreateExpense(input rest.NewExpense) (expenseId int, err error) {
	args := mer.Called(input)
	return args.Int(0), args.Error(1)
}

func (mer *MockExpenseRepository) ReadExpense(input rest.FindUser, expenseId int) (expense entity.Expense, err error) {
	args := mer.Called(input, expenseId)
	return args.Get(0).(entity.Expense), args.Error(1)
}

func (mer *MockExpenseRepository) ReadExpenses(input rest.FindUser) (expenses []entity.Expense, err error) {
	args := mer.Called(input)
	return args.Get(0).([]entity.Expense), args.Error(1)
}

func (mer *MockExpenseRepository) UpdateExpense(input rest.NewExpense, expenseId int) (err error) {
	args := mer.Called(input, expenseId)
	return args.Error(0)
}

func (mer *MockExpenseRepository) DeleteExpense(input rest.FindUser, expenseId int) (err error) {
	args := mer.Called(input, expenseId)
	return args.Error(0)
}
