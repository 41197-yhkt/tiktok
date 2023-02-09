package rpc

//import (
	//"tiktok/kitex_gen/composite/compositeservice"
	// "time"

// 	// "pkg/constants"
// 	// "github.com/cloudwego/kitex/client"
// 	// "github.com/cloudwego/kitex/pkg/retry"
// 	// etcd "github.com/kitex-contrib/registry-etcd"
// 	// trace "github.com/kitex-contrib/tracer-opentracing"
// )

// var videoClient compositeservice.Client

// func initVideoRPC(){
	// r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	// if err != nil {
	// 	panic(err)
	// }

	// 
	// c, err := videoservice.NewClient(
	// 	constants.VideoServiceName,
	// 	//client.WithMiddleware(middleware.CommonMiddleware),
	// 	//client.WithInstanceMW(middleware.ClientMiddleware),
	// 	client.WithMuxConnection(1),                       // mux
	// 	client.WithRPCTimeout(3*time.Second),              // rpc timeout
	// 	client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
	// 	client.WithFailureRetry(retry.NewFailurePolicy()), // retry
	// 	client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
	// 	client.WithResolver(r),                            // resolver
	// )
	// if err != nil {
	// 	// panic(err)
	// }
	// videoClient = c
//}