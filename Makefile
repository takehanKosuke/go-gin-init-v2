up:
	docker-compose up --build -d

down:
	docker-compose down

# ローカル環境でGoサーバーを起動
local:
	# docker build -t go-gin-init-v2-api ./app
	# docker run -e MYSQL_HOST=host.docker.internal --rm --name go-gin-init-v2-api_1 -p 8080:8080 go-gin-init-v2-api
	go run ./app/cmd/wire_gen.go ./app/cmd/main.go

generate:
	go generate ./...

# テストカバレッジをhtmlで表示する
cover:
	go test -cover ./... -coverprofile=cover.out.tmp
	# 自動生成コードをカバレッジ対象から外し、カバレッジファイルを作成
	cat cover.out.tmp | grep -v "wire_gen.go" > cover.out
	rm cover.out.tmp
	go tool cover -html=cover.out -o cover.html
	open cover.html

lint:
	golangci-lint run -E golint

test:
	go test ./... -v

login_gcp:
	gcloud auth application-default login

cloud_run_apply: login_gcp
	cd ./terraform && terraform init
	cd ./terraform && GOOGLE_APPLICATION_CREDENTIALS=~/.config/gcloud/application_default_credentials.json terraform apply

# docker上のGoサーバーに入る
ssh-api:
	docker exec -it go-gin-init-v2-api_1 /bin/bash

# docker上のMysqlに入る
ssh-mysql:
	docker exec -it go-gin-init-v2_mysql_1 /bin/bash

# migrationファイルの新規作成コマンド
new:
	./scripts/new_migration.sh

migrate-up:
	migrate -source file://app/db/migrations/ -database 'mysql://root:password@tcp(127.0.0.1:3306)/go-gin-init-v2' up

migrate-down:
	migrate -source file://app/db/migrations/ -database 'mysql://root:password@tcp(127.0.0.1:3306)/go-gin-init-v2' down

migrate-ver:
	migrate -source file://app/db/migrations/ -database 'mysql://root:password@tcp(127.0.0.1:3306)/go-gin-init-v2' version

api-generate:
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
		-i local/swagger/index.yml \
		-g go-gin-server \
		-o /local/out/go

dart-cliant-generate:
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
		-i local/swagger/index.yml \
		-g dart \
		-o /local/out/dart

create-template:
	./scripts/create_new_template.sh
