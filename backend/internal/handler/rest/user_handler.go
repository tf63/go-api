package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/tf63/go_api/api/rest"
	"github.com/tf63/go_api/internal/entity"
)

func (sh *serverHandler) PostV1Users(w http.ResponseWriter, r *http.Request) {

	log.Printf(`[POST] ` + r.URL.Path)

	var newUser rest.NewUser
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	log.Printf(`[POST] newUser -> name: ` + *newUser.Name)

	user_id, err := sh.ur.CreateUser(newUser)
	if err == entity.STATUS_SERVICE_UNAVAILABLE {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf(`[POST] 503 Error`)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		log.Printf(`[POST] 501 Error`)
		return
	}

	log.Printf(`[POST] 201 OK`)

	w.Header().Set("Location", r.Host+r.URL.Path+"/"+strconv.Itoa(user_id))
	w.WriteHeader(http.StatusCreated)
}

func (sh *serverHandler) GetV1Users(w http.ResponseWriter, r *http.Request) {

	log.Printf(`[GET] ` + r.URL.Path)

	users, err := sh.ur.ReadUsers()
	if err == entity.STATUS_SERVICE_UNAVAILABLE {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf(`[GET] 503 Error`)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		log.Printf(`[GET] 501 Error`)
		return
	}

	log.Printf(`[GET] 200 OK`)

	response := rest.Users{}
	for _, user := range users {
		response = append(response, UserDTO(&user))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (sh *serverHandler) DeleteV1UsersUserId(w http.ResponseWriter, r *http.Request, userId int) {

	log.Printf(`[DELETE] ` + r.URL.Path)

	err := sh.ur.DeleteUser(userId)
	if err == entity.STATUS_SERVICE_UNAVAILABLE {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf(`[DELETE] 503 Error`)
		return
	} else if err == entity.STATUS_NOT_FOUND {
		w.WriteHeader(http.StatusNotFound)
		log.Printf(`[DELETE] 404 Error`)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		log.Printf(`[DELETE] 501 Error`)
		return
	}

	log.Printf(`[DELETE] 204 OK`)

	w.WriteHeader(http.StatusNoContent)
}

func (sh *serverHandler) GetV1UsersUserId(w http.ResponseWriter, r *http.Request, userId int) {

	log.Printf(`[GET] ` + r.URL.Path)

	user, err := sh.ur.ReadUser(userId)
	if err == entity.STATUS_SERVICE_UNAVAILABLE {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf(`[GET] 503 Error`)
		return
	} else if err == entity.STATUS_NOT_FOUND {
		w.WriteHeader(http.StatusNotFound)
		log.Printf(`[GET] 404 Error`)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		log.Printf(`[GET] 501 Error`)
		return
	}

	log.Printf(`[GET] 200 OK`)

	response := UserDTO(&user)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (sh *serverHandler) PutV1UsersUserId(w http.ResponseWriter, r *http.Request, userId int) {

	log.Printf(`[PUT] ` + r.URL.Path)

	var newUser rest.NewUser
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	log.Printf(`[PUT] newUser -> name: ` + *newUser.Name)

	err = sh.ur.UpdateUser(newUser, userId)
	if err == entity.STATUS_SERVICE_UNAVAILABLE {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf(`[PUT] 503 Error`)
		return
	} else if err == entity.STATUS_NOT_FOUND {
		w.WriteHeader(http.StatusNotFound)
		log.Printf(`[PUT] 404 Error`)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		log.Printf(`[PUT] 501 Error`)
		return
	}

	log.Printf(`[PUT] 204 OK`)

	w.WriteHeader(http.StatusNoContent)
}
