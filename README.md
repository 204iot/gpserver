# gpserver

线上 demo ：http://iot.mogutou.xyz/ui

客户端 HTTP 上报 gps 信息

```
curl --location --request POST 'http://iot.mogutou.xyz/api/v1/device' \
--header 'Content-Type: application/json' \
--data-raw '{
		"id":        "12",
		"longitude": "87.8466796900",
		"latitude":  "43.6440258500"
}'
```

POST json 数据信息

```text
{
    "id":        "12",              # 设备ID
	"longitude": "87.8466796900",   # 经度
	"latitude":  "43.6440258500"    # 纬度
}
```

## 部署

容器启动

```
 docker run -p 8080:90 --name gps xuxu123/gpserver:latest -l 0.0.0.0:90
```

UI: http://127.0.0.1:8080/ui

客户端上报 gps 信息

```
curl --location --request POST '127.0.0.1:8080/api/v1/device' \
--header 'Content-Type: application/json' \
--data-raw '{
		"id":        "12",
		"longitude": "87.8466796900",
		"latitude":  "43.6440258500"
}'
```
