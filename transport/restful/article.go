package restful

// 文章接口
import (
	"errors"
	"net/http"
	"strconv"

	"github.com/leaderwolfpipi/doris"
	"github.com/leaderwolfpipi/zhishi/entity"
	"github.com/leaderwolfpipi/zhishi/helper"
	"github.com/leaderwolfpipi/zhishi/service"
	"github.com/leaderwolfpipi/zhishi/service/server/repository/mysql"
)

// 获取文章
func Article(c *doris.Context) error {
	// 初始化错误
	var err error = nil

	// 初始化结果集
	jsonResult := helper.JsonResult{
		Code:    helper.ArticleOk,
		Message: helper.StatusText(helper.ArticleOk),
	}

	// 绑定内容表
	// content := entity.ArticleContent{}
	// _ = c.Form(&content)

	// 绑定文章表
	article := &entity.Article{}
	_ = c.Form(article)

	// 将内容组合进文章
	// article.ArticleContent = content

	// 实例化repo对象
	repo := mysql.NewRepo(article.GetArticleFunc("findOne"), helper.Database)

	// 传递repo到service层
	service := service.NewService(repo)

	preloads := map[string]string{
		"zs_article_content": "ArticleContent",
	}

	// 查询记录
	tmp := c.FormParam("pre_limit")
	preLimit, _ := strconv.Atoi(tmp)
	// 子表默认取10条
	if preLimit == 0 {
		preLimit = 10
	}

	// 调用service的ArticleByPK接口
	result, err := service.ArticleByPK("article_id", uint64(article.ID), preloads, preLimit)
	if err != nil {
		// 异常状态码返回400
		jsonResult.Code = helper.ArticleGetErr
		jsonResult.Message = helper.StatusText(helper.ArticleGetErr) + " [ origin err: " + err.Error() + " ]"
	} else {
		jsonResult.Result = result
	}

	// 返回结果
	c.IndentedJson(http.StatusOK, jsonResult)
	return err
}

// 文章添加
func ArticleAdd(c *doris.Context) error {
	// 初始化错误
	var err error = nil

	// 初始化结果集
	jsonResult := helper.JsonResult{
		Code:    helper.ArticleOk,
		Message: helper.StatusText(helper.ArticleOk),
	}

	// 绑定内容表
	content := entity.ArticleContent{}
	_ = c.Form(&content)

	// 绑定文章表
	article := &entity.Article{}
	_ = c.Form(article)

	// 将内容组合进文章
	article.ArticleContent = content

	// 实例化repo对象
	repo := mysql.NewRepo(article.GetArticleFunc("add"), helper.Database)

	// 传递repo到service层
	service := service.NewService(repo)

	// 查询条件
	andWhere := map[string]interface{}{
		"title = ?":        article.Title,
		"user_id = ?":      article.UserId,
		"article_type = ?": article.ArticleType,
	}

	// 设置预加载表
	preloads := map[string]string{
		"zs_article_content": "ArticleContent",
		"zs_like":            "Likes",
		"zs_star":            "Stars",
	}

	// 查询记录
	tmp := c.FormParam("pre_limit")
	preLimit, _ := strconv.Atoi(tmp)
	// 子表默认取10条
	if preLimit == 0 {
		preLimit = 10
	}
	// @TODO... // 使用事务处理
	aDupli, _ := service.ArticleByIndex(andWhere, nil, preloads, preLimit)
	if aDupli != nil {
		// 更新数据
		article.ID = aDupli.(*entity.Article).ID
		article.CreateTime = aDupli.(*entity.Article).CreateTime
		err = service.ArticleModify(article)
	} else {
		// 插入数据
		err = service.ArticleAdd(article)
	}

	if err != nil {
		// 异常状态码返回400
		jsonResult.Code = helper.ArticleAddErr
		jsonResult.Message = helper.StatusText(helper.ArticleAddErr) + err.Error()
	}

	// 返回结果
	c.IndentedJson(http.StatusOK, jsonResult)
	return err
}

