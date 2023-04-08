package biz

import (
	"context"

	"github.com/iokr/bbs/api/request"
	"github.com/iokr/bbs/internal/data"
	"github.com/iokr/bbs/internal/errors"
)

type ArticleBiz struct {
	articleRepo *data.ArticleRepo
}

func NewArticleBiz(articleRepo *data.ArticleRepo) *ArticleBiz {
	return &ArticleBiz{articleRepo: articleRepo}
}

func (b *ArticleBiz) CreateArticle(ctx context.Context, articleReq *request.CreateArticleRequest) error {
	if articleReq.UserId <= 0 {
		return errors.UserNotFound
	}

	if articleReq.Title == "" || articleReq.Content == "" {
		return errors.TitleOrContentIsEmpty
	}

	var abstract string
	contentLength := len(articleReq.Content)
	if contentLength <= 300 {
		abstract = articleReq.Content
	} else {
		abstract = articleReq.Content[:300]
	}

	return b.articleRepo.Create(ctx, &data.Article{
		UserId:      articleReq.UserId,
		Title:       articleReq.Title,
		Abstract:    abstract,
		Content:     articleReq.Content,
		IsTop:       articleReq.IsTop,
		IsPublished: articleReq.IsPublished,
	})
}

func (b *ArticleBiz) UpdateArticle(ctx context.Context, articleReq *request.UpdateArticleRequest) error {
	if articleReq.Id <= 0 {
		return errors.ArticleNotFound
	}
	if articleReq.UserId <= 0 {
		return errors.UserNotFound
	}

	if articleReq.Title == "" || articleReq.Content == "" {
		return errors.TitleOrContentIsEmpty
	}

	articleData, err := b.articleRepo.GetByCond(ctx, map[string]interface{}{"id": articleReq.Id})
	if err != nil {
		return err
	}

	if articleData.UserId != articleReq.UserId {
		return errors.ArticleNotFound
	}

	var abstract string
	contentLength := len(articleReq.Content)
	if contentLength <= 300 {
		abstract = articleReq.Content
	} else {
		abstract = articleReq.Content[:300]
	}

	return b.articleRepo.UpdateByIdAndCond(ctx, articleData.ID, map[string]interface{}{
		"title":        articleReq.Title,
		"abstract":     abstract,
		"content":      articleReq.Content,
		"is_top":       articleReq.IsTop,
		"is_published": articleReq.IsPublished,
	})
}
