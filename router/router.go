package router

import (
	"github.com/leaderwolfpipi/doris"
	"github.com/leaderwolfpipi/doris/middleware"
	"github.com/leaderwolfpipi/zhishi/helper"
	"github.com/leaderwolfpipi/zhishi/transport/restful"
)

// Register Routers
// 鉴权路由
func RegisterAuthRoutes(d *doris.Doris) {
	// api路由组
	api := d.Group("/api/v1")

	// JWT验证
	// fmt.Println(helper.GetSignKey())
	api.Use(middleware.JWT(helper.GetSignKey()))

	// 添加文章
	api.POST("/article/add", restful.ArticleAdd)

	// 编辑文章
	api.POST("/article/modify", restful.ArticleModify)

	// 删除文章
	api.POST("/article/del", restful.ArticleDel)

	// 点赞文章
	api.POST("/article/like", restful.ArticleLike)

	// 取消文章点赞
	api.POST("/article/unlike", restful.ArticleUnlike)

	// 收藏文章
	api.POST("/article/star", restful.ArticleStar)

	// 取消收藏
	api.POST("/article/unstar", restful.ArticleUnStar)

	// 添加评论
	api.POST("/comment/add", restful.CommentAdd)

	// 编辑评论
	api.POST("/comment/modify", restful.CommentModify)

	// 删除评论
	api.POST("/comment/del", restful.CommentDel)

	// 评论点赞
	api.POST("/comment/like", restful.CommentLike)

	// 取消评论点赞
	api.POST("/comment/unlike", restful.CommentUnlike)

	// 关注作者
	api.POST("/author/follow", restful.Follow)

	// 取消关注
	api.POST("/author/unfollow", restful.Unfollow)
}

// 非鉴权路由
func RegisterNoAuthRoutes(d *doris.Doris) {
	// 未分组接口
	// 登录路由
	d.POST("/login", restful.Login)

	// 注册路由
	d.POST("/register", restful.Register)

	// 刷新token
	d.POST("/refresh-token", restful.RefreshToken)

	// 重置密码
	d.POST("/reset-pwd", restful.ResetPWD)

	// 首页路由
	d.POST("/", restful.Index)

	// 搜索列表
	d.POST("/search", restful.Search)

	// 获取文章
	d.POST("/article", restful.Article)

	// 文章评论
	d.POST("/article/comments", restful.Comments)
}

// 注册rpc路由
// @TODO...
