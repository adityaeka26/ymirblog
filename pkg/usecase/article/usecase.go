package article

import (
	"context"

	"github.com/abialemuel/ymirblog/pkg/entity"
	"github.com/abialemuel/ymirblog/pkg/persist/ymirblog/ent/article"
	"github.com/abialemuel/ymirblog/pkg/persist/ymirblog/ent/user"
	pkgRest "github.com/kubuskotak/asgard/rest"
	"go.opentelemetry.io/otel/trace"
)

// GetAll returns resource articles.
func (i *impl) GetAll(ctx context.Context, request entity.GetArticlePayload) (entity.ArticlesWithPagination, error) {
	span := trace.SpanFromContext(ctx)
	defer span.End()

	client := i.adapter.YmirblogPersist
	query := client.Article.
		Query().
		WithUser().
		WithTags()

	if request.Title != nil {
		query = query.Where(article.TitleContains(*request.Title))
	}

	if request.UserID != nil {
		query = query.Where(article.HasUserWith(user.IDEQ(*request.UserID)))
	}

	// pagination
	total, err := query.Count(ctx)
	if err != nil {
		return entity.ArticlesWithPagination{}, err
	}
	metadata := pkgRest.Pagination{
		Page:  request.Page,
		Limit: request.Limit,
		Total: total,
	}

	offset := (request.Page - 1) * request.Limit
	items, err := query.
		Limit(request.Limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		return entity.ArticlesWithPagination{}, err
	}

	res := entity.ArticlesWithPagination{}
	for _, a := range items {
		entityArticle := entity.Article{
			ID:    a.ID,
			Title: a.Title,
			Body:  a.Body,
		}
		if a.Edges.User != nil {
			entityArticle.User = &entity.User{
				ID:    a.Edges.User.ID,
				Name:  a.Edges.User.Name,
				Email: a.Edges.User.Email,
			}
		}

		for _, tag := range a.Edges.Tags {
			entityArticle.Tags = append(entityArticle.Tags, entity.Tag{
				ID:   tag.ID,
				Name: tag.Name,
			})
		}
		res.Items = append(res.Items, &entityArticle)
	}
	res.Metadata = metadata

	return res, nil
}
