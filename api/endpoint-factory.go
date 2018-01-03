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

// MakePostStartSessionEndpoint returns an endpoint via the New session service.
func MakePostStartSessionEndpoint(s Service, handler db.Handler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postStartSessionRequest)
		e := s.PostStartSession(ctx, handler, req)
		return nil, nil
	}
}

// MakePostEndSessionEndpoint returns an endpoint via the end session service.
func MakePostEndSessionEndpoint(s Service, handler db.Handler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postEndSessionRequest)
		e := s.PostEndSession(ctx, handler, req)
		return nil, nil
	}
}

// MakeGetSessionEndpoint returns an endpoint via the session service.
func MakeGetSessionEndpoint(s Service, handler db.Handler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getSessionRequest)
		e := s.GetSession(ctx, handler, req)
		return nil, nil
	}
}

// MakePostOrderEndpoint returns an endpoint via the order service.
func MakePostOrderEndpoint(s Service, handler db.Handler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postOrderRequest)
		e := s.PostOrder(ctx, handler, req)
		return nil, nil
	}
}