// 文章编辑
func ArticleModify(c *doris.Context) error {
	// 初始化错误
	var err error = nil

	// 初始化结果集
	jsonResult := helper.JsonResult{
		Code:    helper.ArticleOk,
		Message: helper.StatusText(helper.ArticleOk),
	}

	// 实例化内容表
	content := entity.ArticleContent{}
	_ = c.Form(&content)

	// 实例化文章表
	article := &entity.Article{}
	_ = c.Form(article)
	article.ArticleContent = content

	// 实例化repo对象
	repo := mysql.NewRepo(article.GetArticleFunc("update"), helper.Database)

	// 传递repo到service层
	service := service.NewService(repo)

	// 查询记录
	tmp := c.FormParam("pre_limit")
	preLimit, _ := strconv.Atoi(tmp)
	// 子表默认取10条
	if preLimit == 0 {
		preLimit = 10
	}

	// 按照主键查询
	aDupli, _ := service.ArticleByPK("article_id", article.ID, nil, preLimit)

	if article.ID == 0 || aDupli == nil {
		// 错误提示
		jsonResult.Code = helper.ArticleIdEmptyErr
		jsonResult.Message = helper.StatusText(helper.ArticleIdEmptyErr)
		c.IndentedJson(http.StatusOK, jsonResult)
		return err
	}

	// 调用service的ArticleSave接口
	err = service.ArticleModify(article)
	if err != nil {
		// 异常状态码返回400
		jsonResult.Code = helper.ArticleModifyErr
		jsonResult.Message = helper.StatusText(helper.ArticleModifyErr) + err.Error()
	}

	// 返回结果
	c.IndentedJson(http.StatusOK, jsonResult)
	return err
}

// 文章删除
func ArticleDel(c *doris.Context) error {
	// 初始化错误
	var err error = nil

	// 初始化结果集
	jsonResult := helper.JsonResult{
		Code:    helper.ArticleOk,
		Message: helper.StatusText(helper.ArticleOk),
	}

	// 实例化内容表
	content := entity.ArticleContent{}
	_ = c.Form(&content)

	// 实例化文章表
	article := &entity.Article{}
	_ = c.Form(article)
	article.ArticleContent = content

	// 实例化repo对象
	repo := mysql.NewRepo(article.GetArticleFunc("delete"), helper.Database)

	// 传递repo到service层
	service := service.NewService(repo)

	err = service.ArticleDel("article_id", article.ID)
	if err != nil {
		// 异常状态码返回400
		jsonResult.Code = helper.ArticleDelErr
		jsonResult.Message = helper.StatusText(helper.ArticleDelErr) + err.Error()
	}

	// 返回结果
	c.IndentedJson(http.StatusOK, jsonResult)
	return err
}

// 文章点赞
func ArticleLike(c *doris.Context) error {
	// 初始化错误
	var err error = nil

	// 初始化结果
	jsonResult := helper.JsonResult{
		Code:    helper.ArticleLikeOk,
		Message: helper.StatusText(helper.ArticleLikeOk),
	}

	// 绑定参数
	like := &entity.Like{}
	_ = c.Form(like)

	// 参数校验
	if like.ArticleId == 0 || like.CommentId != 0 {
		// 异常状态码返回400
		jsonResult.Code = helper.ArticleLikeErr
		jsonResult.Message = helper.StatusText(helper.ArticleLikeErr) + " origin err: [ article_id cannot be empty and comment_id must be empty! ] "
	} else {
		// 实例化service
		repo := mysql.NewRepo(like.GetLikeFunc("add"), helper.Database)
		service := service.NewService(repo)

		// 点赞查重
		andWhere := map[string]interface{}{
			"user_id = ? ":    like.UserId,
			"article_id = ? ": like.ArticleId,
		}

		// 重复检测
		if !service.Exist(andWhere) {
			// 执行插入
			err = service.ArticleLike(like)
			// 结果判断
			if err != nil {
				// 异常状态码返回400
				jsonResult.Code = helper.ArticleLikeErr
				jsonResult.Message = helper.StatusText(helper.ArticleLikeErr) + " origin err: [ " + err.Error() + " ] "
			}
		}
	}

	// 返回结果
	c.IndentedJson(http.StatusOK, jsonResult)
	return err
}

