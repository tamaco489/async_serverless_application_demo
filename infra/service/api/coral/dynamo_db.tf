resource "aws_dynamodb_table" "user_table" {
  name = "users" # テーブル名

  # 料金モード
  billing_mode = "PAY_PER_REQUEST" # 従量課金モードに設定

  # キャパシティ設定に関して、従量課金モードの場合は読み書きのスループットはAWSによって管理されるため明示的な設定は不要
  # read_capacity  = 20
  # write_capacity = 20

  # プライマリキー
  hash_key  = "user_id"    # Partition Key
  range_key = "created_at" # Sort Key

  # プライマリキーの属性定義
  attribute {
    name = "user_id"
    type = "S"
  }

  # # インデックスで使用する属性の定義
  attribute {
    name = "email"
    type = "S"
  }

  attribute {
    name = "ekyc_status"
    type = "S"
  }

  attribute {
    name = "created_at"
    type = "S"
  }


  # ローカルセカンダリインデックスの定義
  local_secondary_index {
    name            = "email_index"
    range_key       = "email"     # emailを範囲キーに使用
    projection_type = "KEYS_ONLY" # キーのみをインデックスに含める
  }

  # グローバルセカンダリインデックスの定義
  global_secondary_index {
    name            = "ekyc_status_index"
    hash_key        = "ekyc_status" # ekyc_statusをハッシュキーとして使用
    projection_type = "ALL"         # インデックス内に全ての属性を含める
  }

  # # オプション: DynamoDB Streamsの設定
  # stream_enabled   = true
  # stream_view_type = "NEW_AND_OLD_IMAGES" # 変更前後のデータを保存

  # オプション: TTL (Time to Live) 設定
  ttl {
    attribute_name = "ttl" # TTL(生存期間)を設定する場合
    enabled        = false
  }
}
