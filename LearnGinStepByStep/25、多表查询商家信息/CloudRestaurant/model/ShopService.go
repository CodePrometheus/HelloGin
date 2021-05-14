package model

//商铺服务数据对应表
type ShopService struct {
	ShopId    int64 `xorm:"pk not null" json:"shop_id"`
	ServiceId int64 `xorm:"pk not null" json:"service_id"`
}
