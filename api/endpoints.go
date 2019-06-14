package api

import (
	"github.com/DeveloperBeau/CoffeeTime-Go/db"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints : endpoint struct
type Endpoints struct {
	postNewUserEndpoint      endpoint.Endpoint
	postStartSessionEndpoint endpoint.Endpoint
	postEndSessionEndpoint   endpoint.Endpoint
	getSessionEndpoint       endpoint.Endpoint
	postOrderEndpoint        endpoint.Endpoint
}

// MakeServerEndpoints : Endpoint Factory Generator
func MakeServerEndpoints(s Service, handler db.Handler) Endpoints {
	return Endpoints{
		postNewUserEndpoint:      makePostNewUserEndpoint(s, handler),
		postStartSessionEndpoint: makePostStartSessionEndpoint(s, handler),
		postEndSessionEndpoint:   makePostEndSessionEndpoint(s, handler),
		getSessionEndpoint:       makeGetSessionEndpoint(s, handler),
		postOrderEndpoint:        makePostOrderEndpoint(s, handler),
	}
}
