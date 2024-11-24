# =================================================================
# Common
# =================================================================
VERSION := nautilus_v0.0.0


# =================================================================
# for CI
# =================================================================
.PHONY: lint deps build run up down logs
lint:
	golangci-lint run --timeout 3m

deps:
	go mod download && go mod tidy

build: deps
	go build -o ./build/nautilus_api ./cmd/main.go

run: build
	./build/nautilus_api


# =================================================================
# for Container
# =================================================================
.PHONY: up down logs
up:
	docker compose up -d nautilus-api

down:
	docker compose down nautilus-api

logs:
	docker compose logs -f nautilus-api


# =================================================================
# for ECR
# =================================================================
.PHONY: auth push
auth:
	aws ecr get-login-password --region ap-northeast-1 --profile ${AWS_PROFILE} | docker login --username AWS --password-stdin ${AWS_ACCOUNT_ID}.dkr.ecr.ap-northeast-1.amazonaws.com

# make push AWS_ACCOUNT_ID=${AWS_ACCOUNT_ID} AWS_PROFILE=${AWS_PROFILE}
push: auth
	docker build --platform linux/amd64 --no-cache -t $(AWS_ACCOUNT_ID).dkr.ecr.ap-northeast-1.amazonaws.com/dev-nautilus-api:$(VERSION) -f ./build/Dockerfile.nautilus .
	docker push ${AWS_ACCOUNT_ID}.dkr.ecr.ap-northeast-1.amazonaws.com/dev-nautilus-api:$(VERSION)

image-list:
	aws ecr list-images --repository-name dev-nautilus-api --region ap-northeast-1 --profile ${AWS_PROFILE} --query 'imageIds[*]' | jq .

# make delete-images AWS_PROFILE=${AWS_PROFILE}
delete-images:
	@aws ecr list-images --repository-name dev-nautilus-api --region ap-northeast-1 --profile ${AWS_PROFILE} --query 'imageIds[*]' --output json > /tmp/images.json
	@cat /tmp/images.json | jq -c '.[]' | while read -r image_id; do \
	  aws ecr batch-delete-image --repository-name dev-nautilus-api --region ap-northeast-1 --profile ${AWS_PROFILE} --image-ids $${image_id}; \
	done | jq .


# =================================================================
# for Lambda
# =================================================================
.PHONY: deploy
# make deploy AWS_ACCOUNT_ID=${AWS_ACCOUNT_ID} AWS_PROFILE=${AWS_PROFILE}
deploy: push
	aws lambda update-function-code  --profile ${AWS_PROFILE} \
		--function-name dev-nautilus-api \
		--image-uri $(AWS_ACCOUNT_ID).dkr.ecr.ap-northeast-1.amazonaws.com/dev-nautilus-api:$(VERSION) | jq .
