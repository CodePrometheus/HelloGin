package model

//会员数据结构体定义
type Member struct {
	Id           int64   `xorm:"pk autoincr" json:"id"`
	UserName     string  `xorm:"varchar(20)" json:"user_name"`
	Mobile       string  `xorm:"varchar(11)" json:"mobile"`
	Password     string  `xorm:"varchar(255)" json:"password"`
	RegisterTime int64   `xorm:"bigint" json:"register_time"`
	Avatar       string  `xorm:"varchar(255)" json:"avatar"`
	Balance      float64 `xorm:"double" json:"balance"`
	IsActive     int8    `xorm:"tinyint" json:"is_active"`
	City         string  `xorm:"varchar(10)" json:"city"`
}
