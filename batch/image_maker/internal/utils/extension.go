package utils

// FileExtension: 拡張子を表す独自の型
type FileExtension string

// 画像ファイルの拡張子を表す定数
const (
	ExtJPEG FileExtension = ".jpeg"
	ExtJPG  FileExtension = ".jpg"
	ExtPNG  FileExtension = ".png"
)

// IsValidImageExtension: 拡張子が有効かどうかを確認するメソッド
func (ext FileExtension) IsValidImageExtension() bool {
	switch ext {
	case ExtJPEG, ExtJPG, ExtPNG:
		return true
	default:
		return false
	}
}
