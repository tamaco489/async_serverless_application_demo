# =================================================================
# 発注処理サービス
# =================================================================
resource "aws_ecr_repository" "nautilus_api" {
  name = "${var.env}-nautilus-api"

  # 既存のタグに対して、後から上書きを可能とする設定
  image_tag_mutability = "MUTABLE"

  # イメージがpushされる度に、自動的にセキュリティスキャンを行う設定を有効にする
  image_scanning_configuration {
    scan_on_push = true
  }

  tags = { Name = "${var.env}-nautilus-api-ecr" }
}

# ライフサイクルポリシーの設定
resource "aws_ecr_lifecycle_policy" "nautilus_api" {
  repository = aws_ecr_repository.nautilus_api.name

  policy = jsonencode(
    {
      "rules" : [
        {
          "rulePriority" : 1,
          "description" : "バージョン付きのイメージを20個保持する、21個目がアップロードされた際には古いものから順に削除されていく",
          "selection" : {
            "tagStatus" : "tagged",
            "tagPrefixList" : ["nautilus_v"],
            "countType" : "imageCountMoreThan",
            "countNumber" : 20
          },
          "action" : {
            "type" : "expire"
          }
        },
        {
          "rulePriority" : 2,
          "description" : "タグが設定されていないイメージをアップロードされてから3日後に削除する",
          "selection" : {
            "tagStatus" : "untagged",
            "countType" : "sinceImagePushed",
            "countUnit" : "days",
            "countNumber" : 3
          },
          "action" : {
            "type" : "expire"
          }
        },
        {
          "rulePriority" : 3,
          "description" : "タグが設定されたイメージをアップロードされてから30日後に削除する",
          "selection" : {
            "tagStatus" : "any",
            "countType" : "sinceImagePushed",
            "countUnit" : "days",
            "countNumber" : 30
          },
          "action" : {
            "type" : "expire"
          }
        }
      ]
    }
  )
}

# =================================================================
# 画像処理サービス
# =================================================================
resource "aws_ecr_repository" "ibis_api" {
  name = "${var.env}-ibis-api"

  # 既存のタグに対して、後から上書きを可能とする設定
  image_tag_mutability = "MUTABLE"

  # イメージがpushされる度に、自動的にセキュリティスキャンを行う設定を有効にする
  image_scanning_configuration {
    scan_on_push = true
  }

  tags = { Name = "${var.env}-ibis-api-ecr" }
}

# ライフサイクルポリシーの設定
resource "aws_ecr_lifecycle_policy" "ibis_api" {
  repository = aws_ecr_repository.ibis_api.name

  policy = jsonencode(
    {
      "rules" : [
        {
          "rulePriority" : 1,
          "description" : "バージョン付きのイメージを20個保持する、21個目がアップロードされた際には古いものから順に削除されていく",
          "selection" : {
            "tagStatus" : "tagged",
            "tagPrefixList" : ["ibis_v"],
            "countType" : "imageCountMoreThan",
            "countNumber" : 20
          },
          "action" : {
            "type" : "expire"
          }
        },
        {
          "rulePriority" : 2,
          "description" : "タグが設定されていないイメージをアップロードされてから3日後に削除する",
          "selection" : {
            "tagStatus" : "untagged",
            "countType" : "sinceImagePushed",
            "countUnit" : "days",
            "countNumber" : 3
          },
          "action" : {
            "type" : "expire"
          }
        },
        {
          "rulePriority" : 3,
          "description" : "タグが設定されたイメージをアップロードされてから30日後に削除する",
          "selection" : {
            "tagStatus" : "any",
            "countType" : "sinceImagePushed",
            "countUnit" : "days",
            "countNumber" : 30
          },
          "action" : {
            "type" : "expire"
          }
        }
      ]
    }
  )
}

# =================================================================
# ユーザ情報登録サービス
# =================================================================
resource "aws_ecr_repository" "coral_api" {
  name = "${var.env}-coral-api"

  # 既存のタグに対して、後から上書きを可能とする設定
  image_tag_mutability = "MUTABLE"

  # イメージがpushされる度に、自動的にセキュリティスキャンを行う設定を有効にする
  image_scanning_configuration {
    scan_on_push = true
  }

  tags = { Name = "${var.env}-coral-api-ecr" }
}

