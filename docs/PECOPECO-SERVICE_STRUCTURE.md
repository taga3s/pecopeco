# PECOPECO-SERVICE_STRUCTURE

```
service
├── docker
├── internal
│   ├── config
│   │   └── config.go
│   ├── db
│   │   └── mysql.go
│   ├── domain
│   │   └── user
│   │       ├── user.go
│   │       ├── user_domain_service.go
│   │       ├── user_repository.go
│   │       ├── user_repository_test.go
│   │       └── user_test.go
│   ├── infrastructure
│   │   └── repository
│   │       └── user_repository.go
│   ├── presentation
│   │   ├── responder
│   │   │   └── responder.go
│   │   └──  user
│   │       ├── handler.go
│   │       ├── request.go
│   │       └── response.go
│   ├── server
│   │   ├── middleware
│   │   │   └── auth.go
│   │   ├── route
│   │   │   └── route.go
│   │   └── server.go
│   └──  usecase
│       └── user
│           ├── login_usecase.go
│           └── ...
├── main.go
└── pkg
```

- internal/: アプリケーションの主要な実装は全てここに含まれる。

1.  domain: ビジネスルールを表現するエンティティを配置する。何にも依存しない。
2.  usecase: 複数のリポジトリやエンティティの結合 / ビジネスロジックの実装が含まれる。repository のインターフェースと entity に依存する。
3.  infrastructure: 具体的なデータソース（DB や外部 API）への依存が含まれる。
4.  presentation: クライアントとの入出力の実装が含まれる。
5.  server: エンドポイントのルーティングやミドルウェアの実装が含まれる。

- pkg: ドメインに関係のない便利関数をまとめたものが含まれる。
- main.go: アプリケーションのルート
