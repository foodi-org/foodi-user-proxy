package account

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"github.com/foodi-org/foodi-user-proxy/internal/svc"
	"github.com/foodi-org/foodi-user-proxy/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type (
	JwtLogic struct {
		logx.Logger
		ctx    context.Context
		svcCtx *svc.ServiceContext
	}
)

// get user token
func NewJwtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JwtLogic {
	return &JwtLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (j *JwtLogic) Jwt(req *types.JwtTokenRequest) (resp *types.JwtTokenResponse, err error) {
	var accessExpire = j.svcCtx.Config.Auth.AccessExpire

	now := time.Now().Unix()
	accessToken, err := j.GenToken(now, j.svcCtx.Config.Auth.AccessSecret, nil, accessExpire)
	if err != nil {
		return nil, err
	}

	return &types.JwtTokenResponse{
		Token:        accessToken,
		Expire:       now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (j *JwtLogic) GenToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
