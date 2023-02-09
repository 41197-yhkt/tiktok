package main

import (
	"log"
	"net"

	"github.com/41197-yhkt/tiktok-video/gen/dal"
	video "github.com/41197-yhkt/tiktok-video/kitex_gen/video/douyinservice"
	"github.com/41197-yhkt/tiktok-video/rpc"
	"github.com/41197-yhkt/tiktok-video/service"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	internal_opentracing "github.com/kitex-contrib/tracer-opentracing"

	// 组件的go mod名为pkg
	trace "github.com/41197-yhkt/pkg/trace"
)

func Init() {
	dal.Init()
	service.Init()
	rpc.InitRPC()
}

func main() {
	// 初始化Jaeger并设定opentracing的GlobalTrace
	trace.InitJaeger("kitex-server")

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8890")
	if err != nil {
		panic(err)
	}
	Init()

	svr := video.NewServer(new(DouyinServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "video"}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithSuite(internal_opentracing.NewDefaultServerSuite()),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}
