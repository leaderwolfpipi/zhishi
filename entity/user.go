package entity

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/leaderwolfpipi/zhishi/helper"
)

// 用户实体表
type User struct {
	// 主键
	ID uint64 `gorm:"type:bigint(20);primary_key;column:user_id" json:"user_id" validate:"required"`
	// 用户名
	// Username string `param:"username"`
	Username string `gorm:"type:varchar(30);unique_index;not null;column:username" json:"username" param:"username" form:"username" validate:"required"`
	// 别名
	Nickname string `gorm:"type:varchar(30);column:nickname" json:"nickname" param:"nickname" form:"nickname"`
	// 密码
	Password string `gorm:"type:varchar(100);column:password;not null" json:"-" param:"password" form:"password" validate:"required"`
	// 邮件
	Email string `gorm:"type:varchar(20);column:email;" json:"email" param:"email" form:"email" validate:"required,email"`
	// 电话
	Telephone string `gorm:"type:varchar(20);unique_index;column:telephone;" json:"telephone" param:"telephone" form:"telephone" validate:"required"`
	// qq号
	QQ string `gorm:"type:varchar(20);column:qq" json:"qq" param:"qq" form:"qq"`
	// 微信号
	Weixin string `gorm:"type:varchar(20);column:weixin" json:"weixin" param:"weixin" form:"weixin"`
	// github账户
	Github string `gorm:"type:varchar(20);column:github" json:"github" param:"github" form:"github"`
	// 是否作者
	IsAuthor uint8 `gorm:"type:tinyint(4);column:is_author;default:1;" json:"is_author" param:"-" form:"is_author"`
	// 用户状态
	Status uint8 `gorm:"type:tinyint(4);column:status;default:1;" json:"status"`
	// 创建时间
	CreateTime int `gorm:"type:int(11);column:create_time;default:0;" json:"create_time" param:"-" form:"create_time"`
	// 更新时间
	LastUpdateTime int `gorm:"type:int(11);column:last_update_time;default:0;" json:"last_update_time" param:"-" form:"last_update_time"`
	// 最后登录时间
	LastLoginTime int `gorm:"type:int(11);column:last_login_time;default:0;" json:"last_login_time" param:"-" form:"last_login_time"`
	// 用户认证token
	Token string `gorm:"type:varchar(100);column:token" json:"token" param:"token" form:"token"`
	// 多对多关联：关注者
	Follows []User `gorm:"many2many:zs_follow;" param:"-"`
	// 多对多关联：被关注者
	Followeds []User `gorm:"many2many:zs_follow;" param:"-"`
}

// 初始函数执行建表和添加外键
func init() {
	// 执行数据迁移
	// helper.Database.AutoMigrate(&User{})

	// 设置外键约束
	// helper.Database.Model(&Star{}).AddForeignKey("article_id", "zs_article(article_id)", "CASCADE", "CASCADE")
}

// 设置表名
func (user *User) TableName() string {
	return "zs_user"
}

// 设置创建钩子
// 插入前生成主键并自动设置插入时间
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	nodeId, _ := strconv.Atoi(helper.GetNodesConfig().Server[0].NodeId)
	id, err := helper.GenerateIdBySnowflake(int64(nodeId))
	if err != nil {
		panic("生成ID时发生异常: %s" + err.Error())
	}
	scope.Set("ID", &id)
	user.ID = id

	// 设置create_time（单位秒）
	scope.SetColumn("CreateTime", time.Now().Unix())

	return nil
}

// 设置更新钩子
// 更新操作时自动更新last_update_time（单位秒）
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	// 设置更新时间
	scope.SetColumn("LastUpdateTime", time.Now().Unix())
	return nil
}

// 获取实体处理函数
func (user *User) GetUserFunc(action string) helper.EntityFunc {
	return func() interface{} {
		var ret interface{}
		if action == "findMore" {
			ret = &[]User{}
		} else {
			ret = &User{}
		}

		return ret
	}
}
