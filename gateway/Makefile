gen_service:
	hz update -idl ../idl/gateway.thrift

gen_client:
	kitex -module "github.com/41197-yhkt/tiktok" ../idl/composite.thrift
	kitex -module "github.com/41197-yhkt/tiktok" ../idl/user.thrift
	kitex -module "github.com/41197-yhkt/tiktok" ../idl/video.thrift

gen_comp:
	kitex -module "github.com/41197-yhkt/tiktok" ../idl/composite.thrift

gen_all:
	make gen_service
	make gen_client