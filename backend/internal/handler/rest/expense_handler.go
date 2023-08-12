package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/tf63/go_api/api/rest"
	"github.com/tf63/go_api/internal/entity"
)

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

func (sh *serverHandler) GetV1Expenses(w http.ResponseWriter, r *http.Request) {

	var findUser rest.FindUser
	err := json.NewDecoder(r.Body).Decode(&findUser)

	expenses, err := sh.er.ReadExpenses(findUser)
	if err == entity.STATUS_SERVICE_UNAVAILABLE {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf(`Post 503 [Error]`)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		log.Printf(`Post 501 [Error]`)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expenses)
}

func (sh *serverHandler) DeleteV1ExpensesExpenseId(w http.ResponseWriter, r *http.Request, expenseId rest.ExpenseId) {

}

func (sh *serverHandler) GetV1ExpensesExpenseId(w http.ResponseWriter, r *http.Request, expenseId rest.ExpenseId) {

}

func (sh *serverHandler) PutV1ExpensesExpenseId(w http.ResponseWriter, r *http.Request, expenseId rest.ExpenseId) {

}
