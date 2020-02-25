package restful

import (
	"errors"
	"fmt"
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
	// 存储错误
	var errs []error
	errs = make([]error, 0)

	// 初始化参数
	user := &entity.User{}
	result := helper.JsonResult{Code: 200, Message: "success", Result: nil}

	// debug
	fmt.Println("测试1....")

	// 参数绑定
	_ = c.Form(user)

	fmt.Println("测试2....")

	fmt.Println(user.GetUserFunc("findOne"))

	// 实例化repo对象
	repo := mysql.NewRepo(user.GetUserFunc("findOne"), helper.Database)

	fmt.Println("测试2.5....")

	// 传递repo到service层
	service := service.NewService(repo)

	fmt.Println("测试2.8....")

	fmt.Println(user.Username)

	fmt.Println("测试2.9....")

	// 调用service服务获取用户信息
	tmp, err := service.UserByUsername(user.Username)

	fmt.Println("测试3....")

	userInfo := tmp.(*entity.User)

	fmt.Println("测试4....")

	if err != nil || userInfo.Password != helper.SHA256(user.Password) {
		// 加入错误
		errs = append(errs, err)

		// 验证失败
		result.Code = 401
		result.Message = "用户名或密码错误"

		// 返回结果
		c.IndentedJson(200, result)
	} else {
		// 签发token
		err := createToken(c, userInfo)

		if err != nil {
			// 加入错误
			errs = append(errs, err)
		}
	}

	fmt.Println("测试5....")

	return getErrs(errs)
}

// 创建token和refreshToken函数
func createToken(c *doris.Context, user *entity.User) error {
	// 计算生效时间
	activeTime, _ := strconv.Atoi(strings.TrimRight(helper.GetTokenConfig().T.ActiveTime, "sS "))

	// 计算过期时间
	expiredTime, _ := strconv.Atoi(strings.TrimRight(helper.GetTokenConfig().T.ExpiredTime, "sS "))

	// 负载参数
	claims := helper.CustomClaims{
		UID:       user.ID,
		Username:  user.Username,
		Telephone: user.Telephone,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() + int64(activeTime),       // 生效时间
			ExpiresAt: time.Now().Unix() + int64(expiredTime*3600), // 过期时间
			Issuer:    helper.GetTokenConfig().T.Issuer,
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
	user := &entity.User{}
	result := helper.JsonResult{Code: 200, Message: "success"}
	// 获取用户信息
	// 参数绑定
	_ = c.Form(user)

	// 实例化repo对象
	repo := mysql.NewRepo(user.GetUserFunc("findOne"), helper.Database)

	// 传递repo到service层
	service := service.NewService(repo)

	// 检查用户是否已存在
	tmp, err := service.UserByUsername(user.Username)
	userInfo := tmp.(*entity.User)
	if userInfo != nil {
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

// 重组错误
func getErrs(errs []error) error {
	errStr := ""
	for _, err := range errs {
		errStr += err.Error()
	}

	return errors.New(errStr)
}