# ライフサイクルポリシーの設定
resource "aws_ecr_lifecycle_policy" "coral_api" {
  repository = aws_ecr_repository.coral_api.name

  policy = jsonencode(
    {
      "rules" : [
        {
          "rulePriority" : 1,
          "description" : "バージョン付きのイメージを20個保持する、21個目がアップロードされた際には古いものから順に削除されていく",
          "selection" : {
            "tagStatus" : "tagged",
            "tagPrefixList" : ["coral_v"],
            "countType" : "imageCountMoreThan",
            "countNumber" : 20
          },
          "action" : {
            "type" : "expire"
          }
        },
        {
          "rulePriority" : 2,
          "description" : "タグが設定されていないイメージをアップロードされてから3日後に削除する",
          "selection" : {
            "tagStatus" : "untagged",
            "countType" : "sinceImagePushed",
            "countUnit" : "days",
            "countNumber" : 3
          },
          "action" : {
            "type" : "expire"
          }
        },
        {
          "rulePriority" : 3,
          "description" : "タグが設定されたイメージをアップロードされてから30日後に削除する",
          "selection" : {
            "tagStatus" : "any",
            "countType" : "sinceImagePushed",
            "countUnit" : "days",
            "countNumber" : 30
          },
          "action" : {
            "type" : "expire"
          }
        }
      ]
    }
  )
}


# =================================================================
# サムネイル画像生成バッチ
# =================================================================
resource "aws_ecr_repository" "image_maker_batch" {
  name = "${var.env}-image-maker-batch"

  # 既存のタグに対して、後から上書きを可能とする設定
  image_tag_mutability = "MUTABLE"

  # イメージがpushされる度に、自動的にセキュリティスキャンを行う設定を有効にする
  image_scanning_configuration {
    scan_on_push = true
  }

  tags = { Name = "${var.env}-image-maker-batch-ecr" }
}

# ライフサイクルポリシーの設定
resource "aws_ecr_lifecycle_policy" "image_maker_batch" {
  repository = aws_ecr_repository.image_maker_batch.name

  policy = jsonencode(
    {
      "rules" : [
        {
          "rulePriority" : 1,
          "description" : "バージョン付きのイメージを20個保持する、21個目がアップロードされた際には古いものから順に削除されていく",
          "selection" : {
            "tagStatus" : "tagged",
            "tagPrefixList" : ["image_maker_v"],
            "countType" : "imageCountMoreThan",
            "countNumber" : 20
          },
          "action" : {
            "type" : "expire"
          }
        },
        {
          "rulePriority" : 2,
          "description" : "タグが設定されていないイメージをアップロードされてから3日後に削除する",
          "selection" : {
            "tagStatus" : "untagged",
            "countType" : "sinceImagePushed",
            "countUnit" : "days",
            "countNumber" : 3
          },
          "action" : {
            "type" : "expire"
          }
        },
        {
          "rulePriority" : 3,
          "description" : "タグが設定されたイメージをアップロードされてから30日後に削除する",
          "selection" : {
            "tagStatus" : "any",
            "countType" : "sinceImagePushed",
            "countUnit" : "days",
            "countNumber" : 30
          },
          "action" : {
            "type" : "expire"
          }
        }
      ]
    }
  )
}

# =================================================================
# メッセージ通知バッチ
# =================================================================
resource "aws_ecr_repository" "notification_batch" {
  name = "${var.env}-notification-batch"

  # 既存のタグに対して、後から上書きを可能とする設定
  image_tag_mutability = "MUTABLE"

  # イメージがpushされる度に、自動的にセキュリティスキャンを行う設定を有効にする
  image_scanning_configuration {
    scan_on_push = true
  }

  tags = { Name = "${var.env}-notification-batch-ecr" }
}

# ライフサイクルポリシーの設定
resource "aws_ecr_lifecycle_policy" "notification_batch" {
  repository = aws_ecr_repository.notification_batch.name

  policy = jsonencode(
    {
      "rules" : [
        {
          "rulePriority" : 1,
          "description" : "バージョン付きのイメージを20個保持する、21個目がアップロードされた際には古いものから順に削除されていく",
          "selection" : {
            "tagStatus" : "tagged",
            "tagPrefixList" : ["notification_v"],
            "countType" : "imageCountMoreThan",
            "countNumber" : 20
          },
          "action" : {
            "type" : "expire"
          }
        },
        {
          "rulePriority" : 2,
          "description" : "タグが設定されていないイメージをアップロードされてから3日後に削除する",
          "selection" : {
            "tagStatus" : "untagged",
            "countType" : "sinceImagePushed",
            "countUnit" : "days",
            "countNumber" : 3
          },
          "action" : {
            "type" : "expire"
          }
        },
        {
          "rulePriority" : 3,
          "description" : "タグが設定されたイメージをアップロードされてから30日後に削除する",
          "selection" : {
            "tagStatus" : "any",
            "countType" : "sinceImagePushed",
            "countUnit" : "days",
            "countNumber" : 30
          },
          "action" : {
            "type" : "expire"
          }
        }
      ]
    }
  )
}

