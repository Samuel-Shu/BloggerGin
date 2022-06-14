package main

import (
	"ginblog/Models"
	"ginblog/Routers"
)

func main() {
	// 引用数据库
	Models.InitDb()
    // 引入路由组件
	Routers.InitRouter()

}
