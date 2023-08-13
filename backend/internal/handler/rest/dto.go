package handler

import (
	"errors"

	"github.com/tf63/go_api/api/rest"
	"github.com/tf63/go_api/internal/entity"
)

func ExpenseDTO(e *entity.Expense) rest.Expense {
	id := int(e.ID)
	price := int(e.Price)
	title := e.Title
	userId := int(e.UserID)

	return rest.Expense{
		Id:     &id,
		Price:  &price,
		Title:  &title,
		UserId: &userId,
	}
}

func UserDTO(e *entity.User) rest.User {
	id := int(e.ID)
	name := e.Name

	return rest.User{
		Id:   &id,
		Name: &name,
	}
}

func NewExpenseDTO(s *rest.NewExpense) (entity.NewExpense, error) {

	var price *uint
	if s.Price != nil {
		priceValue := uint(*s.Price)
		price = &priceValue
	} else {
		price = nil
	}

	if s.UserId == nil {
		return entity.NewExpense{}, errors.New("invalid userId")
	}

	return entity.NewExpense{
		Title:  s.Title,
		Price:  price,
		UserID: uint(*s.UserId),
	}, nil
}

func FindUserDTO(s *rest.FindUser) entity.FindUser {
	return entity.FindUser{
		ID: uint(*s.UserId),
	}
}

func NewUserDTO(s *rest.NewUser) entity.NewUser {
	return entity.NewUser{
		Name: s.Name,
	}
}
