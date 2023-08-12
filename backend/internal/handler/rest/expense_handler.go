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

	log.Printf(`[POST] ` + r.URL.Path)

	var newExpense rest.NewExpense
	err := json.NewDecoder(r.Body).Decode(&newExpense)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf(`[POST] 503 Error: invalid request body`)
		return
	}

	log.Printf(`[POST] newExpense -> title: ` + *newExpense.Title + `, price: ` + strconv.Itoa(*newExpense.Price) + `, user_id: ` + strconv.Itoa(*newExpense.UserId))

	expense_id, err := sh.er.CreateExpense(newExpense)
	if err == entity.STATUS_SERVICE_UNAVAILABLE {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf(`[POST] 503 Error: database error`)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		log.Printf(`[POST] 501 Error: unexpected error`)
		return
	}

	log.Printf(`[POST] 201 OK`)

	w.Header().Set("Location", r.Host+r.URL.Path+"/"+strconv.Itoa(expense_id))
	w.WriteHeader(http.StatusCreated)
}

func (sh *serverHandler) GetV1Expenses(w http.ResponseWriter, r *http.Request) {

	log.Printf(`[GET] ` + r.URL.Path)

	var findUser rest.FindUser
	err := json.NewDecoder(r.Body).Decode(&findUser)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf(`[GET] 503 Error: invalid request body`)
		return
	}

	log.Printf(`[GET] findUser -> userId: ` + strconv.Itoa(*findUser.UserId))

	expenses, err := sh.er.ReadExpenses(findUser)
	if err == entity.STATUS_SERVICE_UNAVAILABLE {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf(`[GET] 503 Error: database error`)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		log.Printf(`[GET] 501 Error: unexpected error`)
		return
	}

	log.Printf(`[GET] 200 OK`)

	response := rest.Expenses{}
	for _, expense := range expenses {
		response = append(response, ExpenseDTO(&expense))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (sh *serverHandler) DeleteV1ExpensesExpenseId(w http.ResponseWriter, r *http.Request, expenseId int) {

	log.Printf(`[DELETE] ` + r.URL.Path)

	var findUser rest.FindUser
	err := json.NewDecoder(r.Body).Decode(&findUser)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf(`[DELETE] 503 Error: invalid request body`)
		return
	}

	log.Printf(`[DELETE] findUser -> userId: ` + strconv.Itoa(*findUser.UserId))

	err = sh.er.DeleteExpense(findUser, expenseId)
	if err == entity.STATUS_SERVICE_UNAVAILABLE {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf(`[DELETE] 503 Error: database error`)
		return
	} else if err == entity.STATUS_NOT_FOUND {
		w.WriteHeader(http.StatusNotFound)
		log.Printf(`[DELETE] 404 Error: not found`)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		log.Printf(`[DELETE] 501 Error: unexpected error`)
		return
	}

	log.Printf(`[DELETE] 204 OK`)

	w.WriteHeader(http.StatusNoContent)
}

func (sh *serverHandler) GetV1ExpensesExpenseId(w http.ResponseWriter, r *http.Request, expenseId int) {

	log.Printf(`[GET] ` + r.URL.Path)

	var findUser rest.FindUser
	err := json.NewDecoder(r.Body).Decode(&findUser)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf(`[GET] 503 Error: invalid request body`)
		return
	}

	log.Printf(`[GET] findUser -> userId: ` + strconv.Itoa(*findUser.UserId))

	expense, err := sh.er.ReadExpense(findUser, expenseId)
	if err == entity.STATUS_SERVICE_UNAVAILABLE {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf(`[GET] 503 Error: database error`)
		return
	} else if err == entity.STATUS_NOT_FOUND {
		w.WriteHeader(http.StatusNotFound)
		log.Printf(`[GET] 404 Error: not found`)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		log.Printf(`[GET] 501 Error: unexpected error`)
		return
	}

	log.Printf(`[GET] 200 OK`)

	response := ExpenseDTO(&expense)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (sh *serverHandler) PutV1ExpensesExpenseId(w http.ResponseWriter, r *http.Request, expenseId int) {

	log.Printf(`[PUT] ` + r.URL.Path)

	var newExpense rest.NewExpense
	err := json.NewDecoder(r.Body).Decode(&newExpense)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf(`[PUT] 503 Error: invalid request body`)
		return
	}

	log.Printf(`[PUT] newExpense -> title: ` + *newExpense.Title + `, price: ` + strconv.Itoa(*newExpense.Price) + `, user_id: ` + strconv.Itoa(*newExpense.UserId))

	err = sh.er.UpdateExpense(newExpense, expenseId)
	if err == entity.STATUS_SERVICE_UNAVAILABLE {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf(`[PUT] 503 Error: database error`)
		return
	} else if err == entity.STATUS_NOT_FOUND {
		w.WriteHeader(http.StatusNotFound)
		log.Printf(`[PUT] 404 Error: not found`)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		log.Printf(`[PUT] 501 Error: unexpected error`)
		return
	}

	log.Printf(`[PUT] 204 OK`)

	w.WriteHeader(http.StatusNoContent)
}
