package handler

import (
	"github.com/tf63/go_api/api/rest"
	"github.com/tf63/go_api/internal/repository"
)

type serverHandler struct {
	er repository.ExpenseRepository
	ur repository.UserRepository
}

// Make sure we conform to ServerInterface
var _ rest.ServerInterface = (*serverHandler)(nil)

func NewServerHandler(er repository.ExpenseRepository, ur repository.UserRepository) rest.ServerInterface {
	return &serverHandler{er, ur}
}
