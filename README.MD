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
| -------- | --------------------- |
| $lt      | Less than             |
| $gt      | Greater than          |
| $in      | In the                |
| $lte     | Less than or equal to |
| $ne      | Not equal to          |

