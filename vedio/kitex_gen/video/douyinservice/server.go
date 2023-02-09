// Code generated by Kitex v0.4.4. DO NOT EDIT.
package douyinservice

import (
	server "github.com/cloudwego/kitex/server"
	video "github.com/41197-yhkt/tiktok-video/kitex_gen/video"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler video.DouyinService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
