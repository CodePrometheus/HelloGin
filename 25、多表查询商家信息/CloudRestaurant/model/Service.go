package model

//商铺服务基础表结构体定义
type Service struct {
	Id int64 `xorm:"pk autoincr" json:"id"`
	Name string `xorm:"varchar(20)" json:"name"`
	Description string  `xorm:"varchar(30)" json:"description"`
	IconName string 	`xorm:"varchar(3)" json:"icon_name"`
	IconColor string 	`xorm:"varchar(6)" json:"icon_color"`
}

