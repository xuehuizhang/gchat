package vo

type UserLoginVo struct {
	Nick string `json:"nick"`
	Toke string `json:"toke"`
}

type UserRegisterVo struct {
	Nick   string `json:"nick"`
	Mobile string `json:"mobile"`
}

type UserAuthVo struct {
	UserId int64 `json:"user_id"`
}
