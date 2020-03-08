package entity

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/leaderwolfpipi/zhishi/helper"
	// "github.com/shopspring/decimal"
)

// 文章实体表
type Article struct {
	// 主键
	ID uint64 `gorm:"type:bigint(20);primary_key;column:article_id" json:"article_id" param:"article_id" validate:"required"`
	// 用户名
	Title string `gorm:"type:varchar(50);unique_index;not null;column:title;" json:"title" param:"title" form:"title" validate:"required"`
	// 作者id
	UserId uint64 `gorm:"type:bigint(20);column:user_id;" json:"user_id" param:"user_id"`
	// 用户状态
	ArticleType uint8 `gorm:"type:tinyint(4);column:article_type;default:1;" json:"article_type" param:"article_type" form:"article_type"`
	// 文章价格
	Price string `gorm:"type:varchar(20);column:price;" json:"price" param:"price" form:"price"`
	// 是否免费
	IsFree uint8 `gorm:"type:tinyint(4);column:is_free;default:1;" json:"is_free" param:"is_free" form:"is_free"`
	// 文章状态
	Status uint8 `gorm:"type:tinyint(4);column:status;default:1;" json:"status" param:"status" form:"status"`
	// 创建时间
	CreateTime int `gorm:"type:int(11);column:create_time;default:0;" json:"create_time"`
	// 更新时间
	LastUpdateTime int `gorm:"type:int(11);column:last_update_time;default:0;" json:"last_update_time"`
	// 关联内容表（1:1）
	ArticleContent ArticleContent `gorm:"ForeignKey:ArticleId;association_foreignkey:ID" json:"content"`
	// 关联like表
	Likes []Like `gorm:"ForeignKey:ObjectId;association_foreignkey:ID" json:"likes"`
	// 关联star表
	Stars []Star `gorm:"ForeignKey:ArticleId;association_foreignkey:ID" json:"stars"`
	// 关联Comment表
	Comments []Comment `gorm:"ForeignKey:ArticleId;association_foreignkey:ID" json:"comments"`
}

// 初始函数执行建表和添加外键
func init() {
	// 执行数据迁移
	// helper.Database.AutoMigrate(&Article{})

	// 设置外键约束
	helper.Database.Model(&Article{}).AddForeignKey("user_id", "zs_user(user_id)", "CASCADE", "CASCADE")
}

// 设置表名
func (article *Article) TableName() string {
	return "zs_article"
}

// 设置创建钩子
// 插入前生成主键并自动设置插入时间
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	nodeId, _ := strconv.Atoi(helper.GetNodesConfig().Server[0].NodeId)
	id, err := helper.GenerateIdBySnowflake(int64(nodeId))
	if err != nil {
		panic("生成ID时发生异常: %s" + err.Error())
	}
	scope.Set("ID", &id)
	article.ID = id

	// 设置create_time（单位秒）
	scope.SetColumn("CreateTime", time.Now().Unix())

	return nil
}

// 设置更新钩子
// 更新操作时自动更新last_update_time（单位秒）
func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	// 设置更新时间
	scope.SetColumn("LastUpdateTime", time.Now().Unix())
	return nil
}

// 获取实体处理函数
func (article *Article) GetArticleFunc(action string) helper.EntityFunc {
	return func() interface{} {
		var ret interface{}
		if action == "findMore" {
			ret = &[]Article{}
		} else {
			ret = &Article{}
		}

		return ret
	}
}
