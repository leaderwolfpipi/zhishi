package restful

// 评论相关接口

import (
	"errors"
	"net/http"

	"github.com/leaderwolfpipi/doris"
	"github.com/leaderwolfpipi/zhishi/entity"
	"github.com/leaderwolfpipi/zhishi/helper"
	"github.com/leaderwolfpipi/zhishi/service"
	"github.com/leaderwolfpipi/zhishi/service/server/repository/mysql"
)

// 获取文章评论
func Comments(c *doris.Context) error {
	// 初始化错误
	var err error = nil

	// 初始化结果集
	jsonResult := helper.JsonResult{
		Code:    helper.ArticleCommentOk,
		Message: helper.StatusText(helper.ArticleCommentOk),
	}

	// 提取分页参数
	var pageResult *helper.PageResult
	_ = c.Form(pageResult)

	// 提取atticle_id
	article_id := c.Param("articleId").(int64)
	andWhere := map[string]interface{}{
		"article_id": article_id,
	}

	// 排序条件
	order := map[string]string{
		"create_time": "desc",
	}

	// 实例化repo对象
	comment := &entity.Comment{}
	repo := mysql.NewRepo(comment.GetCommentFunc("findMore"), helper.Database)

	// 传递repo到service层
	service := service.NewService(repo)

	// 调用service的ArticleByPK接口
	pageResult = service.ArticleComment(nil, andWhere, nil, order, pageResult.PageNum, pageResult.PageSize)
	if pageResult == nil {
		// 异常状态码返回400
		err = errors.New(helper.StatusText(helper.ArticleCommentErr))
		jsonResult.Code = helper.ArticleCommentErr
		jsonResult.Message = helper.StatusText(helper.ArticleCommentErr)
	} else {
		jsonResult.Result = pageResult
	}

	// 返回结果
	c.IndentedJson(http.StatusOK, jsonResult)
	return err
}

// 添加评论
func CommentAdd(c *doris.Context) error {
	// 初始化错误
	var err error = nil

	// 初始化结果集
	jsonResult := helper.JsonResult{
		Code:    helper.CommentAddOk,
		Message: helper.StatusText(helper.CommentAddOk),
	}

	// 绑定内容表
	comment := &entity.Comment{}
	_ = c.Form(comment)

	// 实例化repo对象
	repo := mysql.NewRepo(comment.GetCommentFunc("add"), helper.Database)

	// 传递repo到service层
	service := service.NewService(repo)

	// 调用service的ArticleByPK接口
	err = service.CommentAdd(comment)
	if err != nil {
		// 异常状态码返回400
		jsonResult.Code = helper.CommentAddErr
		jsonResult.Message = helper.StatusText(helper.CommentAddErr) + err.Error()
	}

	// 返回结果
	c.IndentedJson(http.StatusOK, jsonResult)
	return err
}

// 编辑评论
func CommentModify(c *doris.Context) error {
	// 初始化错误
	var err error = nil

	// 初始化结果集
	jsonResult := helper.JsonResult{
		Code:    helper.CommentModifyOk,
		Message: helper.StatusText(helper.CommentModifyOk),
	}

	// 绑定内容表
	comment := &entity.Comment{}
	_ = c.Form(comment)

	// 实例化repo对象
	repo := mysql.NewRepo(comment.GetCommentFunc("update"), helper.Database)

	// 传递repo到service层
	service := service.NewService(repo)

	// 调用service的ArticleByPK接口
	err = service.CommentModify(comment)
	if err != nil {
		// 异常状态码返回400
		jsonResult.Code = helper.CommentModifyErr
		jsonResult.Message = helper.StatusText(helper.CommentModifyErr) + err.Error()
	}

	// 返回结果
	c.IndentedJson(http.StatusOK, jsonResult)
	return err
}

// 删除评论
func CommentDel(c *doris.Context) error {
	// 初始化错误
	var err error = nil

	// 初始化结果集
	jsonResult := helper.JsonResult{
		Code:    helper.ArticleOk,
		Message: helper.StatusText(helper.ArticleOk),
	}

	// 获取参数
	comment := &entity.Comment{}
	articleId := c.Param("articleId").(int64)

	// 实例化repo对象
	repo := mysql.NewRepo(comment.GetCommentFunc("delete"), helper.Database)

	// 传递repo到service层
	service := service.NewService(repo)

	// 调用删除接口
	err = service.CommentDel(articleId)
	if err != nil {
		// 异常状态码返回400
		jsonResult.Code = helper.ArticleDelErr
		jsonResult.Message = helper.StatusText(helper.ArticleDelErr) + err.Error()
	}

	// 返回结果
	c.IndentedJson(http.StatusOK, jsonResult)
	return err
}

// 针对评论点赞
func CommentLike(c *doris.Context) error {
	// 初始化错误
	var err error = nil

	// 初始化结果
	jsonResult := helper.JsonResult{
		Code:    helper.CommentLikeOk,
		Message: helper.StatusText(helper.CommentLikeOk),
	}

	// 绑定参数
	like := &entity.Like{}
	_ = c.Form(like)

	// 实例化service
	repo := mysql.NewRepo(like.GetLikeFunc("add"), helper.Database)
	service := service.NewService(repo)

	// 执行插入
	err = service.CommentLike(like)

	// 结果判断
	if err != nil {
		// 异常状态码返回400
		jsonResult.Code = helper.CommentLikeErr
		jsonResult.Message = helper.StatusText(helper.CommentLikeErr) + err.Error()
	}

	// 返回结果
	c.IndentedJson(http.StatusOK, jsonResult)
	return err
}

// 取消评论点赞
func CommentUnlike(c *doris.Context) error {
	// 初始化错误
	var err error = nil

	// 初始化结果
	jsonResult := helper.JsonResult{
		Code:    helper.CommentUnLikeOk,
		Message: helper.StatusText(helper.CommentUnLikeOk),
	}

	// 绑定参数
	like := &entity.Like{}
	_ = c.Form(like)

	// 实例化service
	repo := mysql.NewRepo(like.GetLikeFunc("delete"), helper.Database)
	service := service.NewService(repo)

	// 执行插入
	err = service.CommentUnlike(like.UserId, like.ObjectId)

	// 结果判断
	if err != nil {
		// 异常状态码返回400
		jsonResult.Code = helper.CommentUnLikeErr
		jsonResult.Message = helper.StatusText(helper.CommentUnLikeErr) + err.Error()
	}

	// 返回结果
	c.IndentedJson(http.StatusOK, jsonResult)
	return err
}
