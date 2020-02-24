package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/leaderwolfpipi/zhishi/helper"
	"github.com/shopspring/decimal"
)

// 用户点赞表
type Like struct {
	// 主键
	ID uint64 `gorm:"type:bigint(20);primary_key;column:like_id;" json:"like_id" param:"like_id" validate:"required"`
	// 点赞类型
	LikeType uint8 `gorm:"type:tinyint(4);column:like_type;default:1;" json:"like_type" param:"like_type" form:"like_type"`
	// 评论人id
	UserId uint64 `gorm:"type:bigint(20);column:user_id;" json:"user_id" param:"user_id"`
	// 内容id
	ObjectId uint64 `gorm:"type:bigint(20);column:object_id;" json:"object_id" param:"object_id"`
	// 创建时间
	CreateTime int `gorm:"type:int(11);column:create_time;default:0;" json:"create_time" param:"create_time" form:"create_time"`
}

// 设置表名
func (like *Like) TableName() string {
	return "zs_like"
}

// 设置创建钩子
// 插入前生成主键并自动设置插入时间
func (like *Like) BeforeCreate(scope *gorm.Scope) error {
	nodeId := (int64)helper.GetNodesConfig()[0].NodeId
	id, err := helper.GenerateIdBySnowflake(nodeId)
	if err != nil {
		panic("生成ID时发生异常: %s", err)
	}
	scope.Set("ID", &id)
	like.ID = id

	// 设置create_time（单位秒）
	scope.SetColumn("CreateTime", time.Now().Unix())

	return nil
}
