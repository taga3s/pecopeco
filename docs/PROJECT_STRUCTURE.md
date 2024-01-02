```
├── cmd
│ ├── root.go
│ ├── your_command.go
│ └── your_other_command.go
├── internal
│ ├── usecase
│ │ ├── user.go
│ │ └── user_test.go
│ ├── entity
│ │ ├── user.go
│ │ └── user_test.go
│ └── repository
│ │ ├── user.go
│ │ ├── user_test.go
│ │ └── interfaces
│ │     ├── user_repository.go
│ │     └── user_repository_test.go
│ └── presentation
│   ├── handler.go
│   └── handler_test.go
└── main.go
```

1. cmd: Cobraのコマンドはcmdディレクトリの中に入るので、usecaseで書いたものを使うだけにする。internalパッケージに依存する。
2. internal: アプリケーションの主要な実装は全てここに含まれる。
   - usecase: 複数のリポジトリやエンティティの結合 / 特定のビジネスロジックがここに入る。repositoryのインターフェースとentityに依存する。
   - repository: 具体的なデータソース（DBや外部API）への依存性はここに含まれる。
   - entity: ビジネスルールを表現するエンティティを配置する。何にも依存しない。
3. main.go: アプリケーションのエントリーポイント。ここで依存性の注入が行われ、具体的なrepositoryの実装がusecaseに提供される。
4. presentation: 与えられたコマンドに応じて適切なusecaseを呼び出す。usecaseに依存する。
