package main

import (
	_ "bootcamp-go/weatherAPI/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

