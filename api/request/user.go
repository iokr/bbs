package request

type RegisterRequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginReply struct {
	Token string `json:"token"`
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
