package api

import (
	"CoffeeTime-Go/db"
	"context"

	"github.com/go-kit/kit/endpoint"
)

// MakePostNewUserEndpoint returns an endpoint via the New User service.
func MakePostNewUserEndpoint(s Service, handler db.Handler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postNewUserRequest)
		e := s.PostNewUser(ctx, handler, req)
		return postNewUserResponse{Err: e}, nil
	}
}
