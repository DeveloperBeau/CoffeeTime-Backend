package api

//newUser Request struct
type postNewUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Token     string `json:"token"`
}

//newUser Response struct
type postNewUserResponse struct {
	Err error `json:"err,omitempty"`
}

func (r postNewUserResponse) error() error { return r.Err }
