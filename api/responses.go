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

//newSession Request struct
type postStartSessionRequest struct {
	ID string `json:"id"`
}

//newSession Response struct
type postStartSessionResponse struct {
	ID       string `json:"sessionID,omitempty"`
	Err      error  `json:"err,omitempty"`
	IsActive bool   `json:"isActive,omitempty"`
}

func (r postStartSessionResponse) error() error { return r.Err }

//endSession Request struct
type postEndSessionRequest struct {
	ID string `json:"id"`
}

//endSession Response struct
type postEndSessionResponse struct {
	ID       string `json:"sessionID,omitempty"`
	Err      error  `json:"err,omitempty"`
	IsActive bool   `json:"isActive,omitempty"`
}

func (r postEndSessionResponse) error() error { return r.Err }

//session Request struct
type getSessionRequest struct {
	ID string `json:"id"`
}

//session Response struct, users who respond to a new session will receive this
type getUserSessionResponse struct {
	ID    string `json:"sessionId,omitempty"`
	Err   error  `json:"err,omitempty"`
	Order order  `json:"yourOrder,omitempty"`
}

//session Response for whole group struct, only the person who started the session will receive this
type getGroupSessionResponse struct {
	ID     string  `json:"sessionId,omitempty"`
	Err    error   `json:"err,omitempty"`
	Orders []order `json:"orders,omitempty"`
}

func (r getUserSessionResponse) error() error  { return r.Err }
func (r getGroupSessionResponse) error() error { return r.Err }

//PostOrderRequest Request struct
type postOrderRequest struct {
	SessionID string `json:"SessionID"`
	Order     order  `json:"order"`
}

//PostOrderResponse Response struct
type postOrderResponse struct {
	ID           string `json:"id,omitempty"`
	Err          error  `json:"err,omitempty"`
	Result       string `json:"response,omitempty"`
	IsSuccessful bool   `json:"isSuccessful,omitempty"`
}

func (r postOrderResponse) error() error { return r.Err }

//Order generic struct
type order struct {
	UserID  string `json:"userID"`
	Request string `json:"request"`
}
