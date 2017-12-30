package api

import (
	"CoffeeTime-Go/db"
	"CoffeeTime-Go/db/model"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

// Service : Service functions for each endpoint
type Service interface {
	PostNewUser(ctx context.Context, handler db.Handler, user postNewUserRequest) error
}

var (
	// ErrCorruptData : Corrupt data entry err
	ErrCorruptData = errors.New("Data error - Please verify your details are correct")
	// ErrAlreadyExists : already exists error
	ErrAlreadyExists = errors.New("Input error - Already Exists")
	// ErrNotFound : Not found error
	ErrNotFound = errors.New("Input error - not found")
)

type e interface {
	error() error
}

type serviceHandler struct{}

func (sh serviceHandler) PostNewUser(ctx context.Context, handler db.Handler, user postNewUserRequest) error {
	newUser := model.User{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Token: user.Token}
	err := handler.AddUser(newUser)
	if err != nil {
		log.Println(err)
	}
	return err
}

//Decoders

func decodePostNewUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	decoder := json.NewDecoder(r.Body)
	var u postNewUserRequest
	err = decoder.Decode(&u)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}
	return u, nil
}

//Encoders

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(e); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrAlreadyExists, ErrCorruptData, db.ErrUserExists:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
