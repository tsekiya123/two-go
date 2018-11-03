package main

import (
	_ "hello/routers"
	"github.com/astaxie/beego"
)

var ConstString = "HOGEHOGEHOGEHOGEHOGEHOGE"

func main() {
	beego.Run()
}

