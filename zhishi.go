// 知识解构app入口
package main // import "github.com/leaderwolfpipi/zhishi"

import (
	"github.com/leaderwolfpipi/doris"
	"github.com/leaderwolfpipi/doris/middleware"
	"github.com/leaderwolfpipi/zhishi/helper"
	"github.com/leaderwolfpipi/zhishi/router"
)

func main() {
	// 实例化框架对象
	d := doris.New()

	// 展示banner
	d.ShowBanner = true

	// 关闭调试模式
	d.Debug = false

	// 全局中间件
	d.Use(middleware.Logger())
	d.Use(middleware.Recovery())

	// 设置token秘钥
	helper.SetSignKey(helper.GetTokenConfig().Token.SignKey)

	// 注册路由
	router.RegisterAuthRoutes(d)
	router.RegisterNoAuthRoutes(d)

	// 启动服务
	d.Run("localhost:9527")
}
