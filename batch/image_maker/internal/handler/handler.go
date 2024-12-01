package handler

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/disintegration/imaging"
	"github.com/tamaco489/async_serverless_application_demo/batch/image_maker/internal/configuration"
	"github.com/tamaco489/async_serverless_application_demo/batch/image_maker/internal/utils"
)

type ImageMakerHandler struct {
	s3Client        *s3.Client
	dstBucket       string
	thumbnailWidth  int
	thumbnailHeight int
}

var (
	defaultWidth  = 100
	defaultHeight = 100
)

func NewImageMakerHandler(ctx context.Context) (*ImageMakerHandler, error) {
	if err := configuration.Load(ctx); err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(configuration.Get().AWSConfig)

	return &ImageMakerHandler{
		s3Client:        s3Client,
		dstBucket:       configuration.Get().S3.BucketName.ThumbnailImage,
		thumbnailWidth:  defaultWidth,
		thumbnailHeight: defaultHeight,
	}, nil
}

func (h *ImageMakerHandler) Handler(ctx context.Context, s3Event events.S3Event) error {

	log.Println("image maker batch server started...")

	for _, r := range s3Event.Records {
		slog.Info("Start Lambda event.",
			"event_time:", r.EventTime, // 2024-10-20T15:09:44.149Z
			"bucket_name:", r.S3.Bucket.Name, // dev-nautilus-original-image
			"object_key:", r.S3.Object.Key, // profiles/15726031/d132bc66-d075-dbd8-9e77-5ad1b7129dca.png
		)

		srcBucket, key := r.S3.Bucket.Name, r.S3.Object.Key
		originExt := filepath.Ext(key)
		ext := utils.FileExtension(strings.ToLower(originExt))
		if !ext.IsValidImageExtension() {
			slog.WarnContext(ctx, "warning image extension", "not supported extension: ", originExt)
		}

		// オリジナル画像コンテンツの取得
		originalImg, err := h.downloadImage(ctx, srcBucket, key)
		if err != nil {
			slog.WarnContext(ctx, "warning download error", "error", err)
			continue
		}

		// サムネイル画像生成
		thumbnailImg := imaging.Thumbnail(originalImg, h.thumbnailWidth, h.thumbnailHeight, imaging.Lanczos)

		// サムネイル画像のアップロード
		if err := h.uploadThumbnail(ctx, thumbnailImg, key, ext); err != nil {
			slog.WarnContext(ctx, "warning upload error", "error", err)
			continue
		}

		log.Println("thumbnail image uploaded successfully...")
	}

	return nil
}

// downloadImage: S3から画像コンテンツをダウンロード
func (h *ImageMakerHandler) downloadImage(ctx context.Context, bucket, key string) (image.Image, error) {

	resp, err := h.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return img, err
	}

	return img, nil
}

// uploadThumbnail サムネイルを別のS3バケットにアップロード
func (h *ImageMakerHandler) uploadThumbnail(ctx context.Context, img image.Image, key string, ext utils.FileExtension) error {

	buf := new(bytes.Buffer)

	switch ext {
	case utils.ExtJPEG, utils.ExtJPG:
		if err := jpeg.Encode(buf, img, nil); err != nil {
			return err
		}

	case utils.ExtPNG:
		if err := png.Encode(buf, img); err != nil {
			return err
		}

	default:
		return fmt.Errorf("unsupported image type")
	}

	contentType := aws.String("image/" + string(ext[1:])) // 先頭の1文字を除外し、Content-Typeを設定 (ex: .jpn -> jpg)

	_, err := h.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(h.dstBucket),
		Key:         aws.String(key),
		Body:        bytes.NewBuffer(buf.Bytes()),
		ContentType: contentType,
	})
	if err != nil {
		return err
	}

	return nil
}
