package restful

// 文章接口
import (
	"net/http"

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

	// 提取atticle_id
	article_id := c.Param("articleId").(int64)
	article := &entity.Article{}

	// 实例化repo对象
	repo := mysql.NewRepo(article.GetArticleFunc("findOne"), helper.Database)

	// 传递repo到service层
	service := service.NewService(repo)

	// 调用service的ArticleByPK接口
	result, err := service.ArticleByPK("article_id", article_id, "Content")
	if err != nil {
		// 异常状态码返回400
		jsonResult.Code = helper.ArticleGetErr
		jsonResult.Message = helper.StatusText(helper.ArticleGetErr)
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

	// 调用service的ArticleByPK接口
	err = service.ArticleAdd(article)

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

	// 获取参数
	articleId := c.Param("articleId").(int64)
	article := &entity.Article{}

	// 实例化repo对象
	repo := mysql.NewRepo(article.GetArticleFunc("delete"), helper.Database)

	// 传递repo到service层
	service := service.NewService(repo)

	// 调用删除接口
	err = service.ArticleDel(articleId)
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

	// 实例化service
	repo := mysql.NewRepo(like.GetLikeFunc("add"), helper.Database)
	service := service.NewService(repo)

	// 执行插入
	err = service.ArticleLike(like)

	// 结果判断
	if err != nil {
		// 异常状态码返回400
		jsonResult.Code = helper.ArticleLikeErr
		jsonResult.Message = helper.StatusText(helper.ArticleLikeErr) + err.Error()
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
		"user_id":   like.UserId,
		"object_id": like.ObjectId,
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

	// 实例化service
	repo := mysql.NewRepo(star.GetStarFunc("add"), helper.Database)
	service := service.NewService(repo)

	// 执行插入
	err = service.ArticleStar(star)

	// 结果判断
	if err != nil {
		// 异常状态码返回400
		jsonResult.Code = helper.ArticleStarErr
		jsonResult.Message = helper.StatusText(helper.ArticleStarErr) + err.Error()
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
		"user_id":    star.UserId,
		"article_id": star.ArticleId,
	}

	// 实例化service
	repo := mysql.NewRepo(star.GetStarFunc("delete"), helper.Database)
	service := service.NewService(repo)

	// 执行插入
	err = service.ArticleUnStar(andWhere)

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
