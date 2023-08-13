package handler

import "github.com/tf63/go_api/internal/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Er repository.ExpenseRepository
	Ur repository.UserRepository
}
