<p align="center">
  <img src='assets/pecopeco.png' width="100%"/>
</p>

## About this app

エンジニアとして活動している人たちが、何かご飯を食べたいと思った時に、ブラウザを開かずともすぐさま飲食店を検索できるCLIアプリです。現在、以下のような機能があります。

- 飲食店検索機能
- 飲食店 LINE 通知機能
- 飲食店シェア掲示板機能

## Docs

- API 定義はこちら
  - [pecopeco-api-docs](https://taga3s.github.io/pecopeco/)

## Tech Stacks

### Language

- [Go](https://go.dev/)

### Frameworks

- [Cobra](https://cobra.dev/)
- [Viper](https://github.com/spf13/viper)

### DB

- MySQL

### API

- REST API

### External APIs

- [ホットペッパー API](https://webservice.recruit.co.jp/doc/hotpepper/reference.html)
- [LINE Notify API](https://notify-bot.line.me/doc/ja/)

1. ホットペッパー API は、多くの飲食店情報を提供する。
2. LINE Notify API を利用することで、ユーザーがトークンを発行し、任意のグループ等と連携を行うことで、LINE に通知することができる。

### Dev Environment

- [Docker](https://docs.docker.com/)

## System Structure

<img src="docs/system-structure.png" width="100%"/>

## Usage

- run app
```
$ pecopeco run
```

- configure `LINE_NOTIFY_API_TOKEN`
```
$ pecopeco config --token <LINE_NOTIFY_API_TOKEN>
```

## Setup

### 共通

1.  `/cli`と`/service`にある`.env.example`をそれぞれコピーして、`.env`として配置します。

```sh
cp .env.example .env
```

2. cli, api, db のコンテナを全て起動します。

```sh
make run
```

### cli, api コンテナに接続し、Go を実行する

1. コンテナに接続します。

```sh
make it-cli
```

or

```sh
make it-api
```

2. main.go を実行します。

```sh
go run main.go
```

### db コンテナに接続し、mysql に root ユーザとしてログインする

1. コンテナに接続します。

```sh
make it-db
```

2. mysql に root ユーザとしてログインします。

```sh
mysql -uroot -ppeco_password
```
