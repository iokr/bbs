package errors

const (
	Success               Error = 0
	Default               Error = 10000000
	UserNotFound          Error = 10000001
	EmailIsExist          Error = 10000002
	NickNameIsEmpty       Error = 10000003
	EmailIsEmpty          Error = 10000004
	PasswordIsEmpty       Error = 10000005
	PasswordIsError       Error = 10000006
	EmailIsLocked         Error = 10000007
	TitleOrContentIsEmpty Error = 10000008
	CategoryNotFound      Error = 10000009
	CategoryIsEmpty       Error = 10000010
	CategoryIsExist       Error = 10000011
)

var errorMap = InitErrorMap()

func InitErrorMap() map[Error]string {
	info := make(map[Error]string)

	info[Success] = "success"
	info[Default] = "system error"
	info[UserNotFound] = "用户不存在"
	info[EmailIsExist] = "邮箱已存在"
	info[NickNameIsEmpty] = "用户昵称不能为空"
	info[EmailIsEmpty] = "邮箱不能为空"
	info[PasswordIsEmpty] = "密码不能为空"
	info[PasswordIsError] = "用户名密码不正确"
	info[EmailIsLocked] = "该邮箱已经锁定,请联系管理员解锁"
	info[TitleOrContentIsEmpty] = "标题或内容不能为空"
	info[CategoryNotFound] = "分类不存在"
	info[CategoryIsEmpty] = "分类不能为空"
	info[CategoryIsExist] = "分类已存在"
	return info
}
