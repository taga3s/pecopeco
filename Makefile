# コンテナの起動
.PHONY: run
run: run-api run-cli

.PHONY: run-api
run-api:
	docker compose up api db -d

.PHONY: run-cli
run-cli:
	docker compose up cli -d

# コンテナの停止
.PHONY: stop
stop: stop-api stop-cli

.PHONY: stop-api
stop-api:
	docker compose stop api db

.PHONY: stop-cli
stop-cli:
	docker compose stop cli

# コンテナを停止し、 up によって作成されたコンテナ、ネットワークを削除
.PHONY: down
down: down-api down-cli

.PHONY: down-api
down-api:
	docker compose down api db

.PHONY: down-cli
down-cli:
	docker compose down cli

# コンテナに接続する
.PHONY: it-api
it-api:
	docker exec -it api bash

.PHONY: it-cli
it-cli:
	docker exec -it cli bash

.PHONY: it-db
it-db:
	docker exec -it db bash

# ログの出力
.PHONY: logs-api
logs-api:
	docker logs api

.PHONY: logs-cli
logs-cli:
	docker logs cli

.PHONY: logs-db
logs-db:
	docker logs db