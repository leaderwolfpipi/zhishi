package mysql

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/leaderwolfpipi/zhishi/helper"
)

// mysql repo implement DbRepo
type repo struct {
	// 实体对象
	// 由service层传入
	entity helper.EntityFunc

	// 数据库实例
	db *gorm.DB
}

var r *repo = &repo{}

// 实例化mysql repo
func NewRepo(entity helper.EntityFunc, db *gorm.DB) *repo {
	r.entity = entity
	r.db = db
	return r
}

// 插入
func (r *repo) Insert(item interface{}) error {
	err := r.db.Create(item).Error
	return err
}

// 事务操作
// 参数：map[string]interface{}键是操作类型，值是操作实体
func (r *repo) Transaction(trans map[string]interface{}) error {
	// 初始化错误
	var err error = nil

	// 开启事务
	tx := r.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	// 退出操作
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 循环操作
	for k, v := range trans {
		if strings.ToLower(k) == "add" {
			// 插入操作
			if err := tx.Create(v).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else if strings.ToLower(k) == "update" {
			// 更新操作
			if err := tx.Save(v).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else if strings.ToLower(k) == "delete" {
			// 删除操作
			if err := tx.Delete(v).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// 提交事务
	tx.Commit()

	return err
}

// 更新
func (r *repo) Update(item interface{}) error {
	// 剔除对create_time的更新
	err := r.db.Omit("create_time").Save(item).Error
	return err
}

// 更新改变的并且非空字段
// 会绕过钩子函数
func (r *repo) UpdateColumns(item interface{}) error {
	entity := r.entity() // 获取实体
	// 开启调试模式
	// return r.db.Model(entity).Debug().Updates(item).Error

	return r.db.Model(entity).Updates(item).Error
}

// 根据where条件删除
func (r *repo) Delete(
	andWhere map[string]interface{},
	orWhere map[string]interface{}) error {
	// 设置and删除条件
	for k, v := range andWhere {
		r.db = r.db.Where(k, v)
	}

	// 设置or删除条件
	for k, v := range orWhere {
		r.db = r.db.Or(k, v)
	}

	// 调用获取实体的方法
	entity := r.entity()
	err := r.db.Delete(entity).Error

	return err
}

// 根据主键删除
func (r *repo) DeleteByPK(pk string, value uint64) error {
	key := pk + " = ?"
	entity := r.entity()
	err := r.db.Where(key, value).Delete(entity).Error
	return err
}

// 级联删除
func (r *repo) DeleteCascade(obj interface{}) error {
	entity := r.entity()
	return r.db.Delete(entity).Error
}

// 主键查找
func (r *repo) FindByPK(
	joinTables map[string]string,
	pk string,
	value uint64,
	preloads map[string]string,
	preLimit int) (interface{}, error) {
	// 联表查询
	// 优先使用join查询语句
	if joinTables != nil && len(joinTables) > 0 {
		for k, v := range joinTables {
			r.db = r.db.Joins("left join " + k + " on " + v)
		}
	}

	// 定制preload
	pcall := func(key string) func(*gorm.DB) *gorm.DB {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order(key + ".create_time DESC").Limit(preLimit)
		}
	}

	// 联表查询
	if preloads != nil && len(preloads) > 0 {
		for k, v := range preloads {
			r.db = r.db.Preload(v, pcall(k))
		}
	}

	// 设置where查询条件
	row := r.entity()
	key := pk + " = ?"
	err := r.db.Where(key, value).First(row).Error

	if err != nil {
		// 查询出错
		return nil, err
	}

	return row, nil
}

// 查找一条
func (r *repo) FindOne(
	andWhere map[string]interface{},
	orWhere map[string]interface{},
	order map[string]string,
	preloads map[string]string,
	preLimit int) (interface{}, error) {
	// 联表查询
	//	if joinTable != nil && len(joinTable) > 0 {
	//		for k, v := range joinTable {
	//			r.db = r.db.Joins("left join " + k + " on " + v)
	//		}
	//	}

	// 定制preload
	pcall := func(key string) func(*gorm.DB) *gorm.DB {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order(key + ".create_time DESC").Limit(preLimit)
		}
	}

	// 联表查询
	if preloads != nil && len(preloads) > 0 {
		for k, v := range preloads {
			r.db = r.db.Preload(v, pcall(k))
		}
	}

	// 设置and条件
	if andWhere != nil && len(andWhere) > 0 {
		for k, v := range andWhere {
			r.db = r.db.Where(k, v)
		}
	}

	// 设置or条件
	// 至少有一个and条件才可设置
	if andWhere != nil && len(andWhere) > 0 && orWhere != nil && len(orWhere) > 0 {
		for k, v := range orWhere {
			r.db = r.db.Or(k, v)
		}
	}

	// 设置排序
	if order != nil && len(order) > 0 {
		for k, v := range order {
			r.db.Order(k + " " + v)
		}
	}

	// 查询单条记录
	row := r.entity()
	err := r.db.First(row).Error

	if err != nil {
		return nil, err
	}

	return row, nil
}

// 查找全部
// 建议使用FindPage执行分页查询
// 避免返回数据过多对数据库造成压力
func (r *repo) Find(
	joinTable map[string]string,
	andWhere map[string]interface{},
	orWhere map[string]interface{},
	order map[string]string) (interface{}, error) {
	// 存储变量
	rows := r.entity()
	count := 0

	// 联表查询
	if joinTable != nil && len(joinTable) > 0 {
		for k, v := range joinTable {
			r.db = r.db.Joins("left join " + k + " on " + v)
		}
	}

	// 设置and条件
	if andWhere != nil && len(andWhere) > 0 {
		for k, v := range andWhere {
			r.db = r.db.Where(k, v)
		}
	}

	// 设置or条件
	// 至少有一个and条件才可设置
	if andWhere != nil && len(andWhere) > 0 && orWhere != nil && len(orWhere) > 0 {
		for k, v := range orWhere {
			r.db = r.db.Or(k, v)
		}
	}

	// 设置排序
	if order != nil && len(order) > 0 {
		for k, v := range order {
			r.db.Order(k + " " + v)
		}
	}

	// 返回满足条件的全部记录
	err := r.db.Find(rows).Count(&count).Error

	if err != nil {
		return nil, err
	}

	return rows, nil
}

// 分页查找
func (r *repo) FindPage(
	preloads map[string]string,
	andWhere map[string]interface{},
	orWhere map[string]interface{},
	order map[string]string,
	pageNum int,
	pageSize int) *helper.PageResult {
	// 存储变量
	rows := r.entity()
	count := 0

	// 定制preload
	pcall := func(key string) func(*gorm.DB) *gorm.DB {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order(key + ".create_time DESC").Limit(pageSize)
		}
	}

	// 联表查询
	if preloads != nil && len(preloads) > 0 {
		for k, v := range preloads {
			r.db = r.db.Preload(v, pcall(k))
		}
	}

	// 设置and条件
	if andWhere != nil && len(andWhere) > 0 {
		for k, v := range andWhere {
			r.db = r.db.Where(k, v)
		}
	}

	// 设置or条件
	// 至少有一个and条件才可设置
	if andWhere != nil && len(andWhere) > 0 && orWhere != nil && len(orWhere) > 0 {
		for k, v := range orWhere {
			r.db = r.db.Or(k, v)
		}
	}

	// 设置排序
	if order != nil && len(order) > 0 {
		for k, v := range order {
			r.db = r.db.Order(k + " " + v)
		}
	}

	// 设置分页
	if pageNum > 0 && pageSize > 0 {
		r.db = r.db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}

	// 执行查询
	// 分页查询和统计数目冲突
	// count数目计算放到上游
	// r.db.Find(rows).Count(&count)
	r.db.Find(rows)
	pageData := &helper.PageResult{PageNum: pageNum, PageSize: pageSize, Total: count, Rows: rows}

	return pageData
}
