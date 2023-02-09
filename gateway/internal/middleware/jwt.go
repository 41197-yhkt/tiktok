/*
* @Author: 滚~韬
* @Date:   2023/1/28 15:00
 */
package middleware

import (
	"context"
	"net/http"
	"tiktok-gateway/internal/handler"
	"tiktok-gateway/kitex_gen/user"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	identityKey   = "identity"
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key:         []byte("tiktok"), //密钥
		Timeout:     time.Hour * 24,
		MaxRefresh:  time.Hour * 24,
		TokenLookup: "header: token, query: token, form: token",
		// TokenHeadName: "Bearer",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*user.UserLoginResponse); ok {
				return jwt.MapClaims{
					identityKey: v.UserId,
				}
			}
			return jwt.MapClaims{}
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			id, _ := c.Get("user_id")
			c.JSON(http.StatusOK, utils.H{
				"status_code": code,
				"token":       token,
				"user_id ":    id,
				"status_msg":  "success",
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			method, _ := handler.DouyinUserLoginMethod(ctx, c)
			c.Set("user_id", method.(*user.UserLoginResponse).UserId)
			return method, nil
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt err information => %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, statusCode int, statusMsg string) {
			c.JSON(http.StatusOK, utils.H{
				"status_code": statusCode,
				"status_msg":  statusMsg,
			})
		},
	})
	if err != nil {
		panic(err)
	}
}
