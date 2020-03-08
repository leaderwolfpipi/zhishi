// repository interface
// 定义依赖的第三方接口
// 接口实现在server子目录中
package repository

import (
	"github.com/leaderwolfpipi/zhishi/helper"
)

// 公共SQL仓库接口
type DbRepo interface {
	// 插入
	Insert(item interface{}) error

	// 更新
	Update(item interface{}) error

	// 更新改变且非空的部分
	UpdateColumns(item interface{}) error

	// 删除
	Delete(
		where map[string]interface{},
		orWhere map[string]interface{}) error

	// 根据主键删除
	DeleteByPK(pk string, value uint64) error

	// 根据结构对象级联删除
	DeleteCascade(obj interface{}) error

	// 主键查找
	FindByPK(
		joinTables map[string]string,
		pk string,
		value uint64,
		preloads map[string]string,
		preLimit int) (interface{}, error)

	// 查找一条
	FindOne(
		andWhere map[string]interface{},
		orWhere map[string]interface{},
		order map[string]string,
		preloads map[string]string,
		preLimit int) (interface{}, error)

	// 查找全部
	Find(
		joinTables map[string]string,
		andWhere map[string]interface{},
		orWhere map[string]interface{},
		order map[string]string) (interface{}, error)

	// 分页查找
	FindPage(
		preloads map[string]string, // 预加载表
		andWhere map[string]interface{},
		orWhere map[string]interface{},
		order map[string]string,
		page int,
		pageSize int) *helper.PageResult

	// 事务操作
	Transaction(map[string]interface{}) error
}

// 公共cache仓库接口
type CacheRepo interface {
	// 设置
	Set(kv interface{}) error

	// 删除
	Del(key string) error

	// 查找
	Get(key string) interface{}

	// 扫描
	Scan(start int, offset int) interface{}
}

// 公共的文档类数据库接口
type DocRepo interface {
	// 插入文档
	Insert(item interface{}) error

	// 删除文档
	Delete(id string) error

	// 查询文档
	Find(where map[string]interface{}) (interface{}, error)
}

// @TODO...
