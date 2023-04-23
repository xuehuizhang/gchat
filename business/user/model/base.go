package model

type Base struct {
	Id         int64 `json:"id"`
	Status     int   `json:"status"`
	CreateTime int64 `json:"create_time" gorm:"autoCreateTime:milli"`
	UpdateTime int64 `json:"update_time" gorm:"autoUpdateTime:milli"`
	DeleteTime int64 `json:"delete_time" gorm:"autoDeleteTime:milli"`
}
