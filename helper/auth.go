package helper

// jwt-auth相关

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 载荷添加部分系统需要的信息
type CustomClaims struct {
	UID       uint64 `json:"user_id"`
	Username  string `json:"username"`
	Telephone string `json:"telephone"`
	AuthType  string `json:"auth_type"`
	jwt.StandardClaims
}

// jwt签名结构
type JWT struct {
	SignKey []byte
}

// 常量定义
var (
	TokenExpired     error  = errors.New("Token 已经过期")
	TokenNotValidYet error  = errors.New("Token 尚未激活")
	TokenMalformed   error  = errors.New("Token 格式错误")
	TokenInvalid     error  = errors.New("Token 无法解析")
	SignKey          []byte = []byte("82040620FEFAC4511FC65000ADAB0F77")
)

// 新建一个 jwt 实例
func NewJWT() *JWT {
	return &JWT{[]byte(GetSignKey())}
}

// 获取signKey
func GetSignKey() []byte {
	return SignKey
}

// 设置signKey
func SetSignKey(key string) {
	SignKey = []byte(key)
}

// token签发
func (j *JWT) CreateToken(claims CustomClaims) (string, string, error) {
	// 初始化错误
	var tmpErr error = nil

	// 设置accessToken过期时间为一天
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 设置refreshToken过期时间为一月
	expired := strings.TrimRight(GetTokenConfig().Token.RefreshExpiredTime, "dD ")
	expiredInt, _ := strconv.Atoi(expired)
	claims.ExpiresAt = time.Now().Unix() + int64(expiredInt*24*3600)
	claims.AuthType = "refresh" // 标识刷新token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenRet, err := accessToken.SignedString(j.SignKey)
	refreshTokenRet, errRefresh := refreshToken.SignedString(j.SignKey)

	// tmpErr代表出错的那一个
	if err != nil {
		tmpErr = err
	} else {
		tmpErr = errRefresh
	}

	// 返回请求token和刷新token
	return accessTokenRet, refreshTokenRet, tmpErr
}

// token解析
func (j *JWT) ResolveToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SignKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// token刷新
// 注意：使用refreshToken来换取新的accessToken
func (j *JWT) RefreshToken(refreshToken string) (string, string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Now()
	}
	token, err := jwt.ParseWithClaims(refreshToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SignKey, nil
	})

	if err != nil {
		return "", "", err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && claims.AuthType == "refresh" && token.Valid {
		claims.AuthType = "" // 情况refresh标志位
		jwt.TimeFunc = time.Now
		expired := strings.TrimRight(GetTokenConfig().Token.ExpiredTime, "dD ")
		// fmt.Println("expired = ", expired)
		// fmt.Println("cliams = ", claims)
		expiredInt, _ := strconv.Atoi(expired)
		claims.StandardClaims.ExpiresAt = time.Now().Add(time.Duration(expiredInt) * 24 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}

	return "", "", TokenInvalid
}
