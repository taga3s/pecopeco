# PECOPECO-CLI_STRUCTURE

```
cli
├── api
│   ├── client
│   │   └── app
│   │       └── http_client.go
│   ├── factory
│   │   └── restaurant
│   │       ├── restaurant.go
│   │       └── params.go
│   ├── model
│   │   └── restaurant.go
│   └── repository
│       └── restaurant
│           ├── repository.go
│           └── request.go
│           └── response.go
├── cmd
│   ├── run.go
│   └── root.go
└── main.go
```

- api/:

1. client: 実際に API 通信を行うための処理を記述する。
2. factory: repository で定義したデータ取得処理をインスタンス化し、アプリケーションで利用できるようにする 。
3. model: アプリケーションでのデータの持ち方を定義する。
4. repository: データ取得等のロジックを定義する。

- cmd: cobra のコマンド及び UI の処理を記述する。
