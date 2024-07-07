package account

import (
	"context"
	"github.com/foodi-org/foodi-user-proxy/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
	"time"
)

func TestJwtLogic_GenToken(t *testing.T) {
	type fields struct {
		Logger logx.Logger
		ctx    context.Context
		svcCtx *svc.ServiceContext
	}
	type args struct {
		iat       int64
		secretKey string
		payloads  map[string]interface{}
		seconds   int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test generate token",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: &svc.ServiceContext{},
			},
			args: args{
				iat:       time.Now().Unix(),
				secretKey: "fdkeshjfkljkljgeklrbckjbwnkbf",
				payloads:  map[string]interface{}{},
				seconds:   3600,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JwtLogic{
				Logger: tt.fields.Logger,
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
			}
			got, err := j.GenToken(tt.args.iat, tt.args.secretKey, tt.args.payloads, tt.args.seconds)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
