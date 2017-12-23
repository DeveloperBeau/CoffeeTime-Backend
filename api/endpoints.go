package api

import (
	"CoffeeTime-Go/db"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	PostNewUserEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service, handler db.Handler) Endpoints {
	return Endpoints{
		PostNewUserEndpoint: MakePostNewUserEndpoint(s, handler),
	}
}
