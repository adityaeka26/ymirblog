// Package rest is port handler.
package rest

import (
	"net/http"

	"github.com/abialemuel/ymirblog/pkg/entity"
	"github.com/abialemuel/ymirblog/pkg/usecase/article"
	"github.com/go-chi/chi/v5"

	"github.com/kubuskotak/asgard/rest"
)

// ArticleOption is a struct holding the handler options.
type ArticleOption func(Article *Article)

// Article handler instance data.
type Article struct {
	UcArticle article.T
}

// NewArticle creates a new Article handler instance.
//
//	var ArticleHandler = rest.NewArticle()
//
//	You can pass optional configuration options by passing a Config struct:
//
//	var adaptor = &adapters.Adapter{}
//	var ArticleHandler = rest.NewArticle(rest.WithArticleAdapter(adaptor))
func NewArticle(opts ...ArticleOption) *Article {
	// Create a new handler.
	var handler = &Article{}

	// Assign handler options.
	for o := range opts {
		var opt = opts[o]
		opt(handler)
	}

	// Return handler.
	return handler
}

func WithArticleUsecase(u article.T) ArticleOption {
	return func(a *Article) {
		a.UcArticle = u
	}
}

// Register is endpoint group for handler.
func (a *Article) Register(router chi.Router) {
	// PLEASE EDIT THIS EXAMPLE, how to register handler to router
	router.Get("/articles", rest.HandlerAdapter[GetArticleRequest](a.GetArticle).JSON)
}

// GetArticle endpoint func. /** PLEASE EDIT THIS EXAMPLE, return handler response */.
func (a *Article) GetArticle(w http.ResponseWriter, r *http.Request) (GetArticleResponse, error) {
	var (
		request GetArticleRequest
	)
	request, err := rest.GetBind[GetArticleRequest](r)
	if err != nil {
		return GetArticleResponse{}, rest.ErrBadRequest(w, r, err)
	}

	payload := entity.GetArticlePayload{
		Title:  request.Title,
		UserID: request.UserID,
		Limit:  request.Limit,
		Page:   request.Page,
	}
	articlePagination, err := a.UcArticle.GetAll(r.Context(), payload)
	if err != nil {
		return GetArticleResponse{}, err
	}
	// articlePagination := entity.ArticlesWithPagination{}

	rest.Paging(r, rest.Pagination{
		Page:  articlePagination.Metadata.Page,
		Limit: articlePagination.Metadata.Limit,
		Total: articlePagination.Metadata.Total,
	})

	result := GetArticleResponse{}
	for _, article := range articlePagination.Items {
		articleResponse := ArticleResponse{
			ID:    article.ID,
			Title: article.Title,
			Body:  article.Body,
		}
		if article.User != nil {
			articleResponse.User = &SimpleUserResponse{
				ID:   article.User.ID,
				Name: article.User.Name,
			}
		}
		for _, tag := range article.Tags {
			articleResponse.Tags = append(articleResponse.Tags, SimpleTagResponse{
				ID:   tag.ID,
				Name: tag.Name,
			})
		}
		result.Items = append(result.Items, &articleResponse)
	}

	return result, nil
}
