package entity

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/leaderwolfpipi/zhishi/helper"
	// "github.com/shopspring/decimal"
)

// 用户评论表
type Comment struct {
	// 主键
	ID uint64 `gorm:"type:bigint(20);primary_key;column:comment_id;" json:"comment_id" param:"comment_id" validate:"required"`
	// 评论内容
	Comment string `gorm:"type:varchar(300);column:comment;" json:"comment" param:"comment" form:"comment"`
	// 父评论id
	PID uint64 `gorm:"type:bigint(20);column:pid;" json:"pid" param:"pid"`
	// 评论人id
	UserId uint64 `gorm:"type:bigint(20);column:user_id;" json:"user_id" param:"user_id"`
	// 文章id
	ArticleId uint64 `gorm:"type:bigint(20);column:article_id;" json:"article_id" param:"article_id"`
	// 文章状态
	Status uint8 `gorm:"type:tinyint(4);column:status;default:1;" json:"status" param:"status" form:"status"`
	// 创建时间
	CreateTime int `gorm:"type:int(11);column:create_time;default:0;" json:"create_time"`
	// 更新时间
	LastUpdateTime int `gorm:"type:int(11);column:last_update_time;default:0;" json:"last_update_time"`
	// 关联like表
	Likes []Like `gorm:"ForeignKey:ObjectId;association_foreignkey:ID" json:"-"`
}

// 初始函数执行建表和添加外键
func init() {
	// 执行数据迁移
	// helper.Database.AutoMigrate(&Comment{})

	// 设置外键约束
	helper.Database.Model(&Comment{}).AddForeignKey("article_id", "zs_article(article_id)", "CASCADE", "CASCADE")
	helper.Database.Model(&Comment{}).AddForeignKey("user_id", "zs_user(user_id)", "CASCADE", "CASCADE")
}

// 设置表名
func (comment *Comment) TableName() string {
	return "zs_comment"
}

// 设置创建钩子
// 插入前生成主键并自动设置插入时间
func (comment *Comment) BeforeCreate(scope *gorm.Scope) error {
	nodeId, _ := strconv.Atoi(helper.GetNodesConfig().Server[0].NodeId)
	id, err := helper.GenerateIdBySnowflake(int64(nodeId))
	if err != nil {
		panic("生成ID时发生异常: %s" + err.Error())
	}
	scope.Set("ID", &id)
	comment.ID = id

	// 设置create_time（单位秒）
	scope.SetColumn("CreateTime", time.Now().Unix())

	return nil
}

// 设置更新钩子
// 更新操作时自动更新last_update_time（单位秒）
func (comment *Comment) BeforeUpdate(scope *gorm.Scope) error {
	// 设置更新时间
	scope.SetColumn("LastUpdateTime", time.Now().Unix())
	return nil
}

// 获取实体处理函数
func (comment *Comment) GetCommentFunc(action string) helper.EntityFunc {
	return func() interface{} {
		var ret interface{}
		if action == "findMore" {
			ret = &[]Comment{}
		} else {
			ret = &Comment{}
		}

		return ret
	}
}
