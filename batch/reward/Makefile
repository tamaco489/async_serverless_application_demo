# =================================================================
# Common
# =================================================================
VERSION := reward_v0.0.0


# =================================================================
# 報酬管理バッチ
# =================================================================
.PHONY: build run
build:
	docker build -f build/Dockerfile.reward -t reward-batch .

run:
	docker run --rm -it reward-batch:latest


# =================================================================
# for ECR
# =================================================================
.PHONY: auth push
auth:
	aws ecr get-login-password --region ap-northeast-1 --profile ${AWS_PROFILE} | docker login --username AWS --password-stdin ${AWS_ACCOUNT_ID}.dkr.ecr.ap-northeast-1.amazonaws.com

# make push AWS_ACCOUNT_ID=${AWS_ACCOUNT_ID} AWS_PROFILE=${AWS_PROFILE}
push: auth
	docker build --platform linux/amd64 --no-cache -t $(AWS_ACCOUNT_ID).dkr.ecr.ap-northeast-1.amazonaws.com/dev-reward-batch:$(VERSION) -f ./build/Dockerfile.reward .
	docker push ${AWS_ACCOUNT_ID}.dkr.ecr.ap-northeast-1.amazonaws.com/dev-reward-batch:$(VERSION)

image-list:
	aws ecr list-images --repository-name dev-reward-batch --region ap-northeast-1 --profile ${AWS_PROFILE} --query 'imageIds[*]' | jq .

# make delete-images AWS_PROFILE=${AWS_PROFILE}
delete-images:
	@aws ecr list-images --repository-name dev-reward-batch --region ap-northeast-1 --profile ${AWS_PROFILE} --query 'imageIds[*]' --output json > /tmp/images.json
	@cat /tmp/images.json | jq -c '.[]' | while read -r image_id; do \
	  aws ecr batch-delete-image --repository-name dev-reward-batch --region ap-northeast-1 --profile ${AWS_PROFILE} --image-ids $${image_id}; \
	done | jq .


# =================================================================
# for Lambda
# =================================================================
# make deploy AWS_ACCOUNT_ID=${AWS_ACCOUNT_ID} AWS_PROFILE=${AWS_PROFILE}
deploy: push
	aws lambda update-function-code \
		--profile ${AWS_PROFILE} \
		--function-name dev-reward-batch \
		--image-uri $(AWS_ACCOUNT_ID).dkr.ecr.ap-northeast-1.amazonaws.com/dev-reward-batch:$(VERSION) | jq .
