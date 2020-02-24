package mysql

import (
	"helper"
	"strings"

	"github.com/jinzhu/gorm"
)

// mysql repo implement DbRepo
type repo struct {
	// 实体对象
	// 由service层传入
	entity interface{}

	// 数据库实例
	db *gorm.DB
}

var r repo = &repo{}

// 实例化mysql repo
func NewRepo(entity interface{}, db *gorm.DB) *repo {
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
	err := r.db.Save(item).Error
	return err
}

// 根据where条件删除
func (r *repo) Delete(
	andWhere map[string]interface{},
	orWhere map[string]interface{}) error {
	// 设置删除条件
	for k, v := range andWhere {
		r.db = r.db.Where(k, v)
	}
	err := r.db.Delete(r.entity).Error
	return err
}

// 根据主键删除
func (r *repo) DeleteByPK(pk string, value int64) error {
	key := pk + " = ?"
	err := r.db.Where(key, value).Delete(r.entity).Error
	return err
}

// 主键查找
func (r *repo) FindByPK(
	joinTables map[string]string,
	pk string,
	value int64,
	joinTable2 string) (interface{}, error) {
	// 联表查询
	// 优先使用join查询语句
	if joinTables != nil && len(joinTables) > 0 {
		for k, v := range joinTables {
			r.db = r.db.Joins("left join " + k + " on " + v)
		}
	}

	// joinTable没传时使用外键查询
	if joinTables == nil && len(joinTable2) > 0 {
		r.db = r.db.Preload(joinTable2)
	}

	// 设置where查询条件
	key := pk + " = ?"
	err := r.db.Where(key, value).First(r.entity).Error

	if err != nil {
		// 查询出错
		return nil, err
	}

	return r.entity, nil
}

// 查找一条
func (r *repo) FindOne(
	joinTable map[string]string,
	andWhere map[string]interface{},
	orWhere map[string]interface{},
	order map[string]string) (interface{}, error) {
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

	// 查询单条记录
	err := r.db.First(r.entity).Error

	if err != nil {
		return nil, err
	}

	return r.entity, nil
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
	rows := make([]repo.entity, 0)
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
	err := r.db.find(&rows).Count(&count).Error

	if err != nil {
		return nil, err
	}

	return rows, nil
}

// 分页查找
func (r *repo) FindPage(joinTable map[string]string,
	andWhere map[string]interface{},
	orWhere map[string]interface{},
	order map[string]string,
	pageNum int,
	pageSize int) *helper.PageResult {
	// 存储变量
	rows := make([]repo.entity, 0)
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

	// 设置分页
	if pageNum > 0 && pageSize > 0 {
		r.db = r.db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}

	// 执行查询
	r.db.find(&rows).Count(&count)

	pageData := &PageResult{PageNum: pageNum, PageSize: pageSize, Total: count, Rows: rows}
	return pageData

}
