// Package rest is port handler.
package rest

// GetUserRequest Get a User request.  /** PLEASE EDIT THIS EXAMPLE, request handler */.
type GetUserRequest struct {
}

// GetUserRequestID Get a User request by ID.  /** PLEASE EDIT THIS EXAMPLE, request handler */.
type GetUserRequestID struct {
	ID int `json:"id"`
}

// GetUserCSVRequest Get a User request.  /** PLEASE EDIT THIS EXAMPLE, request handler */.
type GetUserCSVRequest struct {
}

// CreateUserRequest create a User request.  /** PLEASE EDIT THIS EXAMPLE, request handler */.
type CreateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"email,required"`
}

// UpdateUserRequest update a User request.  /** PLEASE EDIT THIS EXAMPLE, request handler */.
type UpdateUserRequest struct {
	ID    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"email,required"`
}

// GetUserResponse Get a User response.  /** PLEASE EDIT THIS EXAMPLE, return handler response */.
type GetUserResponse struct {
	Message string         `json:"message"`
	User    *UserResponse  `json:"user,omitempty"`
	Users   []UserResponse `json:"users,omitempty"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"email,required"`
}

type GetAllUserRespone struct {
	Message string         `json:"message"`
	Items   []UserResponse `json:"items"`
}
