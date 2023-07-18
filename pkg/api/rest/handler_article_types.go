// Package rest is port handler.
package rest

// GetArticleRequest Get a Article request.  /** PLEASE EDIT THIS EXAMPLE, request handler */.
type GetArticleRequest struct {
	Title  *string
	UserID *int
	Limit  int `validate:"gte=0,default=10"`
	Page   int `validate:"gte=0,default=1"`
}

// GetArticleResponse Get a Article response.  /** PLEASE EDIT THIS EXAMPLE, return handler response */.
type GetArticleResponse struct {
	Items []*ArticleResponse `json:"items"`
}

type ArticleResponse struct {
	ID    int                 `json:"id"`
	Title string              `json:"title"`
	Body  string              `json:"body"`
	User  *SimpleUserResponse `json:"user,omitempty"`
	Tags  []SimpleTagResponse `json:"tags"`
}

type SimpleUserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SimpleTagResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
