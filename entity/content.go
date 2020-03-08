package entity

import (
	"github.com/leaderwolfpipi/zhishi/helper"
)

// 文章内容表
type ArticleContent struct {
	// 主键: 文章id
	ArticleId uint64 `gorm:"type:bigint(20);primary_key;column:article_id;" json:"article_id" param:"article_id"`
	// 文章内容
	Content string `gorm:"type:text;column:content;" json:"content" param:"content" form:"content"`
	// 创建时间
	CreateTime int `gorm:"type:int(11);column:create_time;default:0;" json:"create_time"`
	// 更新时间
	LastUpdateTime int `gorm:"type:int(11);column:last_update_time;default:0;" json:"last_update_time"`
}

// 初始函数执行建表和添加外键
func init() {
	// 执行数据迁移
	// helper.Database.AutoMigrate(&ArticleContent{})

	// 设置外键约束
	helper.Database.Model(&ArticleContent{}).AddForeignKey("article_id", "zs_article(article_id)", "CASCADE", "CASCADE")
}

// 设置表名
func (articleContent *ArticleContent) TableName() string {
	return "zs_article_content"
}

// 设置创建钩子
// 插入前生成主键并自动设置插入时间
//func (article *ArticleContent) BeforeCreate(scope *gorm.Scope) error {
//	// 设置create_time（单位秒）
//	scope.SetColumn("CreateTime", time.Now().Unix())

//	return nil
//}

// 获取实体处理函数
func (ac *ArticleContent) GetContentFunc(action string) helper.EntityFunc {
	return func() interface{} {
		var ret interface{}
		if action == "findMore" {
			ret = &[]ArticleContent{}
		} else {
			ret = &ArticleContent{}
		}

		return ret
	}
}
