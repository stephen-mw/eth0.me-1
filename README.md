# eth0.me

server side:
```bash
$ go run main.go
2016/06/24 22:58:35 Runing eth0.me server on port [HTTPS]: 443
2016/06/24 22:58:35 Runing eth0.me server on port [HTTP]: 80
2016/06/24 22:58:40 0:443: GET / request from 127.0.0.1.
```

client side:
```bash
$ curl https://0  -k
127.0.0.1
```
