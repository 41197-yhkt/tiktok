# tiktok-gateway
Gateway for simple-tiktok, using Hertz, including authentication, rooter, serve HTTP request

## Project Layout
```
tiktok-gateway
├── configs
├── global
├── internal
│   ├── handler
│   ├── idl
│   ├── middleware
│   ├── routers
│   └── rpc
├── pkg
│   ├── errno
│   ├── logger
│   └── tracer
└── script
```

* configs：配置文件。
* global：全局变量。
* internal：内部模块。
    * handler：项目核心业务逻辑。
    * idl：idl文件。
    * rpc：封装RPC调用函数。
    * middleware：HTTP 中间件。
    * routers：路由相关逻辑处理。
* pkg：项目相关的模块包。
* scripts：各类构建，安装，分析等操作的脚本。