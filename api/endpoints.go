package api

import (
	"CoffeeTime-Go/db"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints : endpoint struct
type Endpoints struct {
	PostNewUserEndpoint      endpoint.Endpoint
	PostStartSessionEndpoint endpoint.Endpoint
	PostEndSessionEndpoint   endpoint.Endpoint
	GetSessionEndpoint       endpoint.Endpoint
	PostOrderEndpoint        endpoint.Endpoint
}

// MakeServerEndpoints : Endpoint Factory Generator
func MakeServerEndpoints(s Service, handler db.Handler) Endpoints {
	return Endpoints{
		PostNewUserEndpoint:      MakePostNewUserEndpoint(s, handler),
		PostStartSessionEndpoint: MakePostStartSessionEndpoint(s, handler),
		PostEndSessionEndpoint:   MakePostEndSessionEndpoint(s, handler),
		GetSessionEndpoint:       MakeGetSessionEndpoint(s, handler),
		PostOrderEndpoint:        MakePostOrderEndpoint(s, handler),
	}
}
