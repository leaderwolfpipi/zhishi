// repository interface
// 定义依赖的第三方接口
// 接口实现在server子目录中
package repository

import (
	"helper"
)

// 公共SQL仓库接口
type DbRepo interface {
	// 插入
	Insert(item interface{}) error

	// 更新
	Update(item interface{}) error

	// 删除
	Delete(where map[string]interface{}) error

	// 根据主键删除
	DeleteByPK(pk string, value int64) error

	// 主键查找
	FindByPK(
		joinTables map[string]string,
		pk string,
		value int64,
		joinTable2 string) (interface{}, error)

	// 查找一条
	FindOne(
		joinTables map[string]string,
		andWhere map[string]interface{},
		orWhere map[string]interface{},
		order map[string]string) (interface{}, error)

	// 查找全部
	Find(
		joinTables map[string]string,
		andWhere map[string]interface{},
		orWhere map[string]interface{},
		order map[string]string) (interface{}, error)

	// 分页查找
	FindPage(joinTables map[string]string,
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
	Find(where map[string]interface{})(interface,error)
}

// @TODO...