// 取消点赞
func ArticleUnlike(c *doris.Context) error {
	// 初始化结果
	jsonResult := helper.JsonResult{
		Code:    helper.ArticleUnLikeOk,
		Message: helper.StatusText(helper.ArticleUnLikeOk),
	}

	// 绑定参数
	like := &entity.Like{}
	_ = c.Form(like)
	andWhere := map[string]interface{}{
		"user_id = ? ":    like.UserId,
		"article_id = ? ": like.ArticleId,
	}

	// 实例化service
	repo := mysql.NewRepo(like.GetLikeFunc("delete"), helper.Database)
	service := service.NewService(repo)

	// 执行插入
	err := service.ArticleUnlike(andWhere)

	// 结果判断
	if err != nil {
		// 异常状态码返回400
		jsonResult.Code = helper.ArticleUnLikeErr
		jsonResult.Message = helper.StatusText(helper.ArticleUnLikeErr) + err.Error()
	}

	// 返回结果
	c.IndentedJson(http.StatusOK, jsonResult)
	return err
}

// 收藏文章
func ArticleStar(c *doris.Context) error {
	// 初始化错误
	var err error = nil

	// 初始化结果
	jsonResult := helper.JsonResult{
		Code:    helper.ArticleStarOk,
		Message: helper.StatusText(helper.ArticleStarOk),
	}

	// 绑定参数
	star := &entity.Star{}
	_ = c.Form(star)

	// 校验参数
	if star.UserId == 0 || star.ArticleId == 0 {
		err = errors.New("user_id or article_id cannot be empty!")
	} else {
		// 实例化service
		repo := mysql.NewRepo(star.GetStarFunc("add"), helper.Database)
		service := service.NewService(repo)

		// 收藏查重
		andWhere := map[string]interface{}{
			"user_id = ? ":    star.UserId,
			"article_id = ? ": star.ArticleId,
		}
		dupli := service.Exist(andWhere)
		if !dupli {
			// 执行插入
			err = service.ArticleStar(star)
		}
	}

	// 结果判断
	if err != nil {
		// 异常状态码返回400
		jsonResult.Code = helper.ArticleStarErr
		jsonResult.Message = helper.StatusText(helper.ArticleStarErr) + " [ origin err: " + err.Error() + " ]"
	}

	// 返回结果
	c.IndentedJson(http.StatusOK, jsonResult)
	return err
}

// 取消收藏文章
func ArticleUnStar(c *doris.Context) error {
	// 初始化错误
	var err error = nil

	// 初始化结果
	jsonResult := helper.JsonResult{
		Code:    helper.ArticleUnStarOk,
		Message: helper.StatusText(helper.ArticleUnStarOk),
	}

	// 绑定参数
	star := &entity.Star{}
	_ = c.Form(star)
	andWhere := map[string]interface{}{
		"user_id = ? ":    star.UserId,
		"article_id = ? ": star.ArticleId,
	}

	// 参数检验
	// @TODO... 参数校验应放到绑定环节
	if star.UserId == 0 || star.ArticleId == 0 {
		err = errors.New("user_id or article_id cannot be empty!")
	} else {
		// 实例化service
		repo := mysql.NewRepo(star.GetStarFunc("delete"), helper.Database)
		service := service.NewService(repo)

		// 执行删除
		err = service.ArticleUnStar(andWhere)
	}

	// 结果判断
	if err != nil {
		// 异常状态码返回400
		jsonResult.Code = helper.ArticleUnStarErr
		jsonResult.Message = helper.StatusText(helper.ArticleUnStarErr) + err.Error()
	}

	// 返回结果
	c.IndentedJson(http.StatusOK, jsonResult)
	return err
}
