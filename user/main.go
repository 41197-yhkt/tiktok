package main

import (
	"github.com/41197-yhkt/pkg/trace"
	user "github.com/41197-yhkt/tiktok-user/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	// Jaeger Init
	closer := trace.InitJaeger("user")
	defer closer.Close()

	// ETCD Init, https://www.cloudwego.io/docs/tutorials/framework-exten/registry/#integrate-into-kitex
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	userServer := user.NewServer(new(UserServiceImpl), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user"}), server.WithRegistry(r), server.WithServiceAddr(addr))
	err = userServer.Run()
	if err != nil {
		log.Fatal(err)
	}
}
