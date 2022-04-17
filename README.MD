# brwswg

学习《Building RESTful Web services with Go》的练习代码

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