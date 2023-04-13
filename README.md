# gin-server
去中心化合约跟单go的web服务


#### 运行步骤
```shell
修改conf配置文件

cp ./conf/local.yaml ./conf/config.yaml

go mod tidy

go run src/main.go

```

#### 附加命令
```shell
# go代码统一风格格式化， -w 后面指定文件，格式化指定文件；放 . 则表示格式化整个目录
gofmt -l -w .
```

#### 项目结构说明
```text
不要受MVC模式的影响给代码做结构分层

go 的思想就是一个业务都放一个目录，将来做微服务提取也方便，避免前缀相同的引用
```

#### 项目结构
```text

├── README.md
├── conf
│ ├── config.yaml
│ ├── dev.yaml
├── go.mod 依赖模块管理
├── go.sum 依赖模块管理
├── src 
│ ├── enum 枚举
│ │ └── error_enum.go
│ ├── jobs 定时任务
│ │ └── rank_job
│ │     └── rank_job.go
│ ├── listen_block
│ │ └── listen_block_dao.go
│ ├── main.go 启动入口
│ ├── model
│ │ └── chain_listen_block.go
│ ├── user_info 用户信息
│ │ ├── user_account_api.go
│ │ ├── user_info.go
│ │ ├── user_info_dao.go
│ │ └── user_info_service.go
│ └── utils 工具类
│     ├── config
│     │ └── config.go
│     ├── database
│     │ └── db.go
│     ├── errors
│     │ └── errors_handle.go
│     ├── format
│     │ └── api_result.go
│     ├── http
│     ├── localtime
│     │ └── local_time.go
│     ├── mylogs
│     │ └── log.go
│     └── redis
│         └── redis.go
└── test 测试文件目录
    └── test_struct_len.go

```