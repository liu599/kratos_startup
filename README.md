# Bilibili go-kratos 微服务Demo

在kratos new的基础上注册了一个MySQL服务的Demo, 写一下学习路径

数据库用表

```sql
CREATE TABLE `blablabla` (
    	`reg_id` INT(11) NOT NULL AUTO_INCREMENT,
    	`reg_name` CHAR(50) NOT NULL COLLATE 'utf8mb4_bin',
    	`author` CHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_bin',
    	PRIMARY KEY (`reg_name`) USING BTREE,
    	INDEX `id` (`reg_id`) USING BTREE
    )
    COLLATE='utf8mb4_bin'
    ENGINE=InnoDB
    AUTO_INCREMENT=1
    ;
```

自己加个数据吧。
    
## Start up

请自备全局梯子
- clone project 克隆
- protoc.exe需要加入系统的环境变量
- copy db.toml to my-db.toml and change your mysql settings in dsn and readDSN 创建一个my-db.toml, 这个是自己的配置
- cd src/small-service/cmd  
- go build

## 代码结构和最简单开发流程
- model 下加个type
```go
type Article struct {
	RegId int64
	RegName string
	Author string
}
```
- 在internal/dao下
- 1. db.go (相当于model层对数据库操作)
- 2. dao.go (go generate), 比较重要的是按照interface去注册model层， 并按照注释中的bts生成缓存。
```go

    // Dao dao interface
    type Dao interface {
        Close()
        Ping(ctx context.Context) (err error)
        // bts: -nullcache=&model.Article{RegId:-1} -check_null_code=$!=nil&&$.RegId==-1
        Article(c context.Context, regId int64) (*model.Article, error)
        // 这部分之后如果有新的model操作在这里添加
    }
```
- 3. wire.go是一个依赖注入工具。 go generate
- 4. service.go 相当于controller, 这里可以看到使用s.dao可以取得dao的函数（前提是要使用go generate)生成依赖，那么可以知道
protoc这个工具只是帮我们按照约定的交换格式创建API和controller进行连接。

```go

    // Service service.
    type Service struct {
        ac  *paladin.Map
        dao dao.Dao
    }

    
    func (s *Service) RequestItem(ctx context.Context, req *pb.HelloReq) (reply *pb.HelloResp, err error) {
    	d, err2 := s.dao.Article(ctx, 1)
    	if err2!=nil {
    		reply = &pb.HelloResp{
    			Content: "hello " + err2.Error(),
    		}
    		return
    	}
    	reply = &pb.HelloResp{
    		Content: "hello " + d.RegName,
    	}
    	return
    }


```

- 最后到api的api.proto中, 添加路由支持RequestItem

```go

    rpc RequestItem(HelloReq) returns (HelloResp) {
        option (google.api.http) = {
          get: "/small-service/request"
        };
      }


```

- 用工具生成api

```shell
# generate all
kratos tool protoc api.proto
# generate gRPC
kratos tool protoc --grpc api.proto
# generate BM HTTP
kratos tool protoc --bm api.proto
# generate ecode
kratos tool protoc --ecode api.proto
# generate swagger
kratos tool protoc --swagger api.proto
```

- go build 大功告成

```shell script

    cd small-service/cmd
    go build
    ./cmd -conf ../configs

```


后面继续学的话会在这里更新。