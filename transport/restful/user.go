package restful

import (
	"helper"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/leaderwolfpipi/doris"
	"github.com/leaderwolfpipi/zhishi/entity"
	"github.com/leaderwolfpipi/zhishi/service"
	"github.com/leaderwolfpipi/zhishi/service/server/repository/mysql"
)

// 用户相关接口
func Login(c *doris.Context) error {
	user := entity.User{}
	result := helper.JsonResult{Code: 200, Message: "success", Content: nil}
	// 获取用户信息
	// 参数获取
	username := c.Form("username")
	password := c.Form("password")

	// 实例化repo对象
	repo := mysql.NewRepo(entity.User, *gorm.DB)

	// 传递repo到service层
	service := service.NewService(repo)

	// 调用service的Index接口
	user := service.UserByUsername(username).(*entity.User)
	if user.Password != helper.SHA256(password) {
		// 验证失败
		result.Code = 401
		result.Message = "用户名或密码错误"
		// 返回结果
		c.IndentedJson(200, result)
	} else {
		// 签发token
		err := createToken(c, user)
	}

	return nil
}

// 创建token和refreshToken函数
func createToken(c *doris.Context, user *entity.User) error {
	// 计算生效时间
	activeTime := strconv.Atoi(strings.TrimRight(helper.GetTokenConfig().T.ActiveTime, "sS "))

	// 计算过期时间
	expiredTime := strconv.Atoi(strings.TrimRight(helper.GetTokenConfig().T.ExpiredTime, "sS "))

	// 负载参数
	claims := helper.CustomClaims{
		UID:       user.ID,
		Username:  user.username,
		Telephone: user.telephone,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() + activeTime),       // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + expiredTime*3600), // 过期时间
			Issuer:    helper.GetTokenConfig().Issuer,
		},
	}
	j := helper.NewJWT()
	accessToken, refreshToken, err := j.CreateToken(claims)

	// 出错
	if err != nil {
		c.IndentedJson(http.StatusOK, helper.JsonResult{
			Code:    helper.LoginStatusErr,
			Message: helper.StatusText(helper.LoginStatusErr) + " : " + err.Error(),
		})
		// 退出
		c.Abort()
		return err
	}

	// 成功
	c.IndentedJson(http.StatusOK, helper.JsonResult{
		Code:    helper.LoginStatusOK,
		Message: helper.StatusText(helper.LoginStatusOK),
		Result: doris.D{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
			"user":         user,
		}})

	return nil
}

// 注册接口
func Register(c *doris.Context) error {
	var err error = nil
	// 获取用户信息
	user := entity.User{}
	result := helper.JsonResult{Code: 200, Message: "success"}
	// 获取用户信息
	// 参数绑定
	_ := c.Form(&user)
	username := c.Form("username")

	// 实例化repo对象
	repo := mysql.NewRepo(entity.User, *gorm.DB)

	// 传递repo到service层
	service := service.NewService(repo)

	// 检查用户是否已存在
	user := service.UserByUsername(username).(*entity.User)
	if user != nil {
		// 验证失败
		result.Code = helper.ExistUsernameErr
		result.Message = helper.StatusText(helper.ExistUsernameErr)
		// 返回结果
		c.IndentedJson(200, result)
		c.Abort()
	} else {
		// 执行插入
		err = service.UserAdd(user)
	}

	return err
}
