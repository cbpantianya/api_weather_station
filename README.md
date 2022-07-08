# api_weather_station

天气收集平台后端部分源码

---
参考代码:[微信公众号后端的快速开发框架](https://github.com/hduhelp/wechat-template)

---
[嵌入式部分博客]()

[前端源码]()

## 配置文件说明

请进入```config/config.toml.ep```文件夹查看配置文件说明
```toml
[httpEngine]
host = 'localhost'
port = 9956

[mysql]
host = 'localhost'
port = 3306
user = 'root'
passwd = '123456'
db = 'api_weather_station'

[redis]
host = 'localhost'
port = 6379
db = 0
passwd = '123456'

[log]
env = 'dev' # 生产还是测试环境
path = './logs' # 日志文件存放路径
```

## 如何运行

1. 初始化Mysql数据库表
```sql
create table records
(
    id          int auto_increment
        primary key,
    temperature float    null,
    humidity    float    null,
    time        bigint   null,
    sn          longtext null,
    constraint records_id_uindex
        unique (id)
);

create table rw_test
(
    update_time time null,
    test_id     int  null
);
```

2. 填写配置文件```config/config.toml.ep```并重命名为```config.toml```
3. 启动程序

```bash
$ cd api_weather_station
$ go run main.go
```