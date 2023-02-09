gen_service:
	hz update -idl ../idl/gateway.thrift

gen_client:
	kitex -module "tiktok-gateway" ../idl/composite.thrift
	kitex -module "tiktok-gateway" ../idl/user.thrift
	kitex -module "tiktok-gateway" ../idl/video.thrift

gen_comp:
	kitex -module "tiktok-gateway" ../idl/composite.thrift

gen_all:
	make gen_service
	make gen_client