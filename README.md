# uploader

### Requirements:

- Docker
- Go 1.18+
- make
- nc (to run locally integration tests)
- [golangci-lint](https://github.com/golangci/golangci-lint)
- [swag](https://github.com/swaggo/swag)

### Commands:

- Start uploader `make start`
- Stop uploader `make stop`
- Rebuild uploader container `make rebuild`
- Run uploader on local machine `make start-locally`
- Run unit tests `make unit-test`
- Run integration tests `make integration-test`
- Run linter `make linter`

### Configuration

Default config is located in `.env_default`:
```dotenv
HOST=127.0.0.1
PORT=3333
UPLOAD_PATH=./uploads
CACHE_PATH=./cache
ALLOWED_MIME_TYPES=image/jpeg,image/png
ALLOWED_WIDTHS=1024,800,600,200
CORS_ORIGINS=*
JWT_SECRET=my_secret
BODY_LIMIT=1M
```

### Start server

1. Copy `.env_default` to `.env` and then:
2. Create 2 folders: `mkdir uploads` and `mkdir cache`
3. `make start` using Docker or `make start-locally` without Docker

### Predefined JWT token

For testing can be used:

User 1 JWT: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxIn0.TA1WENwC5LZ3dqlnfXKVziVI30uG8-7QSboAP9xoyzg`

User 2 JWT: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIyIn0.sC-bPVfoCozp-QFRRQ9S5MQu5493FG4G2x_aWFwM_GQ`

### Example requests

Upload test receipt:
```shell
 curl -H 'Accept: application/json' -X POST -Ffile=@./tests/testdata/receipt2.jpg -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxIn0.TA1WENwC5LZ3dqlnfXKVziVI30uG8-7QSboAP9xoyzg" http://127.0.0.1:3333/v1/receipts
```

Get list of receipts:
```shell
 curl -H 'Accept: application/json' -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxIn0.TA1WENwC5LZ3dqlnfXKVziVI30uG8-7QSboAP9xoyzg" http://127.0.0.1:3333/v1/receipts
```

Download receipt:
```shell
 curl -H 'Accept: application/json' -X GET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxIn0.TA1WENwC5LZ3dqlnfXKVziVI30uG8-7QSboAP9xoyzg" http://127.0.0.1:3333/v1/receipts/2022/12/15/1671097462498380000.jpg
```

Download receipt with resizing:
```shell
 curl -H 'Accept: application/json' -X GET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxIn0.TA1WENwC5LZ3dqlnfXKVziVI30uG8-7QSboAP9xoyzg" http://127.0.0.1:3333/v1/receipts/2022/12/15/1671097462498380000.jpg?width=200 --output test.jpg
```

List all receipts:
```shell
 curl -H 'Accept: application/json' -X GET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxIn0.TA1WENwC5LZ3dqlnfXKVziVI30uG8-7QSboAP9xoyzg" http://127.0.0.1:3333/v1/receipts
```

OCR receipt:
```shell
 curl -H 'Accept: application/json' -X GET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxIn0.TA1WENwC5LZ3dqlnfXKVziVI30uG8-7QSboAP9xoyzg" http://127.0.0.1:3333/v1/receipts/ocr/2022/12/15/1671109072229953000.jpg
```
