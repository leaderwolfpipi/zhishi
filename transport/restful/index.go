package restful

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/leaderwolfpipi/doris"
	"github.com/leaderwolfpipi/entity"
	"github.com/leaderwolfpipi/helper"
	"github.com/leaderwolfpipi/service"
	"github.com/leaderwolfpipi/service/server/repository/mysql"
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
	pageNum := c.Form("pageNum")
	pageSize := c.Form("pageSize")

	// 实例化repo对象
	repo := mysql.NewRepo(entity.Article, helper.Database)

	// 传递repo到service层
	service := service.NewService(repo)

	// 调用service的Index接口
	pageResult, err := service.Articles(nil, nil, nil, nil, pageNum, pageSize)
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
