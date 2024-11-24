# =================================================================
# for SQS
# =================================================================
QUEUE_URL_LIST := https://sqs.ap-northeast-1.amazonaws.com/${AWS_ACCOUNT_ID}/dev-to-notification-dlq https://sqs.ap-northeast-1.amazonaws.com/${AWS_ACCOUNT_ID}/dev-to-notifications-queue

.PHONY: list-sqs get-sqs-attributes get-sqs-details
# SQSのリストを取得
# make list-sqs AWS_PROFILE=${AWS_PROFILE}
list-sqs:
	aws sqs list-sqs --region ap-northeast-1 --profile ${AWS_PROFILE} | jq .

# SQSの属性情報を取得
# make get-sqs-attributes AWS_ACCOUNT_ID=${AWS_ACCOUNT_ID} AWS_PROFILE=${AWS_PROFILE} AWS_SQS_NAME=dev-to-notification-dlq
get-sqs-attributes:
	aws sqs get-sqs-attributes \
		--queue-url https://sqs.ap-northeast-1.amazonaws.com/${AWS_ACCOUNT_ID}/${AWS_SQS_NAME} \
		--attribute-name All \
		--region ap-northeast-1 \
		--profile ${AWS_PROFILE} | jq .

# SQSの詳細情報を取得、「--attribute-name」で取得したい項目を指定
get-sqs-details:
	aws sqs get-sqs-attributes \
		--queue-url https://sqs.ap-northeast-1.amazonaws.com/${AWS_ACCOUNT_ID}/${AWS_SQS_NAME} \
		--attribute-name ApproximateNumberOfMessages DelaySeconds ApproximateNumberOfMessagesNotVisible \
		--profile ${AWS_PROFILE} | jq .


.PHONY: get-sqs-attributes-list get-sqs-details-list
# 複数のSQSの属性を一括で取得
get-sqs-attributes-list:
	@for url in $(QUEUE_URL_LIST); do \
		echo "Getting attributes for queue: $$url"; \
		aws sqs get-sqs-attributes --queue-url $$url --attribute-name All --profile ${AWS_PROFILE} | jq .; \
	done

# 複数のSQSの詳細情報を取得、「--attribute-name」で取得したい項目を指定
get-sqs-details-list:
	@for url in $(QUEUE_URL_LIST); do \
		echo "Getting ApproximateNumberOfMessages for queue: $$url"; \
		aws sqs get-sqs-attributes --queue-url $$url \
			--attribute-name ApproximateNumberOfMessages DelaySeconds ApproximateNumberOfMessagesNotVisible \
			--profile ${AWS_PROFILE} | jq .; \
	done


.PHONY: get-sqs-message send-sqs-message delete-sqs-message
# SQSのキューを受信
get-sqs-message:
	aws sqs receive-message \
		--queue-url https://sqs.ap-northeast-1.amazonaws.com/${AWS_ACCOUNT_ID}/dev-to-notification-dlq \
		--max-number-of-messages 1 \
		--wait-time-seconds 10 \
		--profile ${AWS_PROFILE} | jq .

# 指定したキューに再度流し込む
send-sqs-message:
	aws sqs send-message \
		--queue-url https://sqs.ap-northeast-1.amazonaws.com/${AWS_ACCOUNT_ID}/dev-to-notifications-queue \
		--message-body '{"user_id":10010024,"order_id":"9a1882e2-b20e-4c48-93ef-a352d9502c18","status":"COMPLETED"}' \
		--profile ${AWS_PROFILE} | jq .

# メッセージの削除
delete-sqs-message:
	aws sqs delete-message \
		--queue-url https://sqs.ap-northeast-1.amazonaws.com/${AWS_ACCOUNT_ID}/dev-to-notification-dlq \
		--receipt-handle AQEBqVY61V5FSuoXKVaoeV27K05aiWprdunxcQ0i09DpDOE2zPfSDOsNjjsIjC1j/9OVNIawjge7fA54HuNKEoL3zYu4BmA2sFhaBDQNAdcgXiJQflVcQwTLXH9B3haBowQEnUDZrWc5t7YoIJlbClME5VjK7lOgBbkMhJSpnMs/b0m6t3BBYC71wWVat52y11hefK4NkufiiNauk8UXb4mWEe9EGDgo5TadIpfQMyjBbPotWCjU2vXilhgkF86q7VAE45ZNWnl+oYKVzLX1XKp0qrA84XDnzxz0O3Jek/MDQxOz2tIJ1NsGbCoQ3d2UW0HjVqDPDBIcEMQkjp85eNMjFJ5A3ZraL7T/qENcEMCs0Ki0Fi5XBm/VlZV2O0AXU/NaEqHDsbDQsBSYpTCrFCkLk2B0yAplDbf+Eo0+nN7tYKY= \
		--profile ${AWS_PROFILE} | jq .
