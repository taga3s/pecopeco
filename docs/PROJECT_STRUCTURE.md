```
├── cmd
│   ├── root.go
│   ├── your_command.go
│   └── your_other_command.go
├── core
│   ├── usecase
│   │   ├── user.go
│   │   └── user_test.go
│   ├── repository
│   │   ├── user.go
│   │   └── user_test.go
│   └── entity
│       ├── user.go
│       └── user_test.go
└── main.go
```

1. `cmd`: Cobraのコマンドはcmdディレクトリの中に入るので、usecaseで書いたものを使うだけにする。coreに依存する
2. `core`: 基本実装は全部ここ
   - `usecase`: 複数のリポジトリやエンティティの結合 / 特定のビジネスロジックがここに入る。repositoryとentityに依存する
   - `repository`: 具体的なデータソース（DBや外部API）への依存性はここに含まれる。entityに依存する。
   - `entity`: ビジネスルールを表現するエンティティを配置する。何にも依存しない。
3. `main.go`: アプリケーションのエントリーポイント
