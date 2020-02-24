package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/leaderwolfpipi/zhishi/helper"
	"github.com/shopspring/decimal"
)

// 文章内容表
type ArticleContent struct {
	// 主键
	ID uint64 `gorm:"type:bigint(20);primary_key;column:content_id;" json:"content_id" param:"content_id" validate:"required"`
	// 文章id
	ArticleId uint64 `gorm:"type:bigint(20);column:article_id;" json:"article_id" param:"article_id"`
	// 内容
	Content string `gorm:"type:text;column:content;" json:"content" param:"content" form:"content"`
}

// 设置表名
func (articleContent *ArticleContent) TableName() string {
	return "zs_article_content"
}

// 设置创建钩子
// 插入前生成主键并自动设置插入时间
func (content *ArticleContent) BeforeCreate(scope *gorm.Scope) error {
	nodeId := (int64)helper.GetNodesConfig()[0].NodeId
	id, err := helper.GenerateIdBySnowflake(nodeId)
	if err != nil {
		panic("生成ID时发生异常: %s", err)
	}
	scope.Set("ID", &id)
	content.ID = id
	
	// 查询主表id并赋值
	

	return nil
}
