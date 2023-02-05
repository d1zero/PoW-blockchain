# Simple Proof-of-Work blockchain

## Basic PoW-blockchain, written on golang

To launch:
- Paste to config/config.yml next content
```yaml
http_server:
  host: 0.0.0.0
  port: 8080

logger:
  level: 0
```

- Run 
```shell
go run cmd/api/main.go
```

How to use:

- Use this request to get all blocks in blockchain
```shell
curl http://127.0.0.1:8080/api/block/ -X GET
```

- Use this request add new block to blockchain (data should be any integer)
```shell
curl http://127.0.0.1:8080/api/block/ -X POST -H "Content-Type: application/json" -d '{"data": 235}'
```