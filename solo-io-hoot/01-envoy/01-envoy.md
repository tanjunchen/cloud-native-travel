# intro-envoy-01

## 架构流程

+-------------------+          +----------+           +----------------------------+           +------------------+
|                   |          |          |           |                            |           |                  |
| downstream client +--------->+ listener +---------->+ filters (routing decision) +---------->+ upstream cluster |
|                   |          |          |           |                            |           |                  |
+-------------------+          +----------+           +----------------------------+           +------------------+

route filter 如何转发配置呢？
+-------------+         +------------+          +--------------+           +---------------+          +----------------+
|             |         |            |          |              |           |               |          |                |
| TCP filters +-------->+ HCM filter +--------->+ http filters +---------->+ router filter +--------->+ host selection |
|             |         |            |          |              |           |               |          |                |
+-------------+         +------------+          +--------------+           +---------------+          +----------------+

## demo 测试

```
fuser -k 8082/tcp
fuser -k 10000/tcp
fuser -k 10004/tcp
```

### 测试案例1

```
go run server_tcp.go > /tmp/aa.log 2>&1 &
envoy -c simple_tcp.yaml > /tmp/envoy.log 2>&1 &
```

### 测试案例2

```
go run server.go&
envoy -c simple.yaml&
curl http://localhost:10000 -dhi
```

### cors

```
go run server.go&
envoy -c cors.yaml&
curl -XOPTIONS http://localhost:10000 -H"Origin: solo.io" -v
curl -XOPTIONS http://localhost:10000 -H"Origin: example.com" -v
```

### fault filter

```
go run server.go&
envoy -c simple_fault.yaml -l debug&

for i in $(seq 10); do
curl http://localhost:10000 -s -o /dev/null -w "%{http_code}" 
echo
done
```

### header manipulation

```
go run server.go&
envoy -c response-header.yaml&
curl http://localhost:10000 -v
```
