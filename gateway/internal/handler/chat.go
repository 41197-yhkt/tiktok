/*
* @Author: 滚~韬
* @Date:   2023/2/3 23:09
 */
package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"net/http"
)

func ChatMethod(ctx context.Context, c *app.RequestContext) {
	//cli, err := client.NewClient()
	//if err != nil {
	//	panic(err)
	//}
	//r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	//if err != nil {
	//	panic(err)
	//}
	//cli.Use(sd.Discovery(r))
	//status, body, err := cli.Get(context.Background(), nil, "http://chatsserver/ws", config.WithSD(true))
	//
	//if err != nil {
	//	hlog.Fatal(err)
	//}
	//hlog.Infof("HERTZ: code=%d,body=%s", status, string(body))
	c.Redirect(http.StatusMovedPermanently, []byte("http://127.0.0.1:8001/ws"))
	c.JSON(consts.StatusOK, "OK")
}
