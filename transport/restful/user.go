package restful

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/leaderwolfpipi/doris"
	"github.com/leaderwolfpipi/zhishi/entity"
	"github.com/leaderwolfpipi/zhishi/helper"
	"github.com/leaderwolfpipi/zhishi/service"
	"github.com/leaderwolfpipi/zhishi/service/server/repository/mysql"
)

// 用户相关接口
func Login(c *doris.Context) error {
	// 初始化参数
	var accessToken string  // 用于请求
	var refreshToken string // 用于刷新
	user := &entity.User{}
	result := helper.JsonResult{Code: 200, Message: "success", Result: nil}

	// 参数绑定
	err := c.Form(user)
	if err != nil {
		// 绑定失败
		result.Code = helper.ParamBindErr
		result.Message = helper.StatusText(helper.ParamBindErr) + " [ " + err.Error() + " ]"
		c.IndentedJson(200, result)
		return err
	}

	// 实例化repo对象
	repo := mysql.NewRepo(user.GetUserFunc("findOne"), helper.Database)

	// 传递repo到service层
	service := service.NewService(repo)

	// fmt.Println("user.Username", user.Username)
	// 调用service服务获取用户信息
	tmp, err := service.UserByUsername(user.Username)
	if err != nil && tmp == nil {
		// 失败返回
		result.Code = helper.LoginNotFoundErr
		result.Message = helper.StatusText(helper.LoginNotFoundErr) + " [ " + err.Error() + " ] "
		c.IndentedJson(200, result)
		return err
	} else {
		// 类型转换
		userInfo := tmp.(*entity.User)
		if userInfo.Password != helper.SHA256(user.Password) {
			result.Code = helper.LoginStatusErr
			result.Message = helper.StatusText(helper.LoginStatusErr)
			c.IndentedJson(200, result)
			return errors.New(helper.StatusText(helper.LoginStatusErr))
		}

		// 更新user
		user = userInfo

		// 签发token
		accessToken, refreshToken, err = createToken(c, userInfo)
		if err != nil {
			result.Code = helper.LoginTokenErr
			result.Message = helper.StatusText(helper.LoginTokenErr) + " [ " + err.Error() + " ] "
			c.IndentedJson(http.StatusOK, result)
			c.Abort()
			return err
		}
	}

	// 返回正确
	result.Code = helper.LoginStatusOK
	result.Message = helper.StatusText(helper.LoginStatusOK)
	result.Result = doris.D{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user":          user,
	}
	c.IndentedJson(http.StatusOK, result)
	c.Abort()
	return nil
}

// 创建token和refreshToken函数
func createToken(c *doris.Context, user *entity.User) (string, string, error) {
	// 计算生效时间
	activeTime, _ := strconv.Atoi(strings.TrimRight(helper.GetTokenConfig().Token.ActiveTime, "sS "))

	// 计算过期时间
	expiredTime, _ := strconv.Atoi(strings.TrimRight(helper.GetTokenConfig().Token.ExpiredTime, "dD "))

	// 负载参数
	claims := helper.CustomClaims{
		UID:       user.ID,
		Username:  user.Username,
		Telephone: user.Telephone,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() + int64(activeTime),          // 生效时间
			ExpiresAt: time.Now().Unix() + int64(expiredTime*24*3600), // 过期时间
			Issuer:    helper.GetTokenConfig().Token.Issuer,
		},
	}

	// 计算access_token和refresh_token
	j := helper.NewJWT()
	return j.CreateToken(claims)
}

// 注册接口
func Register(c *doris.Context) error {
	// 获取用户信息
	user := &entity.User{}
	result := helper.JsonResult{
		Code:    helper.UserAddOk,
		Message: helper.StatusText(helper.UserAddOk)}
	// 获取用户信息
	// 参数绑定
	err := c.Form(user)
	if err != nil {
		result.Code = helper.ParamBindErr
		result.Message = helper.StatusText(helper.ParamBindErr)
		c.IndentedJson(200, result)
		c.Abort()
		return err
	}

	// 实例化repo对象
	repo := mysql.NewRepo(user.GetUserFunc("findOne"), helper.Database)

	// 传递repo到service层
	service := service.NewService(repo)

	// 检查用户是否已存在
	tmp, err := service.UserByUsername(user.Username)
	if err != nil && tmp == nil {
		// 用户不存在
		user.Password = helper.SHA256(user.Password)
		err = service.UserAdd(user)

		// 插入失败
		if err != nil {
			result.Code = helper.UserAddErr
			result.Message = helper.StatusText(helper.UserAddErr) + " [ " + err.Error() + " ] "
			// 返回结果
			c.IndentedJson(200, result)
			c.Abort()
			return err
		}

		// 返回成功
		c.IndentedJson(200, result)
		c.Abort()
		return nil

	} else {
		// 用户已存在
		result.Code = helper.ExistUsernameErr
		result.Message = helper.StatusText(helper.ExistUsernameErr)
		// 返回结果
		c.IndentedJson(200, result)
		c.Abort()
		return errors.New(helper.StatusText(helper.ExistUsernameErr))
	}

	return nil
}

// 刷新token
func RefreshToken(c *doris.Context) error {
	// 初始化结果
	result := helper.JsonResult{
		Code:    helper.RefreshTokenOk,
		Message: helper.StatusText(helper.RefreshTokenOk)}

	// 获取刷新token和请求token
	refreshToken := c.FormParam("refresh_token")

	// 重新计算access_token
	j := helper.NewJWT()

	newToken, _, err := j.RefreshToken(refreshToken)
	if err != nil {
		// 反回错误信息
		result.Code = helper.RefreshTokenErr
		result.Message = helper.StatusText(helper.RefreshTokenErr) + " [ " + err.Error() + " ] "
		result.Result = doris.D{
			"access_token":  "",
			"refresh_token": refreshToken,
		}
		c.IndentedJson(http.StatusOK, result)
		c.Abort()
		return nil
	}

	// 返回正确
	result.Code = helper.RefreshTokenOk
	result.Message = helper.StatusText(helper.RefreshTokenOk)
	result.Result = doris.D{
		"access_token":  newToken,
		"refresh_token": refreshToken,
	}
	c.IndentedJson(http.StatusOK, result)
	c.Abort()
	return nil
}

// 重置用户密码
func ResetPWD() {

	// @TODO...

}
