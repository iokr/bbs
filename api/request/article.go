package request

type CreateArticleRequest struct {
	UserId      uint   `form:"user_id" json:"user_id"`
	Title       string `form:"title" json:"title"`
	Content     string `form:"content" json:"content"`
	IsTop       int32  `form:"is_top" json:"is_top"`
	IsPublished int32  `form:"is_published" json:"is_published"`
}

type UpdateArticleRequest struct {
	Id          uint   `form:"id" json:"id"`
	UserId      uint   `form:"user_id" json:"user_id"`
	Title       string `form:"title" json:"title"`
	Content     string `form:"content" json:"content"`
	IsTop       int32  `form:"is_top" json:"is_top"`
	IsPublished int32  `form:"is_published" json:"is_published"`
}
