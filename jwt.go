package utils

import (
	"Project/global"
	"Project/model/common/request"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.KBB_CONFIG.JWT.Header),
	}
}

type MyClaims struct {
	LibraryId int64 `json:"library_id"`
	UserID    int64 `json:"user_id"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func (j *JWT) GenToken(middleInfo request.MiddleInfo) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		middleInfo.LibraryId.LibraryId,
		middleInfo.UserInfo.Id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(global.KBB_CONFIG.JWT.Timeout) * time.Second * 18).Unix(), // 过期时间
			Issuer:    "my-project",                                                                           // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString([]byte(global.KBB_CONFIG.JWT.Secret))
}

// ParseToken 解析JWT
func (j *JWT) ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(global.KBB_CONFIG.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
