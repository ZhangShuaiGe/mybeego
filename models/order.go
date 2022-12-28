package models

import "time"

type Order struct {
	// 对应表字段名为: id
	Id int
	// 对应表字段名为: shop_id , 下面字段名转换规则以此类推。
	ShopId int
	// struct字段名跟表字段名不一样，通过orm标签指定表字段名为customer_id
	Uid      int `orm:"column(customer_id)" json:"uid"`
	Nickname string
	Address  string
	// 数据库init_time字段是datetime类型，支持自动转换成Go的time.Time类型，但是数据库连接参数必须设置参数parseTime=true
	InitTime time.Time
}

// 指定Order结构体默认绑定的表名
func (o *Order) TableName() string {
	return "orders"
}
