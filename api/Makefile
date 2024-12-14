# =================================================================
# 共通
# =================================================================
AWS_API_v1_DOMAIN := api.<api-domain>
AWS_API_v2_DOMAIN := apiv2.<api-domain>
AWS_API_v3_DOMAIN := apiv3.<api-domain>


# =================================================================
# 発注処理サービス
# =================================================================
.PHONY: nautilus-healthcheck get-products get-product-details reserve purchase

# ヘルスチェック
nautilus-healthcheck:
	curl -sX GET 'https://${AWS_API_v1_DOMAIN}/nautilus/v1/healthcheck' \
		-H "Content-Type: application/json" | jq .

# 商品一覧取得API
get-products:
	curl -sX GET 'https://${AWS_API_v1_DOMAIN}/nautilus/v1/products' \
		-H "Content-Type: application/json" | jq .

# 商品詳細取得API
get-product-details:
	curl -sX GET 'https://${AWS_API_v1_DOMAIN}/nautilus/v1/products/10010001' \
		-H "Content-Type: application/json" | jq .

# 在庫確保API
reserve:
	curl -sX POST 'https://${AWS_API_v1_DOMAIN}/nautilus/v1/products/reserve' \
		-H "Content-Type: application/json" \
		-d '{"ensure_product_list":[{"id":10010001,"quantity":2},{"id":10010002,"quantity":1}],"reserved_id":""}' | jq .

# 商品購入API
purchase:
	curl -sX POST 'https://${AWS_API_v1_DOMAIN}/nautilus/v1/products/purchase' \
		-H "Content-Type: application/json" \
		-H "X_API_KEY: 7GfPb5sDXkRSZAXBj3HFjaA3ayMfHQdzbBhT5UKXHKzMw8rXTtm7VHKASf3ZwHGw" \
		-d '{"reserved_id": "DEDBCC81-4B8E-4E50-AB71-E9DEA3EE2A15", "price_with_tax": 9960}' | jq .


# =================================================================
# 画像処理サービス
# =================================================================
.PHONY: ibis-healthcheck ibis-image-download ibis-image-upload ibis-client-image-upload

# ヘルスチェック
ibis-healthcheck:
	curl -sX GET 'https://${AWS_API_v2_DOMAIN}/ibis/v2/healthcheck' \
		-H "Content-Type: application/json" | jq .

# ダウンロード用署名付きURL発行
ibis-image-download:
	curl -sX GET 'https://${AWS_API_v2_DOMAIN}/ibis/v2/image/download/11020096' \
		-H "Content-Type: application/json" | jq .

# アップロード用署名付きURL発行
ibis-image-upload:
	curl -sX GET 'https://${AWS_API_v2_DOMAIN}/ibis/v2/image/upload/15726031' \
		-H "Content-Type: application/json" | jq .

