package restful

import (
	"net/http"

	"github.com/leaderwolfpipi/doris"
	"github.com/leaderwolfpipi/zhishi/entity"
	"github.com/leaderwolfpipi/zhishi/helper"
	"github.com/leaderwolfpipi/zhishi/service"
	"github.com/leaderwolfpipi/zhishi/service/server/repository/mysql"
)

// 首页接口
func Index(c *doris.Context) error {
	var err error = nil
	// 初始化结果集
	jsonResult := helper.JsonResult{
		Code:    helper.ArticlesPageOk,
		Message: helper.StatusText(helper.ArticlesPageOk),
	}
	// 参数获取与校验
	article := &entity.Article{}
	content := entity.ArticleContent{}
	article.ArticleContent = content
	pageResult := &helper.PageResult{}
	_ = c.Form(pageResult)

	// 0默认为第一页
	if pageResult.PageNum == 0 {
		pageResult.PageNum = 1
	}
	// 默认单页100条
	if pageResult.PageSize == 0 {
		pageResult.PageSize = 100
	}

	// 实例化repo对象
	repo := mysql.NewRepo(article.GetArticleFunc("findMore"), helper.Database)

	// 传递repo到service层
	service := service.NewService(repo)

	// 设置预加载模型
	preloads := map[string]string{
		"zs_article_content": "ArticleContent",
		"zs_like":            "Likes",
		"zs_star":            "Stars",
		"zs_comment":         "Comments",
	}

	// 设置排序参数
	orderField := c.FormParam("orderby")
	order := c.FormParam("order")
	if orderField == "" {
		// 默认按创建时间排序
		orderField = "create_time"
	}
	if order == "" {
		// 默认降序
		order = "desc"
	}
	orders := map[string]string{
		orderField: order,
	}

	// 调用service的Index接口
	pageResult = service.Articles(preloads, nil, nil, orders, pageResult.PageNum, pageResult.PageSize)
	pageResult.Total = len(*pageResult.Rows.(*[]entity.Article))

	// 组织返回结果
	if pageResult == nil {
		// 异常状态码返回400
		jsonResult.Code = helper.ArticlesPageErr
		jsonResult.Message = helper.StatusText(helper.ArticlesPageErr)
	} else {
		jsonResult.Result = pageResult
	}

	// 返回结果
	c.IndentedJson(http.StatusOK, jsonResult)

	return err
}
