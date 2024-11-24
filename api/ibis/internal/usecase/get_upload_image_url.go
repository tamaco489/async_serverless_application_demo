package usecase

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/takeuchima0/async_serverless_application_demo/api/ibis/internal/configuration"
	"github.com/takeuchima0/async_serverless_application_demo/api/ibis/internal/library/storage"
)

type GetUploadImageURLUseCase struct{}

func NewGetUploadImageURL() *GetUploadImageURLUseCase {
	return &GetUploadImageURLUseCase{}
}

func (uc *GetUploadImageURLUseCase) GetUploadImageURL(ctx context.Context, uid string) (string, error) {

	bucketName := configuration.Get().S3.BucketName.OriginalImage
	s3Service := storage.NewStorageService(bucketName)

	// NOTE: 本来はこのタイミングでuidに紐づくレコードが存在有無をチェックするのが望ましい
	// NOTE: 存在しない場合はinsert処理、存在する場合はupdate処理を行う
	// NOTE: このAPIが実行された時点では 'uploaded_at' 等のカラムをNULLとして設定しておく
	// NOTE: クライアントより、画像が実際にアップロードされた場合は、別途S3のPUTイベントLambdaを用意し、そっちでレコードを更新を行う

	uidStr, err := strconv.Atoi(uid)
	if err != nil {
		return "", fmt.Errorf("invalid user ID: %w", err)
	}

	// NOTE: 本来、画像の拡張子はクライアントのリクエスト経由でもらったほうが良い
	filePath := fmt.Sprintf("profiles/%d/d132bc66-d075-dbd8-9e77-5ad1b7129dca.png", uidStr)

	expire := 30 * time.Minute // 有効期限を30分に設定

	url, err := s3Service.GetUploadPresignedURL(ctx, filePath, expire)
	if err != nil {
		return "", err
	}

	return url, nil
}
