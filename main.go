package main

import (
	_ "mybeego/models"
	_ "mybeego/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	//an official log.Logger
	beego.Run()
}
