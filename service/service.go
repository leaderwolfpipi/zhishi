package service

import (
	"github.com/leaderwolfpipi/zhishi/helper"
	"github.com/leaderwolfpipi/zhishi/service/server/repository"
)

// 定义全局服务接口：domain interface
type Service interface {
	// 添加用户
	UserAdd(user interface{}) error

	// 修改用户
	UserModify(user interface{}) error

	// 根据pk查询
	UserByPK(pk string, value int64, joinTable string) (interface{}, error)

	// 根据username查询
	UserByUsername(username string) (interface{}, error)

	// 根据telephone查询
	UserByTelephone(telephone string) (interface{}, error)

	//获取用户列表
	Users(
		joinTables map[string]string,
		andWhere map[string]interface{},
		orWhere map[string]interface{},
		order map[string]string,
		pageNum int,
		pageSize int) *helper.PageResult

	// 添加文章
	ArticleAdd(article interface{}) error

	// 获取单篇文章
	ArticleByPK(pk string, value int64, joinTable2 string) (interface{}, error)

	// 分页获取文章
	Articles(
		joinTables map[string]string,
		andWhere map[string]interface{},
		orWhere map[string]interface{},
		order map[string]string,
		pageNum int,
		pageSize int) *helper.PageResult

	// 编辑文章
	ArticleModify(article interface{}) error

	// 删除文章
	ArticleDel(articleId int64) error

	// 点赞文章
	ArticleLike(like interface{}) error

	// 取消点赞
	ArticleUnlike(andWhere map[string]interface{}) error

	// 收藏文章
	ArticleStar(star interface{}) error

	// 取消收藏
	ArticleUnStar(andWhere map[string]interface{}) error

	// 获取文章评论
	ArticleComment(
		joinTables map[string]string,
		andWhere map[string]interface{},
		orWhere map[string]interface{},
		order map[string]string,
		pageNum int,
		pageSize int) *helper.PageResult

	// 添加评论
	CommentAdd(comment interface{}) error

	// 编辑评论
	CommentModify(comment interface{}) error

	// 删除评论
	CommentDel(commentId int64) error

	// 点赞评论
	CommentLike(like interface{}) error

	// 取消点赞
	CommentUnlike(userId uint64, commentId uint64) error

	// 关注作者
	AuthorFollow(follow interface{}) error

	// 取消关注
	AuthorUnFollow(andWhere map[string]interface{}) error
}

// 定义Service接口的实现结构
type service struct {
	repo repository.DbRepo
}

var serv *service = &service{}

// 实例化Service对象
func NewService(repo repository.DbRepo) Service {
	serv.repo = repo
	return serv
}

// 添加用户
func (s *service) UserAdd(user interface{}) error {
	return s.repo.Insert(user)
}

// 修改用户
func (s *service) UserModify(user interface{}) error {
	return s.repo.Update(user)
}

// 根据主键查询用户
func (s *service) UserByPK(pk string, value int64, joinTable2 string) (interface{}, error) {
	return s.repo.FindByPK(nil, pk, value, joinTable2)
}

// 根据用户名查询用户
func (s *service) UserByUsername(username string) (interface{}, error) {
	andWhere := map[string]interface{}{
		"username": username,
	}
	return s.repo.FindOne(nil, andWhere, nil, nil)
}

// 根据电话号码查询
func (s *service) UserByTelephone(telephone string) (interface{}, error) {
	andWhere := map[string]interface{}{
		"telephone": telephone,
	}
	return s.repo.FindOne(nil, andWhere, nil, nil)
}

// 获取用户列表
func (s *service) Users(
	joinTables map[string]string,
	andWhere map[string]interface{},
	orWhere map[string]interface{},
	order map[string]string,
	pageNum int,
	pageSize int) *helper.PageResult {
	return s.repo.FindPage(joinTables, andWhere, orWhere, order, pageNum, pageSize)
}

// 首页接口
func (s *service) Articles(
	joinTables map[string]string,
	andWhere map[string]interface{},
	orWhere map[string]interface{},
	order map[string]string,
	pageNum int,
	pageSize int) *helper.PageResult {
	return s.repo.FindPage(joinTables, andWhere, orWhere, order, pageNum, pageSize)
}

// 获取单篇文章
func (s *service) ArticleByPK(pk string, value int64, joinTable2 string) (interface{}, error) {
	return s.repo.FindByPK(nil, pk, value, joinTable2)
}

// 文章评论列表
func (s *service) ArticleComment(
	joinTables map[string]string,
	andWhere map[string]interface{},
	orWhere map[string]interface{},
	order map[string]string,
	pageNum int,
	pageSize int) *helper.PageResult {
	return s.repo.FindPage(joinTables, andWhere, orWhere, order, pageNum, pageSize)
}

// 添加文章
func (s *service) ArticleAdd(article interface{}) error {
	err := s.repo.Insert(article)
	return err
}

// 编辑文章
func (s *service) ArticleModify(article interface{}) error {
	return s.repo.Update(article)
}

// 删除文章
func (s *service) ArticleDel(articleId int64) error {
	return s.repo.DeleteByPK("article_id", articleId)
}

// 文章点赞
func (s *service) ArticleLike(like interface{}) error {
	return s.repo.Insert(like)
}

// 取消文章点赞
func (s *service) ArticleUnlike(andWhere map[string]interface{}) error {
	return s.repo.Delete(andWhere, nil)
}

// 收藏文章
func (s *service) ArticleStar(star interface{}) error {
	return s.repo.Insert(star)
}

// 取消收藏文章
func (s *service) ArticleUnStar(andWhere map[string]interface{}) error {
	return s.repo.Delete(andWhere, nil)
}

// 添加评论
func (s *service) CommentAdd(comment interface{}) error {
	return s.repo.Insert(comment)
}

// 修改评论
func (s *service) CommentModify(comment interface{}) error {
	return s.repo.Update(comment)
}

// 删除评论
func (s *service) CommentDel(commentId int64) error {
	return s.repo.DeleteByPK("comment_id", commentId)
}

// 评论点赞
func (s *service) CommentLike(like interface{}) error {
	return s.repo.Insert(like)
}

// 取消评论点赞
func (s *service) CommentUnlike(userId uint64, commentId uint64) error {
	andWhere := map[string]interface{}{
		"user_id":   userId,
		"object_id": commentId,
	}
	return s.repo.Delete(andWhere, nil)
}

// 关注作者
func (s *service) AuthorFollow(follow interface{}) error {
	return s.repo.Insert(follow)
}

// 取消关注作者
func (s *service) AuthorUnFollow(andWhere map[string]interface{}) error {
	return s.repo.Delete(andWhere, nil)
}

// END OF SERVICE
