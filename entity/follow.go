package entity

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/leaderwolfpipi/zhishi/helper"
	// "github.com/shopspring/decimal"
)

// 关注表
type Follow struct {
	// 主键
	ID uint64 `gorm:"type:bigint(20);primary_key;column:follow_id;" json:"follow_id" param:"follow_id" validate:"required"`
	// 评论人id
	UserId uint64 `gorm:"type:bigint(20);column:user_id;" json:"user_id" param:"user_id"`
	// 被关注人id
	FollowedId uint64 `gorm:"type:bigint(20);column:followed_id;" json:"followed_id" param:"followed_id"`
	// 创建时间
	CreateTime int `gorm:"type:int(11);column:create_time;default:0;" json:"create_time" param:"create_time" form:"create_time"`
}

// 初始函数执行建表和添加外键
func init() {
	// 执行数据迁移
	// helper.Database.AutoMigrate(&Follow{})

	// 设置外键约束
	helper.Database.Model(&Follow{}).AddForeignKey("user_id", "zs_user(user_id)", "CASCADE", "CASCADE")
	helper.Database.Model(&Follow{}).AddForeignKey("followed_id", "zs_user(user_id)", "CASCADE", "CASCADE")
}

// 设置表名
func (follow *Follow) TableName() string {
	return "zs_follow"
}

// 设置创建钩子
// 插入前生成主键并自动设置插入时间
func (follow *Follow) BeforeCreate(scope *gorm.Scope) error {
	nodeId, _ := strconv.Atoi(helper.GetNodesConfig().Server[0].NodeId)
	id, err := helper.GenerateIdBySnowflake(int64(nodeId))
	if err != nil {
		panic("生成ID时发生异常: %s" + err.Error())
	}
	scope.Set("ID", &id)
	follow.ID = id

	// 设置create_time（单位秒）
	scope.SetColumn("CreateTime", time.Now().Unix())

	return nil
}

// 获取实体处理函数
func (follow *Follow) GetFollowFunc(action string) helper.EntityFunc {
	return func() interface{} {
		var ret interface{}
		if action == "findMore" {
			ret = &[]Follow{}
		} else {
			ret = &Follow{}
		}

		return ret
	}
}
