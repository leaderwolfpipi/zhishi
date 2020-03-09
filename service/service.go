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
	UserByPK(
		pk string,
		value uint64,
		preloads map[string]string,
		preLimit int) (interface{}, error)

	// 根据username查询
	UserByUsername(username string) (interface{}, error)

	// 根据telephone查询
	UserByTelephone(telephone string) (interface{}, error)

	//获取用户列表
	Users(
		preloads map[string]string,
		andWhere map[string]interface{},
		orWhere map[string]interface{},
		order map[string]string,
		pageNum int,
		pageSize int) *helper.PageResult

	// 添加文章
	ArticleAdd(article interface{}) error

	// 获取单篇文章
	ArticleByPK(
		pk string,
		value uint64,
		preloads map[string]string,
		preLimit int) (interface{}, error)

	// 根据索引搜索一条记录
	// 由于or查询会造成索引失效因此尽量不要用or条件
	ArticleByIndex(
		andWhere map[string]interface{},
		order map[string]string,
		preloads map[string]string,
		preLimit int) (interface{}, error)

	// 分页获取文章
	Articles(
		preloads map[string]string,
		andWhere map[string]interface{},
		orWhere map[string]interface{},
		order map[string]string,
		pageNum int,
		pageSize int) *helper.PageResult

	// 编辑文章
	ArticleModify(article interface{}) error

	// 删除文章
	ArticleDel(pk string, value uint64) error

	// 点赞文章
	ArticleLike(like interface{}) error

	// 取消点赞
	ArticleUnlike(andWhere map[string]interface{}) error

	// 收藏文章
	ArticleStar(star interface{}) error

	// 取消收藏
	ArticleUnStar(andWhere map[string]interface{}) error

	// 获取文章评论
	ArticleComments(
		preloads map[string]string,
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
	CommentDel(commentId uint64) error

	// 点赞评论
	CommentLike(like interface{}) error

	// 取消点赞
	CommentUnlike(userId uint64, commentId uint64) error

	// 关注作者
	AuthorFollow(follow interface{}) error

	// 取消关注
	AuthorUnFollow(andWhere map[string]interface{}) error

	// 存在判断
	Exist(andWhere map[string]interface{}) bool
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
func (s *service) UserByPK(
	pk string,
	value uint64,
	preloads map[string]string,
	preLimit int) (interface{}, error) {
	return s.repo.FindByPK(nil, pk, value, preloads, preLimit)
}

// 根据用户名查询用户
func (s *service) UserByUsername(username string) (interface{}, error) {
	andWhere := map[string]interface{}{
		"username = ? ": username,
	}
	return s.repo.FindOne(andWhere, nil, nil, nil, 10)
}

// 根据电话号码查询
func (s *service) UserByTelephone(telephone string) (interface{}, error) {
	andWhere := map[string]interface{}{
		"telephone=?": telephone,
	}
	return s.repo.FindOne(nil, andWhere, nil, nil, 10)
}

// 获取用户列表
func (s *service) Users(
	preloads map[string]string,
	andWhere map[string]interface{},
	orWhere map[string]interface{},
	order map[string]string,
	pageNum int,
	pageSize int) *helper.PageResult {
	return s.repo.FindPage(preloads, andWhere, orWhere, order, pageNum, pageSize)
}

// 首页接口
func (s *service) Articles(
	preloads map[string]string,
	andWhere map[string]interface{},
	orWhere map[string]interface{},
	order map[string]string,
	pageNum int,
	pageSize int) *helper.PageResult {
	return s.repo.FindPage(preloads, andWhere, orWhere, order, pageNum, pageSize)
}

// 获取单篇文章
func (s *service) ArticleByPK(
	pk string,
	value uint64,
	preloads map[string]string,
	preLimit int) (interface{}, error) {
	return s.repo.FindByPK(nil, pk, value, preloads, preLimit)
}

// 基于索引条件查询一条数据
func (s *service) ArticleByIndex(
	andWhere map[string]interface{},
	order map[string]string,
	preloads map[string]string,
	preLimit int) (interface{}, error) {
	return s.repo.FindOne(andWhere, nil, order, preloads, preLimit)
}

// 文章评论列表
func (s *service) ArticleComments(
	preloads map[string]string,
	andWhere map[string]interface{},
	orWhere map[string]interface{},
	order map[string]string,
	pageNum int,
	pageSize int) *helper.PageResult {
	return s.repo.FindPage(preloads, andWhere, orWhere, order, pageNum, pageSize)
}

// 添加文章
func (s *service) ArticleAdd(article interface{}) error {
	err := s.repo.Insert(article)
	return err
}

// 编辑文章
func (s *service) ArticleModify(article interface{}) error {
	return s.repo.UpdateColumns(article)
}

// 删除文章
func (s *service) ArticleDel(pk string, value uint64) error {
	// 级联删除
	return s.repo.DeleteByPK(pk, value)
}

// 点赞查重
func (s *service) Exist(andWhere map[string]interface{}) bool {
	dupli, _ := s.repo.FindOne(andWhere, nil, nil, nil, 1)
	if dupli != nil {
		return true
	}
	return false
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
func (s *service) CommentDel(commentId uint64) error {
	return s.repo.DeleteByPK("comment_id", commentId)
}

// 评论点赞
func (s *service) CommentLike(like interface{}) error {
	return s.repo.Insert(like)
}

// 取消评论点赞
func (s *service) CommentUnlike(userId uint64, commentId uint64) error {
	andWhere := map[string]interface{}{
		"user_id = ? ":    userId,
		"comment_id = ? ": commentId,
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
