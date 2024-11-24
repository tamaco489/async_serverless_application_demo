# # S3バケットリソースの定義
# resource "aws_s3_bucket" "example" {
#   bucket = "supply-chain-demo-bucket"
# }

# # S3バケットのバージョニング設定
# resource "aws_s3_bucket_versioning" "versioning_example" {
#   bucket = aws_s3_bucket.example.id
#   versioning_configuration {
#     status = "Enabled"
#   }
# }

# # publicアクセスを全てブロックする設定
# resource "aws_s3_bucket_public_access_block" "example" {
#   bucket = aws_s3_bucket.example.id

#   block_public_acls       = true
#   block_public_policy     = true
#   ignore_public_acls      = true
#   restrict_public_buckets = true
# }

# # S3バケットのサーバーサイド暗号化設定
# resource "aws_s3_bucket_server_side_encryption_configuration" "sse_configuration_example" {
#   bucket = aws_s3_bucket.example.id

#   rule {
#     apply_server_side_encryption_by_default {
#       sse_algorithm = "AES256"
#     }
#   }
# }

# resource "aws_s3_bucket_lifecycle_configuration" "lifecycle_example" {
#   bucket = aws_s3_bucket.example.id

#   rule {
#     # ルールの一意識別子を設定
#     id = "logs-deletion-rule"

#     # ルールの有効化状態を設定
#     status = "Enabled"

#     # フィルタ設定: プレフィックスとサイズに基づいてオブジェクトをフィルタリング
#     filter {
#       and {
#         # "logs/" プレフィックスで始まるオブジェクトに対してルールを適用
#         prefix = "logs/"

#         # 最大オブジェクトサイズを設定 (例: 10 MB = 10 * 1024 * 1024 バイト)
#         object_size_less_than = 10485760
#       }
#     }

#     # オブジェクトが作成されてから30日後にSTANDARD_IAストレージクラスに移行
#     transition {
#       days          = 30
#       storage_class = "STANDARD_IA"
#     }

#     # オブジェクトが作成されてから60日後にGLACIERストレージクラスに移行
#     transition {
#       days          = 60
#       storage_class = "GLACIER"
#     }

#     # オブジェクトが作成されてから90日後に削除
#     expiration {
#       days = 90
#     }
#   }
# }
