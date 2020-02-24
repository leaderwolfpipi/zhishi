package restful

// 关注相关接口
// 关注作者
func Follow() {
	// 初始化结果
	jsonResult := helper.JsonResult{
		Code:    helper.FollowOk,
		Message: helper.StatusText(helper.FollowOk),
	}

	// 绑定参数
	follow := entity.Follow{}
	_ = c.Form(&follow)

	// 实例化service
	repo := mysql.NewRepo(&entity.Follow{}, helper.Database)
	service := service.NewService(repo)

	// 执行插入
	err := service.AuthorFollow(&follow)

	// 结果判断
	if err != nil {
		// 异常状态码返回400
		jsonResult.Code = helper.FollowErr
		jsonResult.Message = helper.StatusText(helper.FollowErr) + err.Error()
	}

	// 返回结果
	c.IndentedJson(http.StatusOK, jsonResult)
}

// 取消关注
func Unfollow() {
	// 初始化结果
	jsonResult := helper.JsonResult{
		Code:    helper.UnFollowOk,
		Message: helper.StatusText(helper.UnFollowOk),
	}

	// 绑定参数
	follow := entity.Follow{}
	_ = c.Form(&follow)

	// 实例化service
	repo := mysql.NewRepo(&entity.Follow{}, helper.Database)
	service := service.NewService(repo)

	// 执行插入
	err := service.AuthorUnFollow(&follow)

	// 结果判断
	if err != nil {
		// 异常状态码返回400
		jsonResult.Code = helper.UnFollowErr
		jsonResult.Message = helper.StatusText(helper.UnFollowErr) + err.Error()
	}

	// 返回结果
	c.IndentedJson(http.StatusOK, jsonResult)
}
