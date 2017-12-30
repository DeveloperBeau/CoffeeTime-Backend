package api

import (
	"CoffeeTime-Go/db"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints : endpoint struct
type Endpoints struct {
	PostNewUserEndpoint endpoint.Endpoint
}

// MakeServerEndpoints : Endpoint Factory Generator
func MakeServerEndpoints(s Service, handler db.Handler) Endpoints {
	return Endpoints{
		PostNewUserEndpoint: MakePostNewUserEndpoint(s, handler),
	}
}
