package api

import (
	"CoffeeTime-Go/db"
	"flag"
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
	startSessionPath = "/startSession/{id}"
	stopSessionPath  = "/stopSession"
	sessionPath      = "/session"
	orderPath        = "/order/{id}"
)

var s serviceHandler

// Run : Boots up server
func Run(endpoint string, db db.Handler) error {
	r := RunAPIOnRouter(db)
	return http.ListenAndServe(endpoint, r)
}

// RunAPIOnRouter : boots up router
func RunAPIOnRouter(db db.Handler) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(s, db)
	flag.Parse()

	r.Methods(postMethod).Path(addUserPath).Handler(httptransport.NewServer(
		e.PostNewUserEndpoint,
		decodePostNewUserRequest,
		encodeResponse,
	))

	r.Methods(postMethod).Path(startSessionPath).Handler(httptransport.NewServer(
		e.PostStartSessionEndpoint,
		decodePostStartSessionRequest,
		encodeResponse,
	))

	r.Methods(postMethod).Path(stopSessionPath).Handler(httptransport.NewServer(
		e.PostEndSessionEndpoint,
		decodePostEndSessionRequest,
		encodeResponse,
	))

	r.Methods(getMethod).Path(sessionPath).Handler(httptransport.NewServer(
		e.GetSessionEndpoint,
		decodeGetSessionRequest,
		encodeResponse,
	))

	r.Methods(postMethod).Path(orderPath).Handler(httptransport.NewServer(
		e.PostOrderEndpoint,
		decodeGetSessionRequest,
		encodeResponse,
	))

	return r
}
