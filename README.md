# gpserver

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

容器启动

```
 docker run -p 8080:90 --name gps xuxu123/gpserver:latest -l 0.0.0.0:90
```
