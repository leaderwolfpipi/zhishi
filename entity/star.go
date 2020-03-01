package entity

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/leaderwolfpipi/zhishi/helper"
	// "github.com/shopspring/decimal"
)

// 用户收藏表
type Star struct {
	// 主键
	ID uint64 `gorm:"type:bigint(20);primary_key;column:star_id;" json:"star_id" param:"star_id"`
	// 评论人id
	UserId uint64 `gorm:"type:bigint(20);column:user_id;" json:"user_id" param:"user_id" validate:"required"`
	// 内容id
	ArticleId uint64 `gorm:"type:bigint(20);column:article_id;" json:"article_id" param:"article_id" validate:"required"`
	// 作者id
	AuthorId uint64 `gorm:"type:bigint(20);column:author_id;" json:"author_id" param:"author_id"`
	// 创建时间
	CreateTime int `gorm:"type:int(11);column:create_time;default:0;" json:"create_time" param:"create_time" form:"create_time"`
}

// 设置表名
func (star *Star) TableName() string {
	return "zs_star"
}

// 设置创建钩子
// 插入前生成主键并自动设置插入时间
func (star *Star) BeforeCreate(scope *gorm.Scope) error {
	nodeId, _ := strconv.Atoi(helper.GetNodesConfig().Server[0].NodeId)
	id, err := helper.GenerateIdBySnowflake(int64(nodeId))
	if err != nil {
		panic("生成ID时发生异常: %s" + err.Error())
	}
	scope.Set("ID", &id)
	star.ID = id

	// 设置create_time（单位秒）
	scope.SetColumn("CreateTime", time.Now().Unix())

	return nil
}

// 获取实体处理函数
func (star *Star) GetStarFunc(action string) helper.EntityFunc {
	return func() interface{} {
		var ret interface{}
		if action == "findMore" {
			ret = make([]Star, 0)
		} else {
			ret = new(Star)
		}

		return ret
	}
}
