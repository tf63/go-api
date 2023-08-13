package handler

import (
	"strconv"

	"github.com/tf63/go_api/api/graph"
	"github.com/tf63/go_api/internal/entity"
)

func ExpenseDTO(e *entity.Expense) graph.Expense {
	id := strconv.Itoa(int(e.ID))
	price := e.Price
	title := e.Title
	userId := e.UserID

	return graph.Expense{
		ID:        id,
		Price:     price,
		Title:     title,
		UserID:    userId,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func UserDTO(e *entity.User) graph.User {
	id := strconv.Itoa(int(e.ID))
	name := e.Name

	return graph.User{
		ID:        id,
		Name:      name,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func NewExpenseDTO(s *graph.NewExpense) entity.NewExpense {
	return entity.NewExpense{
		Title:  s.Title,
		Price:  s.Price,
		UserID: s.UserID,
	}
}

func FindUserDTO(s *graph.FindUser) entity.FindUser {
	return entity.FindUser{
		ID: s.UserID,
	}
}

func NewUserDTO(s *graph.NewUser) entity.NewUser {
	return entity.NewUser{
		Name: s.Name,
	}
}
