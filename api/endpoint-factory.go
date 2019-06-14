package api

import (
	"context"

	"github.com/DeveloperBeau/CoffeeTime-Go/db"

	"github.com/go-kit/kit/endpoint"
)

const (
	started  = "started"
	finished = "finished"
	failed   = "failed"
)

// makePostNewUserEndpoint returns an endpoint via the New User service.
func makePostNewUserEndpoint(s Service, handler db.Handler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postNewUserRequest)
		e := s.postNewUser(ctx, handler, req)
		return postNewUserResponse{Err: e}, nil
	}
}

// makePostStartSessionEndpoint returns an endpoint via the New session service.
func makePostStartSessionEndpoint(s Service, handler db.Handler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postStartSessionRequest)
		sID, e := s.postStartSession(ctx, handler, req)
		if e != nil {
			return postStartSessionResponse{Err: e, Status: failed}, e
		}
		return postStartSessionResponse{SID: sID, Status: started}, nil
	}
}

// makePostEndSessionEndpoint returns an endpoint via the end session service.
func makePostEndSessionEndpoint(s Service, handler db.Handler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postEndSessionRequest)
		e := s.postEndSession(ctx, handler, req)
		if e != nil {
			return postEndSessionResponse{Err: e, Status: failed}, e
		}
		return postStartSessionResponse{Status: finished}, nil
	}
}

// makeGetSessionEndpoint returns an endpoint via the session service.
func makeGetSessionEndpoint(s Service, handler db.Handler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getSessionRequest)
		se, e := s.getSession(ctx, handler, req)
		return se, e
	}
}

// makePostOrderEndpoint returns an endpoint via the order service.
func makePostOrderEndpoint(s Service, handler db.Handler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postOrderRequest)
		res, e := s.postOrder(ctx, handler, req)
		return res, e
	}
}
