package model

type User struct {
	Base
	Nick   string `json:"nick"`
	Mobile string `json:"mobile"`
}

func (User) TableName() string {
	return "user"
}
