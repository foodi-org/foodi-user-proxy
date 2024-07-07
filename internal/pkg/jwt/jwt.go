package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type (
	TokenOptions struct {
		AccessSecret string
		AccessExpire int64
		Fields       map[string]interface{}
	}

	Token struct {
		AccessToken  string `json:"accessToken"`
		AccessExpire int64  `json:"accessExpire"`
	}

	Claims struct {
		Mobile string
		jwt.MapClaims
	}
)

func New(options TokenOptions) (Token, error) {
	var token Token
	now := time.Now().Add(-time.Minute).Unix()
	accessToken, err := getJwtToken(options.AccessSecret, now, options.AccessExpire, options.Fields)
	if err != nil {
		return token, err
	}
	token.AccessToken = accessToken
	token.AccessExpire = options.AccessExpire

	return token, nil
}

// getJwtToken
//
//	@Description: 生成token
//	@param secret jwt秘钥
//	@param iat 时间戳
//	@param second 过期时间，单位秒
//	@param payload 数据载体
//	@return string token
//	@return error
func getJwtToken(secret string, iat, second int64, payload map[string]interface{}) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + second
	claims["iat"] = iat
	for k, v := range payload {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secret))
}

// ParseToken
//
//	@Description: 解析token
//	@param tokenStr: 加密token
//	@param t: 秘钥
//	@return *Claims: 解密后的 Claims 对象
//	@return error
func ParseToken(tokenStr string, t string) (*Claims, error) {
	mc := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, mc, func(token *jwt.Token) (interface{}, error) {
		return []byte(t), nil
	})
	if err != nil {
		return nil, errors.New("JWT已过期")
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("token校验失败")
}
