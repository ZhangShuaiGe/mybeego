package controllers

import (
	"fmt"
	"mybeego/models"
	"time"

	"github.com/beego/beego/v2/adapter/orm"
	web "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	web.Controller
}

type JSONStruct struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Login struct {
	Username string `orm:"column(name)"`
	Password string `orm:"column(password)"`
}

func (c *LoginController) OrderList() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(3333, err)
		}
	}()

	o := orm.NewOrm()
	order_list := &models.Order{}
	err := c.BindJSON(order_list)
	if err != nil {
		fmt.Println(4444, err)
	}

	err2 := o.Read(order_list)
	if err2 != nil {
		fmt.Print(66666, err2)
		panic(err2)
	}
	c.Data["json"] = order_list
	c.ServeJSON()
}

func (c *LoginController) CreatedOrder() {
	o := orm.NewOrm()

	// 创建一个新的订单
	order := &models.Order{}
	order.InitTime = time.Now()
	err := c.BindJSON(order)
	if err != nil {
		fmt.Print(5555555, err)
		c.Ctx.WriteString(err.Error())
		return
	}

	// 调用orm的Insert函数插入数据
	// 等价sql： INSERT INTO `orders` (`shop_id`, `customer_id`, `nickname`,`address`, `init_time`) VALUES (1, 1002, '大锤', '深圳南山区', '2019-06-24 23:08:57')
	id, err := o.Insert(order)

	if err != nil {
		fmt.Println("插入失败")
	} else {
		// 插入成功会返回插入数据自增字段，生成的id
		fmt.Println("新插入数据的id为:", id)
	}
	c.Data["json"] = order
	c.ServeJSON()
}
