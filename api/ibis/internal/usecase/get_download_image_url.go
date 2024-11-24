package usecase

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/takeuchima0/async_serverless_application_demo/api/ibis/internal/configuration"
	"github.com/takeuchima0/async_serverless_application_demo/api/ibis/internal/library/storage"
)

type GetDownloadImageURLUseCase struct{}

func NewGetDownloadImageURL() *GetDownloadImageURLUseCase {
	return &GetDownloadImageURLUseCase{}
}

func (uc *GetDownloadImageURLUseCase) GetDownloadImageURL(ctx context.Context, uid string) (string, error) {

	bucketName := configuration.Get().S3.BucketName.OriginalImage
	s3Service := storage.NewStorageService(bucketName)

	// NOTE: 本来はこのパスから得られるuidを使用してDBなどにアクセスし、ファイルパスを取得する運用が望ましい
	uidStr, err := strconv.Atoi(uid)
	if err != nil {
		return "", fmt.Errorf("invalid user ID: %w", err)
	}
	filePath := fmt.Sprintf("profiles/%d/c2e188fa-efa9-2521-b5b2-64c5f674020e.jpg", uidStr)

	expire := 30 * time.Minute // 有効期限を30分に設定

	url, err := s3Service.GetDownloadPresignedURL(ctx, filePath, expire)
	if err != nil {
		return "", err
	}

	return url, nil
}
