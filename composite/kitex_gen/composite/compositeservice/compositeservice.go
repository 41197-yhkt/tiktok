// Code generated by Kitex v0.4.4. DO NOT EDIT.

package compositeservice

import (
	"context"
	composite "github.com/41197-yhkt/tiktok-composite/kitex_gen/composite"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return compositeServiceServiceInfo
}

var compositeServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CompositeService"
	handlerType := (*composite.CompositeService)(nil)
	methods := map[string]kitex.MethodInfo{
		"BasicFavoriteActionMethod": kitex.NewMethodInfo(basicFavoriteActionMethodHandler, newCompositeServiceBasicFavoriteActionMethodArgs, newCompositeServiceBasicFavoriteActionMethodResult, false),
		"BasicFavoriteListMethod":   kitex.NewMethodInfo(basicFavoriteListMethodHandler, newCompositeServiceBasicFavoriteListMethodArgs, newCompositeServiceBasicFavoriteListMethodResult, false),
		"BasicFeedMethod":           kitex.NewMethodInfo(basicFeedMethodHandler, newCompositeServiceBasicFeedMethodArgs, newCompositeServiceBasicFeedMethodResult, false),
		"BasicCommentActionMethod":  kitex.NewMethodInfo(basicCommentActionMethodHandler, newCompositeServiceBasicCommentActionMethodArgs, newCompositeServiceBasicCommentActionMethodResult, false),
		"BasicCommentListMethod":    kitex.NewMethodInfo(basicCommentListMethodHandler, newCompositeServiceBasicCommentListMethodArgs, newCompositeServiceBasicCommentListMethodResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "composite",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func basicFavoriteActionMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*composite.CompositeServiceBasicFavoriteActionMethodArgs)
	realResult := result.(*composite.CompositeServiceBasicFavoriteActionMethodResult)
	success, err := handler.(composite.CompositeService).BasicFavoriteActionMethod(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCompositeServiceBasicFavoriteActionMethodArgs() interface{} {
	return composite.NewCompositeServiceBasicFavoriteActionMethodArgs()
}

func newCompositeServiceBasicFavoriteActionMethodResult() interface{} {
	return composite.NewCompositeServiceBasicFavoriteActionMethodResult()
}

func basicFavoriteListMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*composite.CompositeServiceBasicFavoriteListMethodArgs)
	realResult := result.(*composite.CompositeServiceBasicFavoriteListMethodResult)
	success, err := handler.(composite.CompositeService).BasicFavoriteListMethod(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCompositeServiceBasicFavoriteListMethodArgs() interface{} {
	return composite.NewCompositeServiceBasicFavoriteListMethodArgs()
}

func newCompositeServiceBasicFavoriteListMethodResult() interface{} {
	return composite.NewCompositeServiceBasicFavoriteListMethodResult()
}

func basicFeedMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*composite.CompositeServiceBasicFeedMethodArgs)
	realResult := result.(*composite.CompositeServiceBasicFeedMethodResult)
	success, err := handler.(composite.CompositeService).BasicFeedMethod(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCompositeServiceBasicFeedMethodArgs() interface{} {
	return composite.NewCompositeServiceBasicFeedMethodArgs()
}

func newCompositeServiceBasicFeedMethodResult() interface{} {
	return composite.NewCompositeServiceBasicFeedMethodResult()
}

func basicCommentActionMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*composite.CompositeServiceBasicCommentActionMethodArgs)
	realResult := result.(*composite.CompositeServiceBasicCommentActionMethodResult)
	success, err := handler.(composite.CompositeService).BasicCommentActionMethod(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCompositeServiceBasicCommentActionMethodArgs() interface{} {
	return composite.NewCompositeServiceBasicCommentActionMethodArgs()
}

func newCompositeServiceBasicCommentActionMethodResult() interface{} {
	return composite.NewCompositeServiceBasicCommentActionMethodResult()
}

func basicCommentListMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*composite.CompositeServiceBasicCommentListMethodArgs)
	realResult := result.(*composite.CompositeServiceBasicCommentListMethodResult)
	success, err := handler.(composite.CompositeService).BasicCommentListMethod(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCompositeServiceBasicCommentListMethodArgs() interface{} {
	return composite.NewCompositeServiceBasicCommentListMethodArgs()
}

func newCompositeServiceBasicCommentListMethodResult() interface{} {
	return composite.NewCompositeServiceBasicCommentListMethodResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) BasicFavoriteActionMethod(ctx context.Context, req *composite.BasicFavoriteActionRequest) (r *composite.BasicFavoriteActionResponse, err error) {
	var _args composite.CompositeServiceBasicFavoriteActionMethodArgs
	_args.Req = req
	var _result composite.CompositeServiceBasicFavoriteActionMethodResult
	if err = p.c.Call(ctx, "BasicFavoriteActionMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BasicFavoriteListMethod(ctx context.Context, req *composite.BasicFavoriteListRequest) (r *composite.BasicFavoriteListResponse, err error) {
	var _args composite.CompositeServiceBasicFavoriteListMethodArgs
	_args.Req = req
	var _result composite.CompositeServiceBasicFavoriteListMethodResult
	if err = p.c.Call(ctx, "BasicFavoriteListMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BasicFeedMethod(ctx context.Context, req *composite.BasicFeedRequest) (r *composite.BasicFeedResponse, err error) {
	var _args composite.CompositeServiceBasicFeedMethodArgs
	_args.Req = req
	var _result composite.CompositeServiceBasicFeedMethodResult
	if err = p.c.Call(ctx, "BasicFeedMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BasicCommentActionMethod(ctx context.Context, req *composite.BasicCommentActionRequest) (r *composite.BasicCommentActionResponse, err error) {
	var _args composite.CompositeServiceBasicCommentActionMethodArgs
	_args.Req = req
	var _result composite.CompositeServiceBasicCommentActionMethodResult
	if err = p.c.Call(ctx, "BasicCommentActionMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BasicCommentListMethod(ctx context.Context, req *composite.BasicCommentListRequest) (r *composite.BasicCommentListResponse, err error) {
	var _args composite.CompositeServiceBasicCommentListMethodArgs
	_args.Req = req
	var _result composite.CompositeServiceBasicCommentListMethodResult
	if err = p.c.Call(ctx, "BasicCommentListMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}