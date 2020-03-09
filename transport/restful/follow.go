package restful

import (
	"errors"
	"net/http"

	"github.com/leaderwolfpipi/doris"
	"github.com/leaderwolfpipi/zhishi/entity"
	"github.com/leaderwolfpipi/zhishi/helper"
	"github.com/leaderwolfpipi/zhishi/service"
	"github.com/leaderwolfpipi/zhishi/service/server/repository/mysql"
)

// 关注相关接口
// 关注作者
func Follow(c *doris.Context) error {
	// 初始化错误
	var err error = nil

	// 初始化结果
	jsonResult := helper.JsonResult{
		Code:    helper.FollowOk,
		Message: helper.StatusText(helper.FollowOk),
	}

	// 绑定参数
	follow := &entity.Follow{}
	_ = c.Form(follow)

	// 参数校验
	if follow.UserId == 0 || follow.FollowedId == 0 {
		err = errors.New("user_id and followed_id cannot be empty!")
	} else {
		// 实例化service
		repo := mysql.NewRepo(follow.GetFollowFunc("add"), helper.Database)
		service := service.NewService(repo)

		// 关注去重
		andWhere := map[string]interface{}{
			"user_id = ? ":     follow.UserId,
			"followed_id = ? ": follow.FollowedId,
		}
		dupli := service.Exist(andWhere)
		if !dupli {
			// 执行插入
			err = service.AuthorFollow(follow)
		}
	}

	// 结果判断
	if err != nil {
		// 异常状态码返回400
		jsonResult.Code = helper.FollowErr
		jsonResult.Message = helper.StatusText(helper.FollowErr) + err.Error()
	}

	// 返回结果
	c.IndentedJson(http.StatusOK, jsonResult)
	return err
}

// 取消关注
func Unfollow(c *doris.Context) error {
	// 初始化错误
	var err error = nil

	// 初始化结果
	jsonResult := helper.JsonResult{
		Code:    helper.UnFollowOk,
		Message: helper.StatusText(helper.UnFollowOk),
	}

	// 绑定参数
	follow := &entity.Follow{}
	_ = c.Form(follow)

	// 参数校验
	if follow.UserId == 0 || follow.FollowedId == 0 {
		err = errors.New("user_id and followed_id cannot be empty!")
	} else {
		// 设置where查询条件
		andWhere := map[string]interface{}{
			"user_id = ? ":     follow.UserId,
			"followed_id = ? ": follow.FollowedId,
		}

		// 实例化service
		repo := mysql.NewRepo(follow.GetFollowFunc("delete"), helper.Database)
		service := service.NewService(repo)

		// 取消关注
		err = service.AuthorUnFollow(andWhere)
	}

	// 结果判断
	if err != nil {
		// 异常状态码返回400
		jsonResult.Code = helper.UnFollowErr
		jsonResult.Message = helper.StatusText(helper.UnFollowErr) + err.Error()
	}

	// 返回结果
	c.IndentedJson(http.StatusOK, jsonResult)
	return err
}
