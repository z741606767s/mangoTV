## mangoTV - backend
mangoTV - backend

## Requirements
- go1.23.6
- gofiber
- docker
- make
- redis
- mysql
- mongodb
- RabbitMQ

## ***Linux/Mac-Intel***
`docker-compose up -d`


### Event or MQ消息队列事件推送用户
```shell
// 事件推送
_ = svc.ServiceComm.Service.GetEventService().Emit(constants.MqQueueTypeBusiness, utils.JsonMarshal(event.EventMsg{
    EventName: constants.EventTypeNotification,
    Body: string(utils.JsonMarshal(map[string]interface{}{
        "AdminId": 11111,
        "msg":     "ssss",
    })),
}))
```

### Make
使用说明
- 1.编译项目

运行以下命令编译项目，并将生成的可执行文件放入 bin 目录：
```shell
make build
```
- 2.运行项目

编译并运行项目：
```shell
make run
```
- 3.清理生成的文件

清理编译生成的文件：
```shell
make clean
```
- 4.跨平台编译

为 Linux 平台编译：
```shell
make build-linux
```
为 Windows 平台编译：
```shell
make build-windows
```

- 5.查看帮助

显示 Makefile 的使用说明：
```shell
make help
```

- Linux运行
```shell
go run /bin/main.go -config=/Users/yaosen/go/src/mangoTV/app/config/config.yml

nohup /www/mangoTV/mangoTV-linux -config=/www/mangoTV/etc/config.yml > /www/mangoTV/logs/log.log 2>&1 &

ps aux | grep novels-linux

ps aux | grep 'go run main.go'

sudo fuser -k 8089/tcp
```

### Nginx (systemctl | service)
- 配置
```shell
server {
    listen 80;
    server_name www.test.com;

    # 配置根目录
    root /www;

    # 默认页面
    index index.html;

    location / {
        proxy_pass http://localhost:8089; # 将请求代理到应用
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # 配置日志
    access_log  /log/nginx/www.test.com.cn.log;
    error_log  /log/nginx/www.test.com.cn.log;
}
```
- 命令
```shell
sudo service nginx start     //启动     sudo systemctl start nginx
sudo service nginx restart   //重启     sudo systemctl restart nginx
sudo service nginx status    //状态    sudo systemctl status nginx
sudo service nginx stop      //停止     sudo systemctl stop nginx
sudo service nginx reload    


//重载配置  sudo systemctl reload nginx
```

## tree
    项目根目录
    .
    ├── Makefile
    ├── README.md
    ├── app
    │   ├── api
    │   │   └── v1
    │   │       ├── admin
    │   │       │   └── admin.go
    │   │       └── user
    │   │           └── user.go
    │   ├── config
    │   │   ├── base
    │   │   │   └── baseDao.go
    │   │   ├── config.go
    │   │   ├── config.yml
    │   │   ├── constants
    │   │   │   ├── constant.go
    │   │   │   └── errorCode.go
    │   │   ├── mangoTV_gcs.json
    │   │   ├── mongoDB.go
    │   │   ├── mysql.go
    │   │   ├── rabbitMq
    │   │   │   └── rabbitMq.go
    │   │   └── redis.go
    │   ├── cron
    │   │   ├── cron.go
    │   │   ├── fun.go
    │   │   └── job.go
    │   ├── domain
    │   │   ├── logs
    │   │   │   ├── dao
    │   │   │   │   └── logDao.go
    │   │   │   ├── models
    │   │   │   │   └── logs.go
    │   │   │   └── services
    │   │   │       └── logService.go
    │   │   ├── notification
    │   │   │   ├── dao
    │   │   │   │   └── notificationDao.go
    │   │   │   ├── model
    │   │   │   │   ├── customMessage.go
    │   │   │   │   └── notification.go
    │   │   │   └── services
    │   │   │       └── notificationService.go
    │   │   └── users
    │   │       ├── dao
    │   │       │   ├── userDeviceDao.go
    │   │       │   ├── userFlowsDao.go
    │   │       │   └── usersDao.go
    │   │       ├── models
    │   │       │   ├── userDevice.go
    │   │       │   ├── userFlow.go
    │   │       │   └── users.go
    │   │       └── services
    │   │           └── usersService.go
    │   ├── event
    │   │   ├── event.go
    │   │   ├── eventDispatcher.go
    │   │   └── eventHandler.go
    │   ├── main.go
    │   ├── middleware
    │   │   ├── logsMiddleware
    │   │   │   └── logrusLoggerMiddleware.go
    │   │   └── recoverMiddleware
    │   │             └──recoverMiddleware.go
    │   ├── migrations
    │   │   └── migrateTable.go
    │   ├── query
    │   │   └── api
    │   │       └── v1
    │   │           ├── logApiQuery.go
    │   │           ├── notificationApiQuery.go
    │   │           └── userApiQuery.go
    │   ├── routers
    │   │   └── router.go
    │   ├── svc
    │   │   ├── factory
    │   │   │   └── factory.go
    │   │   ├── iface
    │   │   │   ├── event.go
    │   │   │   ├── interface.go
    │   │   │   ├── logService.go
    │   │   │   ├── notificationService.go
    │   │   │   └── usersService.go
    │   │   ├── repository
    │   │   │   └── repository.go
    │   │   └── serviceContext.go
    │   └── utils
    │       ├── base64Captcha
    │       │   └── captcha.go
    │       ├── gosafe.go
    │       ├── helper.go
    │       ├── jwt
    │       │   └── jwt.go
    │       ├── reqLimit
    │       │   └── reqLimit.go
    │       └── responses.go
    ├── build
    │   ├── etc
    │   │   └── config.yml
    │   └── mangoTV-linux
    ├── docker-compose.yml
    ├── go.mod
    ├── go.sum
    ├── logs
    │   ├── 2025-02-14.log
    │   └── 2025-02-15.log
    ├── public
    └── scripts
        └── sql
            └── v1.0-2025-02-15.sql









