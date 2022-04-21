# brwswg

学习《Building RESTful Web services with Go》的练习代码

## project init

```shell

mkdir brwswg

cd brwswg

go mod init github.com/guozhe001/brwswg

go get go.mongodb.org/mongo-driver/mongo

# 删除未使用的依赖项
go mod tidy

```

## chapter3

### 中间件

1. [gorilla/handlers](https://github.com/gorilla/handlers)
2. [justinas/alice](https://github.com/justinas/alice)
3. [urfave/negroni](https://github.com/urfave/negroni)
3. [satori/go.uuid](http://github.com/satori/go.uuid)

### RPC JSON request

```shell
curl -X POST \
http://localhost:1234/rpc \
-H 'cache-control: no-cache' \
-H 'content-type: application/json' \
-d '{
"method": "JSONServer.GiveBookDetail",
"params": [{
"Id": "1234"
}],
"id": "1"
}
```

or :

```shell
curl -X POST http://localhost:1234/rpc -H 'cache-control: no-cache' -H 'content-type: application/json' -d '{"method": "JSONServer.GiveBookDetail","params": [{"Id": "1234"}],"id": "1"}
```

## chapter4

### Popular Go Frameworks

1. [go-restful](https://github.com/emicklei/go-restful)
2. [go-restful-swagger12](https://github.com/emicklei/go-restful-swagger12)
3. [gin](https://github.com/gin-gonic/gin)
3. [revel](https://github.com/revel/revel)

## chapter5

最终使用[mongo-go-driver](https://github.com/mongodb/mongo-go-driver) 替代书中的[mgo](https://gopkg.in/mgo.v2)

### mongoDB命令记录

```
use movies;

db.movies.insertOne({ _id: 5, name: 'Star Trek', year: 2009, directors: ['J.J. Abrams'], writers: ['Roberto Orci', 'Alex Kurtzman'], boxOffice: { budget:150000000, gross:257704099 } } )

db.movies.insertOne({ name: 'The Dark Knight ', year: 2008, directors: ['Christopher Nolan'], writers: ['Jonathan Nolan', 'Christopher Nolan'], boxOffice: { budget:185000000, gross:533316061 } } )

db.movies.findOne();

db.movies.find();

db.movies.find({year : {$eq : 2008}})

db.movies.find({year: 2008})

db.movies.find({'boxOffice.budget': {$gt : 150000000}})

db.movies.find({'boxOffice.budget': {$gt : 150000000}, year:2009})

db.movies.find({$or: [{'boxOffice.budget': {$gt : 150000000}}, {year:2009}]})

db.movies.deleteOne({"_id" : ObjectId('625c2ccadb210749da806846')})

# 创建索引
db.movies.createIndex({year : 1})

```

#### 其他操作关键字

| Operator | Function              |
|----------|-----------------------|
| $lt      | Less than             |
| $gt      | Greater than          |
| $in      | In the                |
| $lte     | Less than or equal to |
| $ne      | Not equal to          |

## chapter5

#### Go内置数据类型与Protobuf类型不一样的类型对应关系表：

| Go type | Protobuf type |
|---------|---------------|
| float32 | float         |
| float64 | double        |
| uint32  | fixed32       |
| uint64  | fixed64       |
| []byte  | bytes         |


#### Protobuf类型的默认值

| Protobuf type                 | Default value |
|-------------------------------|---------------|
| string                        | ""            |
| bytes                         | empty bytes[] |
| bool                          | false         |
| int, int32,int64,float,double | 0             |
| enum                          | 0             |



#### Go程序中如何使用Protobuf

![image-20220421223358457](https://cdn.jsdelivr.net/gh/guozhe001/oss//imgimage-20220421223358457.png)


#### 错误记录:"protoc-gen-go: program not found or is not executable"

```shell
$  protoc --go_out=. *.proto                                   

protoc-gen-go: program not found or is not executable
Please specify a program using absolute path or make sure the program is available in your PATH system variable
--go_out: protoc-gen-go: Plugin failed with status code 1.

```

**解决方案**
1. 配置GO环境变量如下：

   ```shell
   export GOROOT=/usr/local/go
   export GOPATH=$HOME/go
   export GOBIN=$GOPATH/bin
   export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
   ```
   
2. 安装`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`包

参考文档：https://developers.google.com/protocol-buffers/docs/gotutorial



#### 错误记录：unable to determine Go import path for "person.proto"

```shell
$ protoc --go_out=. person.proto                                
protoc-gen-go: unable to determine Go import path for "person.proto"

Please specify either:
        • a "go_package" option in the .proto source file, or
        • a "M" argument on the command line.

See https://developers.google.com/protocol-buffers/docs/reference/go-generated#package for more information.

--go_out: protoc-gen-go: Plugin failed with status code 1.

```
**解决方案**
在person.proto中执行go_package:
```protobuf
option go_package = "./";
```

参考文档：https://stackoverflow.com/questions/70586511/protoc-gen-go-unable-to-determine-go-import-path-for-simple-proto

#### protocol buffers和JSON比较有什么好处？
1. protocol buffers编译成二进制文本小于JSON文本