# クライアント側の画像アップロード処理(事前にアップロード用の署名付きURLを発行して書き換える必要あり)
ibis-client-image-upload:
	curl -i -sX PUT --upload-file \
		/mnt/c/Users/nuro1/Desktop/mt/992_development_image/d132bc66-d075-dbd8-9e77-5ad1b7129dca.png \
		"https://dev-nautilus-original-image.s3.ap-northeast-1.amazonaws.com/profiles/15726031/d132bc66-d075-dbd8-9e77-5ad1b7129dca.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=ASIA4WU2Z6N37IKY3YHI%2F20241026%2Fap-northeast-1%2Fs3%2Faws4_request&X-Amz-Date=20241026T132531Z&X-Amz-Expires=1800&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEJ3%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEaDmFwLW5vcnRoZWFzdC0xIkYwRAIgWc2c0LdTy820Df0qcaZd0b%2FYli0QMN%2FCe2YwECPmB70CIB4D3%2B29AlB%2FlQZzm0NpFC9q9Lq84EFCwVrVSSL%2FE%2BXqKvECCBcQARoMODczMjc2NzY1MDQ3IgyvBq1EunG%2FOYstQpcqzgLgO%2BHv4ynL0kpBHQMspsnGsuFISrNMI1nyP9L1Tl1APhNjcUR%2F%2B7TRj%2FtB9at9rfdKa7uJ09yh7RaiaJTCSTarXqkZu2P%2BdKDceKuQjFBtmw2W8ev4HG3m8jC40tKzR20titTinKo63ijDwwYwDM72zkZ5OvUKSNyNCkx7Wn%2FhlP8c7Wjc3pTpQkZCyjinIvOyPsV6ch2kiYZeKRJIGswNw1ZB%2Fy8o7KOZXG6P2s9sYBBDT9xgd3YJdhm6Hxx50Y305vHjSUdTErZe6q%2BSaikrHWEEUzA6C0DX4iUkDqB0jMEXR227BlmGfLIdKYKOuvfklbq1Rwj0Fu%2FS1y0lZJeHPUWJ7s5xyMePZD7LAY5s18ijRi8cvrfMsCQ%2BWHIX2syBrByraTv5w%2Fw7WsLmaY3QuGoHHuBuB2zXTw1avVOXYa8UrZhjwtIYZo%2F2US1VMMvb87gGOp8BLnXeYVyd8LMq1PRNXTWrharVWB5d5hKkdM88lZknnLFwyT5MOL6xN7TnFGYCs6pV%2BkCbE%2BWGrLIPDT5esBdL9l1CHhhnRhliSSLJwfmtcUltOvakOEekV7Xd0KB9ToiZOmfW%2B9eE7G76gdvTfQXMBHfdzGTISMq3%2FnCwUTlNFg8RM40RlPHv9k4Ikxk7vXj5e8EUyDqwYIO%2Bz0jJkLMV&X-Amz-SignedHeaders=host&x-id=PutObject&X-Amz-Signature=ac3b1aabf4c2fb6de029d0e32031c3741ddc2ff96b48ac805b2aa0c0e82e6680"


# =================================================================
# ユーザ情報登録サービス
# =================================================================
.PHONY: coral-healthcheck coral-create-user

# ヘルスチェック
coral-healthcheck:
	curl -sX GET 'https://${AWS_API_v3_DOMAIN}/coral/v3/healthcheck' \
		-H "Content-Type: application/json" | jq .

# coral-healthcheck:
# 	curl -sX GET 'http://localhost:8082/coral/v3/healthcheck' \
# 		-H "Content-Type: application/json" | jq .

coral-create-user:
	curl -i -sX POST 'https://${AWS_API_v3_DOMAIN}/coral/v3/users' \
		-H "Content-Type: application/json" \
		-d '{"email":"ulala.halu@example.com", "birthday":"2000-12-01", "ekyc_status":"pending", "invite_code":"XYZ123", "is_admin":false}'

# coral-create-user:
# 	curl -i -sX POST 'http://localhost:8082/coral/v3/users' \
# 		-H "Content-Type: application/json" \
# 		-d '{"email":"ulala.halu@example.com", "birthday":"2000-12-01", "ekyc_status":"pending", "invite_code":"XYZ123", "is_admin":false}'

coral-getme-user:
	curl -sX GET 'https://${AWS_API_v3_DOMAIN}/coral/v3/users/me' \
		-H "Content-Type: application/json" | jq .

# coral-getme-user:
# 	curl -sX GET 'http://localhost:8082/coral/v3/users/me' \
# 		-H "Content-Type: application/json" | jq .

coral-get-user:
	curl -sX GET 'https://${AWS_API_v3_DOMAIN}/coral/v3/users/0193c4cc-1e5c-7188-ab2f-3856a4b73566' \
		-H "Content-Type: application/json" | jq .

# coral-get-user:
# 	curl -sX GET 'http://localhost:8082/coral/v3/users/0193c51c-d390-78a0-815e-57c8a763e3b5' \
# 		-H "Content-Type: application/json" | jq .


# =================================================================
# tfstate管理用 S3 backet の作成
# =================================================================
# S3バケットの作成
create-tfstate-bucket:
	aws --profile <aws-profile> s3 mb s3://dev-nautilus-tfstate

# バケットを指定してバージョニングを有効にする
update-tfstate-backet-versioning:
	aws --profile <aws-profile> s3api put-bucket-versioning \
		--bucket dev-nautilus-tfstate \
		--versioning-configuration Status=Enabled

# バケットの中身を全て削除し、バケット自体を削除する
delete-tfstate-bucket:
	aws --profile <aws-profile> s3 rm s3://dev-nautilus-tfstate --recursive
	aws --profile <aws-profile> s3 rb s3://dev-nautilus-tfstate
