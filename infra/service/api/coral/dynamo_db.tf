resource "aws_dynamodb_table" "user_table" {
  name         = "users"           # テーブル名
  billing_mode = "PAY_PER_REQUEST" # 従量課金の料金モード
  hash_key     = "user_id"         # プライマリキー
  attribute {
    name = "user_id"
    type = "S" # SはString
  }

  # オプション: 時間情報の管理
  stream_enabled   = true
  stream_view_type = "NEW_AND_OLD_IMAGES" # 変更前後のデータを保存

  # オプション: テーブルの設定（必要に応じて）
  ttl {
    attribute_name = "ttl" # TTL(生存期間)を設定する場合
    enabled        = false
  }
}
