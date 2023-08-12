package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/tf63/go_api/api/rest"
	"github.com/tf63/go_api/internal/entity"
	"github.com/tf63/go_api/internal/repository"
)

type serverHandler struct {
	er repository.ExpenseRepository
	ur repository.UserRepository
}

// Make sure we conform to ServerInterface
var _ rest.ServerInterface = (*serverHandler)(nil)

func NewServerHandler(er repository.ExpenseRepository) rest.ServerInterface {
	return &serverHandler{er}
}

func (sh *serverHandler) PostV1Expenses(w http.ResponseWriter, r *http.Request) {

	var newExpense rest.NewExpense
	err := json.NewDecoder(r.Body).Decode(&newExpense)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	log.Printf(`newExpense -> title: ` + *newExpense.Title + `, price: ` + strconv.Itoa(*newExpense.Price) + `, user_id: ` + strconv.Itoa(*newExpense.UserId))

	expense_id, err := sh.er.CreateExpense(newExpense)
	if err == entity.STATUS_SERVICE_UNAVAILABLE {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf(`Post 503 [Error]`)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		log.Printf(`Post 501 [Error]`)
		return
	}

	log.Printf(`Post 201 [OK]`)

	w.Header().Set("Location", r.Host+r.URL.Path+"/"+strconv.Itoa(expense_id))
	w.WriteHeader(http.StatusCreated)
}
