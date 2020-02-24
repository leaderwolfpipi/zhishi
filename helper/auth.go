package helper

// jwt-auth相关

import (
	"errors"
	"helper"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 载荷添加部分系统需要的信息
type CustomClaims struct {
	UID       string `json:"user_id"`
	Username  string `json:"username"`
	Telephone string `json:"telephone"`
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
	SignKey          string = "82040620FEFAC4511FC65000ADAB0F77"
)

// 新建一个 jwt 实例
func NewJWT() *JWT {
	return &JWT{[]byte(GetSignKey())}
}

// 获取signKey
func GetSignKey() string {
	return SignKey
}

// 设置signKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// token签发
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	// 初始化错误
	var tmpErr error = nil

	// 设置accessToken过期时间为一天
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 设置refreshToken过期时间为一月
	expired := strings.TrimRight(GetTokenConfig().T.ExpiredTime, "dD ")
	expiredInt := strconv.Atoi(expired)
	claims.ExpiresAt = int64(time.Now().Unix() + expiredInt*24*3600)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenRet, err := token.SignedString(j.SignKey)
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
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SignKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}