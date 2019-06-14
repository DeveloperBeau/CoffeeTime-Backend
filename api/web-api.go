package api

import (
	"github.com/DeveloperBeau/CoffeeTime-Go/db"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

const (
	postMethod = "POST"
	getMethod  = "GET"
)

const (
	addUserPath      = "/newUser"
	startSessionPath = "/startSession/{UID}"
	stopSessionPath  = "/stopSession"
	sessionPath      = "/session"
	orderPath        = "/order/{UID}"
)

var s serviceHandler

// Run : Boots up server
func Run(endpoint string, db db.Handler) error {
	r := RunAPIOnRouter(db)
	return http.ListenAndServe(endpoint, r)
}

// RunAPIOnRouter : boots up router
func RunAPIOnRouter(db db.Handler) http.Handler {
	r := router(db)

	return r
}

func router(db db.Handler) (r *mux.Router) {
	r = mux.NewRouter()
	e := MakeServerEndpoints(s, db)
	r.Methods(postMethod).Path(addUserPath).Handler(httptransport.NewServer(
		e.postNewUserEndpoint,
		decodePostNewUserRequest,
		encodeResponse,
	))

	r.Methods(postMethod).Path(startSessionPath).Handler(httptransport.NewServer(
		e.postStartSessionEndpoint,
		decodePostStartSessionRequest,
		encodeResponse,
	))

	r.Methods(postMethod).Path(stopSessionPath).Handler(httptransport.NewServer(
		e.postEndSessionEndpoint,
		decodePostEndSessionRequest,
		encodeResponse,
	))

	r.Methods(getMethod).Path(sessionPath).Handler(httptransport.NewServer(
		e.getSessionEndpoint,
		decodeGetSessionRequest,
		encodeResponse,
	))

	r.Methods(postMethod).Path(orderPath).Handler(httptransport.NewServer(
		e.postOrderEndpoint,
		decodePostOrderRequest,
		encodeResponse,
	))
	return
}
