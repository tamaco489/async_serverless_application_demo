# =================================================================
# Common
# =================================================================
VERSION := coral_v0.0.0


# =================================================================
# for CI
# =================================================================
.PHONY: lint deps build run
lint:
	golangci-lint run --timeout 3m

deps:
	go mod download && go mod tidy

build: deps
	go build -o ./build/coral_api ./cmd/main.go

run: build
	./build/coral_api


# =================================================================
# for Container
# =================================================================
.PHONY: up down logs
up:
	docker compose up -d coral-api

down:
	docker compose down coral-api

logs:
	docker compose logs -f coral-api


# =================================================================
# for DynamoDB
# =================================================================
.PHONY: dynamodb-up dynamodb-down dynamodb-logs
dynamodb-up:
	docker compose up -d dynamodb-local dynamodb-admin

dynamodb-down:
	docker compose down dynamodb-local dynamodb-admin

dynamodb-logs:
	docker compose logs -f dynamodb-local

# =================================================================
# dynamodb テーブル作成
# =================================================================
.PHONY: get-dynamodb-tables create-dynamodb-table make delete-dynamodb-table
get-dynamodb-tables:
	aws --endpoint-url http://localhost:8000 dynamodb list-tables | tee

create-dynamodb-table:
	aws --endpoint-url http://localhost:8000 dynamodb create-table \
		--table-name users \
		--attribute-definitions AttributeName=user_id,AttributeType=S \
		--key-schema AttributeName=user_id,KeyType=HASH \
		--provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 | jq .

delete-dynamodb-table:
	aws --endpoint-url http://localhost:8000 dynamodb delete-table \
		--table-name users | jq .


# =================================================================
# for ECR
# =================================================================
.PHONY: auth push get-images delete-images
auth:
	aws ecr get-login-password --region ap-northeast-1 --profile ${AWS_PROFILE} | docker login --username AWS --password-stdin ${AWS_ACCOUNT_ID}.dkr.ecr.ap-northeast-1.amazonaws.com

# make push AWS_ACCOUNT_ID=${AWS_ACCOUNT_ID} AWS_PROFILE=${AWS_PROFILE}
push: auth
	docker build --platform linux/amd64 --no-cache -t $(AWS_ACCOUNT_ID).dkr.ecr.ap-northeast-1.amazonaws.com/dev-coral-api:$(VERSION) -f ./build/Dockerfile.coral .
	docker push ${AWS_ACCOUNT_ID}.dkr.ecr.ap-northeast-1.amazonaws.com/dev-coral-api:$(VERSION)

# make get-images AWS_ACCOUNT_ID=${AWS_ACCOUNT_ID} AWS_PROFILE=${AWS_PROFILE}
get-images:
	aws ecr list-images --repository-name dev-coral-api --region ap-northeast-1 --profile ${AWS_PROFILE} --query 'imageIds[*]' | jq .

# make delete-images AWS_PROFILE=${AWS_PROFILE}
delete-images:
	@aws ecr list-images --repository-name dev-coral-api --region ap-northeast-1 --profile ${AWS_PROFILE} --query 'imageIds[*]' --output json > /tmp/images.json
	@cat /tmp/images.json | jq -c '.[]' | while read -r image_id; do \
	  aws ecr batch-delete-image --repository-name dev-coral-api --region ap-northeast-1 --profile ${AWS_PROFILE} --image-ids $${image_id}; \
	done | jq .


# =================================================================
# for Lambda
# =================================================================
.PHONY: deploy
# make deploy AWS_ACCOUNT_ID=${AWS_ACCOUNT_ID} AWS_PROFILE=${AWS_PROFILE}
deploy: push
	aws lambda update-function-code  --profile ${AWS_PROFILE} \
		--function-name dev-coral-api \
		--image-uri $(AWS_ACCOUNT_ID).dkr.ecr.ap-northeast-1.amazonaws.com/dev-coral-api:$(VERSION) | jq .
