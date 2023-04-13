package biz

import (
	"context"

	"github.com/iokr/bbs/api/request"
	"github.com/iokr/bbs/internal/data"
	"github.com/samber/lo"
)

type IndexBiz struct {
	userRepo    *data.UserRepo
	articleRepo *data.ArticleRepo
}

func NewIndexBiz(userRepo *data.UserRepo, articleRepo *data.ArticleRepo) *IndexBiz {
	return &IndexBiz{
		userRepo:    userRepo,
		articleRepo: articleRepo,
	}
}

func (b *IndexBiz) FindMoreArticles(ctx context.Context, param *request.MoreArticleRequest) (
	[]*request.MoreArticleReply, error) {
	articles, err := b.articleRepo.FindOrderByIdAndLimit(ctx, param.CurrentArticleId, param.Limit)
	if err != nil {
		return nil, err
	}

	userIdMap := make(map[uint]struct{})
	lo.ForEach(articles, func(item *data.Article, idx int) {
		userIdMap[item.UserId] = struct{}{}
	})

	var userIds []uint
	for userId := range userIdMap {
		userIds = append(userIds, userId)
	}

	userMap, err := b.userRepo.FindMapByIds(ctx, userIds)
	if err != nil {
		return nil, err
	}

	return lo.Map(articles, func(item *data.Article, index int) *request.MoreArticleReply {
		result := &request.MoreArticleReply{
			UserId: item.UserId,
			Article: &request.MoreArticleResult{
				Id:           item.ID,
				Title:        item.Title,
				Abstract:     item.Abstract,
				CreatedTime:  item.CreatedAt.Format("2006-01-02 15:04:05"),
				LikeCount:    12,
				CommentCount: 12345,
				ViewCount:    1234567,
			},
		}

		if user, ok := userMap[item.UserId]; ok {
			result.UserName = user.UserName
			result.AvatarUrl = user.AvatarUrl
		}
		return result
	}), nil
}
