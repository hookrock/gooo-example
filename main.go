package main

import (
	"fmt"

	"github.com/hookrock/gooo"
)

func main() {
	r := gooo.Default()
	//添加认证
	//r.Use(gooo.AuthMiddleware)
	//实现认证路由看守器
	//
	r.GET("/", func(c *gooo.Context) {
		c.View("index.tmpl", nil)
	})
	// 生成一个读取参数的get请求
	r.GET("/get", func(c *gooo.Context) {
		// name := c.Query("name") // 假设 Query 是获取查询参数的方法
		// if name == "" {
		// 	name = "Guest"
		// }
		name := c.GetParamWithDefault("name", "Guest")

		message := fmt.Sprintf("hello world %s", name)
		c.String(200, message)
	})

	r.Run()
}
