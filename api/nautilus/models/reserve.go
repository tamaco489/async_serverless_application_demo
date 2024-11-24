package models

type ReserveStatus string

const (
	ReserveStatusPending   ReserveStatus = "PENDING"   // 在庫確保中
	ReserveStatusConfirmed ReserveStatus = "CONFIRMED" // 在庫確保済み
	ReserveStatusFailed    ReserveStatus = "FAILED"    // 在庫確保失敗
	ReserveStatusCancelled ReserveStatus = "CANCELLED" // 在庫確保キャンセル
)

type ReserveRequest struct {
	EnsureProductList []EnsureProduct `json:"ensure_product_list"` // 確保する商品のリスト
	ReservedID        string          `json:"reserved_id"`         // 2つ目の商品を在庫確保する場合に使用、初回リクエスト時は空指定
}

type EnsureProduct struct {
	ID       int `json:"id"`       // 確保する商品のID
	Quantity int `json:"quantity"` // 確保する商品数
}

type ReserveResponse struct {
	Status     ReserveStatus `json:"status"`      // 在庫確保ステータス
	ReservedID string        `json:"reserved_id"` // 在庫確保時に一意に払い出されるUUID
}
