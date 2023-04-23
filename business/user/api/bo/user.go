package bo

type UserRegisterBo struct {
	Nick   string `json:"nick"`
	Mobile string `json:"mobile"`
	Code   string `json:"code"`
}

type UserLoginBo struct {
	Mobile string `json:"mobile"`
	Code   string `json:"code"`
}

type UserAuthBo struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}
