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
	api.Use(middleware.JWT(helper.GetSignKey()))

	// 添加文章
	api.POST("/article/add", restful.ArticleAdd)

	// 编辑文章
	api.POST("/article/:articleId/modify", restful.ArticleModify)

	// 删除文章
	api.POST("/article/:articleId/del", restful.ArticleDel)

	// 点赞文章
	api.POST("/article/:articleId/like", restful.ArticleLike)

	// 取消文章点赞
	api.POST("/article/:articleId/unlike", restful.ArticleUnlike)

	// 收藏文章
	api.POST("/article/:articleId/star", restful.ArticleStar)

	// 取消收藏
	api.POST("/article/:articleId/unstar", restful.ArticleUnStar)

	// 添加评论
	api.POST("/comment/add", restful.CommentAdd)

	// 编辑评论
	api.POST("/comment/:commentId/modify", restful.CommentModify)

	// 删除评论
	api.POST("/comment/:commentId/del", restful.CommentDel)

	// 评论点赞
	api.POST("/comment/:commentId/like", restful.CommentLike)

	// 取消评论点赞
	api.POST("/comment/:commentId/unlike", restful.CommentUnlike)

	// 关注作者
	api.POST("/author/:user_id/follow", restful.Follow)

	// 取消关注
	api.POST("/author/:user_id/unfollow", restful.Unfollow)
}

// 非鉴权路由
func RegisterNoAuthRoutes(d *doris.Doris) {
	// api-no-auth路由组
	api := d.Group("/api-no-auth/v1")

	// 首页路由
	api.POST("/", restful.Index)

	// 登录路由
	api.POST("/login", restful.Login)

	// 注册路由
	api.POST("/register", restful.Register)

	// 搜索列表
	api.POST("/search", restful.Search)

	// 获取文章
	api.GET("/article/:articleId", restful.Article)

	// 文章评论
	api.GET("/article/:articleId/comments", restful.Comments)
}

// 注册rpc路由
// @TODO...
