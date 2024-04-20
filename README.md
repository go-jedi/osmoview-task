## Getting Started

First, run the development server:

```bash
Запуск:
go run cmd/task/main.go
# or
sudo docker compose up
```

```bash
Линтер:
make install-golangci-lint
make lint
```

```bash
Тестирование:
go test ./...
go test -bench=. ./...
make cover
```