# =================================================================
# ランキング集計バッチ
# =================================================================
resource "aws_ecr_repository" "rank_batch" {
  name = "${var.env}-rank-batch"

  # 既存のタグに対して、後から上書きを可能とする設定
  image_tag_mutability = "MUTABLE"

  # イメージがpushされる度に、自動的にセキュリティスキャンを行う設定を有効にする
  image_scanning_configuration {
    scan_on_push = true
  }

  tags = { Name = "${var.env}-rank-batch-ecr" }
}

# ライフサイクルポリシーの設定
resource "aws_ecr_lifecycle_policy" "rank_batch" {
  repository = aws_ecr_repository.rank_batch.name

  policy = jsonencode(
    {
      "rules" : [
        {
          "rulePriority" : 1,
          "description" : "バージョン付きのイメージを20個保持する、21個目がアップロードされた際には古いものから順に削除されていく",
          "selection" : {
            "tagStatus" : "tagged",
            "tagPrefixList" : ["rank_v"],
            "countType" : "imageCountMoreThan",
            "countNumber" : 20
          },
          "action" : {
            "type" : "expire"
          }
        },
        {
          "rulePriority" : 2,
          "description" : "タグが設定されていないイメージをアップロードされてから3日後に削除する",
          "selection" : {
            "tagStatus" : "untagged",
            "countType" : "sinceImagePushed",
            "countUnit" : "days",
            "countNumber" : 3
          },
          "action" : {
            "type" : "expire"
          }
        },
        {
          "rulePriority" : 3,
          "description" : "タグが設定されたイメージをアップロードされてから30日後に削除する",
          "selection" : {
            "tagStatus" : "any",
            "countType" : "sinceImagePushed",
            "countUnit" : "days",
            "countNumber" : 30
          },
          "action" : {
            "type" : "expire"
          }
        }
      ]
    }
  )
}

# =================================================================
# 報酬管理バッチ
# =================================================================
resource "aws_ecr_repository" "reward_batch" {
  name = "${var.env}-reward-batch"

  # 既存のタグに対して、後から上書きを可能とする設定
  image_tag_mutability = "MUTABLE"

  # イメージがpushされる度に、自動的にセキュリティスキャンを行う設定を有効にする
  image_scanning_configuration {
    scan_on_push = true
  }

  tags = { Name = "${var.env}-reward-batch-ecr" }
}

# ライフサイクルポリシーの設定
resource "aws_ecr_lifecycle_policy" "reward_batch" {
  repository = aws_ecr_repository.reward_batch.name

  policy = jsonencode(
    {
      "rules" : [
        {
          "rulePriority" : 1,
          "description" : "バージョン付きのイメージを20個保持する、21個目がアップロードされた際には古いものから順に削除されていく",
          "selection" : {
            "tagStatus" : "tagged",
            "tagPrefixList" : ["reward_v"],
            "countType" : "imageCountMoreThan",
            "countNumber" : 20
          },
          "action" : {
            "type" : "expire"
          }
        },
        {
          "rulePriority" : 2,
          "description" : "タグが設定されていないイメージをアップロードされてから3日後に削除する",
          "selection" : {
            "tagStatus" : "untagged",
            "countType" : "sinceImagePushed",
            "countUnit" : "days",
            "countNumber" : 3
          },
          "action" : {
            "type" : "expire"
          }
        },
        {
          "rulePriority" : 3,
          "description" : "タグが設定されたイメージをアップロードされてから30日後に削除する",
          "selection" : {
            "tagStatus" : "any",
            "countType" : "sinceImagePushed",
            "countUnit" : "days",
            "countNumber" : 30
          },
          "action" : {
            "type" : "expire"
          }
        }
      ]
    }
  )
}
