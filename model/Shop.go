package model

// Shop 商家结构体/商户
type Shop struct {
	//id
	Id int64 `xorm:"pk autoincr" json:"id"`
	//商铺名称
	Name string `xorm:"varchar(12)" json:"name"`
	//宣传信息
	PromotionInfo string `xorm:"varchar(30)" json:"promotion_info"`
	//地址
	Address string `xorm:"varchar(100)" json:"address"`
	//联系电话
	Phone string `xorm:"varchar(11)" json:"phone"`
	//店铺营业状态
	Status int `xorm:"tinyint" json:"status"`

	//经度
	Longitude float64 `xorm:"" json:"longitude"`
	//纬度
	Latitude float64 `xorm:"" json:"latitude"`
	//店铺图标
	ImagePath string `xorm:"varchar(255)" json:"image_path"`

	IsNew     bool `xorm:"bool" json:"is_new"`
	IsPremium bool `xorm:"bool" json:"is_premium"`

	//商铺评分
	Rating float32 `xorm:"float" json:"rating"`
	//评分总数
	RatingCount int64 `xorm:"int" json:"rating_count"`
	//当前订单总数
	RecentOrderNum int64 `xorm:"int" json:"recent_order_num"`

	//配送起送价
	MinimumOrderAmount int32 `xorm:"int" json:"minimum_order_amount"`
	//配送费
	DeliveryFee int32 `xorm:"int" json:"delivery_fee"`

	//营业时间
	OpeningHours string `xorm:"varchar(20)" json:"opening_hours"`
}
