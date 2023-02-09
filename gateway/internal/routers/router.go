// Code generated by hertz generator.

package routers

import (
	"tiktok-gateway/internal/handler"
	"tiktok-gateway/internal/middleware"

	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	group := r.Group("/douyin")
	group.POST("/user/register", handler.DouyinUserRegisterMethod, middleware.JwtMiddleware.LoginHandler)
	group.POST("/user/login", middleware.JwtMiddleware.LoginHandler)
	auth := group.Group("/user", middleware.JwtMiddleware.MiddlewareFunc())
	auth.GET("/", handler.DouyinUserMethod)
	auth = group.Group("/relation", middleware.JwtMiddleware.MiddlewareFunc())
	auth.POST("/action", handler.DouyinUserMethod)
	auth.GET("/follow/list", handler.DouyinUserMethod)
	auth.GET("/follower/list", handler.DouyinUserMethod)
	auth.GET("/friend/list", handler.DouyinUserMethod)

	// feed不需jwt
	group.GET("/feed", handler.DouyinFeedMethod)

	auth2 := group.Group("/publish", middleware.JwtMiddleware.MiddlewareFunc())
	auth3 := group.Group("/favorite", middleware.JwtMiddleware.MiddlewareFunc())
	auth4 := group.Group("/comment", middleware.JwtMiddleware.MiddlewareFunc())

	auth2.POST("/action", handler.DouyinPublishActionMethod)
	auth2.GET("/list", handler.DouyinPublishListMethod)
	auth3.POST("/action", handler.DouyinFavoriteActionMethod)
	auth3.GET("/list", handler.DouyinFavoriteListMethod)
	auth4.POST("/action", handler.DouyinCommentActionMethod)
	auth4.GET("/list", handler.DouyinCommentListMethod)

}