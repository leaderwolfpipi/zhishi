package restful

// 搜索相关接口

import (
	"net/http"

	// "github.com/jinzhu/gorm"
	"github.com/leaderwolfpipi/doris"
	"github.com/leaderwolfpipi/zhishi/entity"
	"github.com/leaderwolfpipi/zhishi/helper"
	"github.com/leaderwolfpipi/zhishi/service"
	"github.com/leaderwolfpipi/zhishi/service/server/repository/mysql"
)

// 首页接口
func Search(c *doris.Context) error {
	var err error = nil
	// 初始化结果集
	jsonResult := helper.JsonResult{
		Code:    helper.ArticlesPageOk,
		Message: helper.StatusText(helper.ArticlesPageOk),
	}
	// 参数获取与校验
	article := &entity.Article{}
	pageResult := &helper.PageResult{}
	_ = c.Form(pageResult)

	// 获取搜索关键词
	_ = c.Form(article)

	// 实例化repo对象
	repo := mysql.NewRepo(article.GetArticleFunc("findMore"), helper.Database)

	// 传递repo到service层
	service := service.NewService(repo)

	// 设置and条件
	andWhere := map[string]interface{}{
		"title LIKE ?": article.Title + "%",
	}

	// 调用service的Index接口
	pageResult = service.Articles(nil, andWhere, nil, nil, pageResult.PageNum, pageResult.PageSize)
	if err != nil {
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
