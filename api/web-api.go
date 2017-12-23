package api

import (
	"CoffeeTime-Go/db"
	"flag"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var s serviceHandler

func Run(endpoint string, db db.Handler) error {
	r := RunAPIOnRouter(db)
	return http.ListenAndServe(endpoint, r)
}

func RunAPIOnRouter(db db.Handler) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(s, db)
	flag.Parse()

	r.Methods("POST").Path("/addUser").Handler(httptransport.NewServer(
		e.PostNewUserEndpoint,
		decodePostNewUserRequest,
		encodeResponse,
	))

	return r
}
