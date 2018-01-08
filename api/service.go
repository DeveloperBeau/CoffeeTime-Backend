package api

import (
	"CoffeeTime-Go/db"
	"CoffeeTime-Go/db/model"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

// Service : Service functions for each endpoint
type Service interface {
	postNewUser(ctx context.Context, handler db.Handler, user postNewUserRequest) error
	postStartSession(ctx context.Context, handler db.Handler, session postStartSessionRequest) (string, error)
	postEndSession(ctx context.Context, handler db.Handler, session postEndSessionRequest) error
	getSession(ctx context.Context, handler db.Handler, session getSessionRequest) (interface{}, error)
	postOrder(ctx context.Context, handler db.Handler, session postOrderRequest) (*postOrderResponse, error)
}

var (
	// ErrCorruptData : Corrupt data entry err
	ErrCorruptData = errors.New("Data error - Please verify your details are correct")
	// ErrAlreadyExists : already exists error
	ErrAlreadyExists = errors.New("Input error - Already Exists")
	// ErrNotFound : Not found error
	ErrNotFound = errors.New("Input error - Not found")
	// ErrSessionNotFound : No Session is active
	ErrSessionNotFound = errors.New("Session error - No session is active")
    // ErrSomethingWentWrong : Generic issue error
    ErrSomethingWentWrong = errors.New("Error - Something went wrong, please try again later")
)

type e interface {
	error() error
}

type serviceHandler struct{}

func (sh serviceHandler) postNewUser(ctx context.Context, handler db.Handler, user postNewUserRequest) error {
	newUser := model.User{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Token: user.Token}
	err := handler.AddUser(newUser)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (sh serviceHandler) postStartSession(ctx context.Context, handler db.Handler, sr postStartSessionRequest) (string, error) {
	ns := model.Session{UserID: sr.UID, IsActive: true, Started: time.Now(), Orders: nil}
	// TODO: Need to setup sending push notifications to devices.
	res, err := handler.StartSession(ns)
	return res, err
}

func (sh serviceHandler) postEndSession(ctx context.Context, handler db.Handler, sr postEndSessionRequest) error {
	err := handler.EndSession(sr.UID)
	return err
}

func (sh serviceHandler) getSession(ctx context.Context, handler db.Handler, sr getSessionRequest) (interface{}, error) {
	s := handler.Session(sr.UID)
	if sr.UID == s.UserID {
		gs := getGroupSessionResponse{SID: s.ID, Orders: s.Orders}
		return gs, nil
	} else if len(s.Orders) == 1 {
		us := getUserSessionResponse{SID: s.ID, Order: s.Orders[0]}
		return us, nil
	}
	return nil, ErrSessionNotFound
}

func (sh serviceHandler) postOrder(ctx context.Context, handler db.Handler, o postOrderRequest) (*postOrderResponse, error) {

	order := o.Order
	oID, e := handler.Order(order)
	if oID != "" {
		por := postOrderResponse{OID:oID}
		return &por, e
	}
	return nil, e
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

func decodePostStartSessionRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	decoder := json.NewDecoder(r.Body)
	var s postStartSessionRequest
	err = decoder.Decode(&s)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func decodePostEndSessionRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	decoder := json.NewDecoder(r.Body)
	var s postEndSessionRequest
	err = decoder.Decode(&s)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func decodeGetSessionRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	decoder := json.NewDecoder(r.Body)
	var s getSessionRequest
	err = decoder.Decode(&s)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func decodePostOrderRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	decoder := json.NewDecoder(r.Body)
	var s postOrderRequest
	err = decoder.Decode(&s)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}
	return s, nil
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
