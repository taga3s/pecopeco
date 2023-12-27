# pecopeco

## 技術選定

**言語**

- Go

**フレームワーク**

- Cobra
- Viper

CLI の実装のために、Cobra と Viper を選定する。① 比較的容易に CLI を実装できること、②API 側の実装に時間をかけたいことの 2 点が選定理由である。

**DB**

- PlanetScale

DB に、PlanetScale を選定する。PlanetScale は MySQL のサーバーレスプラットフォームである。① 月に 5GB のストレージ、10 億行の読み取り、1 千万行の書き込みが提供されていること ② デプロイが容易であること ③ 高パフォーマンスは求められないため、データの整合性を重視し RDB を利用したいことの 3 点が選定理由である。

**API**

- REST

**外部 API（仮）**

- ホットペッパー API
- LINE API

**開発環境**

- Docker
