resource "aws_dynamodb_table" "user_table" {
  name = "users" # テーブル名

  # 削除保護設定 (削除されたくない場合はtrueで反映する)
  deletion_protection_enabled = false

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

  # オプション: TTL (Time to Live) 設定
  ttl {
    enabled        = false
    attribute_name = "ttl" # TTL(生存期間)を設定する場合
  }

  tags = { Name = "${var.env}-coral-api-dynamodb-users-table" }
}
