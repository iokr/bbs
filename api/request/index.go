package request

type MoreArticleRequest struct {
	Limit            int  `form:"limit" json:"limit"`
	UserId           uint `form:"user_id" json:"user_id"`
	CurrentArticleId uint `form:"current_article_id" json:"current_article_id"`
}

type MoreArticleReply struct {
	UserId    uint               `json:"user_id"`
	UserName  string             `json:"user_name"`
	AvatarUrl string             `json:"avatar_url"`
	Article   *MoreArticleResult `json:"article"`
}

type MoreArticleResult struct {
	Id           uint     `json:"id"`
	Title        string   `json:"title"`
	Abstract     string   `json:"abstract"`
	CreatedTime  string   `json:"created_time"`
	LikeCount    int      `json:"like_count"`
	CommentCount int      `json:"comment_count"`
	ViewCount    int      `json:"view_count"`
	Tags         []string `json:"tags"`
}
