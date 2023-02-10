# IDL
The IDL project contains all services' idls in the project

## 研发流程推荐
1. 如果需要新建对应的idl文件，或者往现有的idl文件新增内容，请在idl目录中进行修改。为了方便更新代码，可以在本地编写Makefile，以后一键就能通过idl生成对应的代码，何乐而不为？
2. 并且需要用到对应的客户端和服务端，请仔细阅读对应的代码生成工具指令（例如hz & kitex），并将对应的代码生成更新命令添加在Makefile中的相应位置，下面给了一个Makefile的小例子
   - 如果需要生成服务端代码，请将对应逻辑，添加在gen_service位置
   - 如果需要生成客户端代码，请将对应逻辑，添加在gen_client位置
```makefile
gen_service:
	hz new -mod yhkt/gateway -idl ../idl/gateway.thrift
	kitex -module "KitexTest" -service KitexTest ../idl/gateway.thrift

gen_client:
	kitex -module "yhkt/social" ../idl/social.thrift
	kitex -module "yhkt/interact" ../idl/interact.thrift
	kitex -module "yhkt/basic" ../idl/basic.thrift

gen_all:
	make gen_service
	make gen_client
```

3. 如果在已有thrift文件中，进行了数据结构的修改或者接口的更新，请在本地拉取最新的idl分支，重新执行gen_client / gen_service，并更新业务代码中用到的数据结构昂！
