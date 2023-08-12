package handler

import (
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
