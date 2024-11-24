package models

import "time"

type OriginalMessage interface{}

type DLQMessage struct {
	OriginalMessage OriginalMessage
	Error           Error
	TimeStamp       time.Time
	LambdaFunction  string
}

type ErrorCode string

const (
	ErrorCodeInvalidStatus    ErrorCode = "INVALID_STATUS"    // 無効なステータス
	ErrorCodeProcessingFailed ErrorCode = "PROCESSING_FAILED" // 処理失敗
	ErrorCodeUnknownError     ErrorCode = "UNKNOWN_ERROR"     // 不明なエラー
)

type Error struct {
	Code    ErrorCode
	Message string
}
