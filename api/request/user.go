package request

type RegisterRequest struct {
	UserName string `form:"user_name" json:"user_name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type LoginRequest struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type LoginReply struct {
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
}

// ------------------------------ type -----------------------------------

type UserResult struct {
	Id        uint   `json:"id"`
	UserName  string `json:"user_name"`
	NickName  string `json:"nick_name"`
	Uid       string `json:"uid"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	AvatarUrl string `json:"avatar_url"`
}
