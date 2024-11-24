package models

type PurchaseStatus string

const (
	PurchaseStatusPending    PurchaseStatus = "PENDING"    // 購入処理待ち
	PurchaseStatusProcessing PurchaseStatus = "PROCESSING" // 購入処理中
	PurchaseStatusCompleted  PurchaseStatus = "COMPLETED"  // 購入完了
	PurchaseStatusFailed     PurchaseStatus = "FAILED"     // 購入失敗
	PurchaseStatusCancelled  PurchaseStatus = "CANCELLED"  // 購入キャンセル
)

type PurchaseRequest struct {
	ReservedID   string  `json:"reserved_id"`    // 在庫確保APIで払い出されたID
	PriceWithTax float64 `json:"price_with_tax"` // 注文する商品の合計金額 (税込)
}

type PurchaseResponse struct {
	Status  PurchaseStatus `json:"status"`   // 商品購入ステータス
	OrderID string         `json:"order_id"` // 注文受付時に払い出されるUUID
}

// SQSに送信するキューの構造体
type PurchaseQueueMessage struct {
	UserID  int            `json:"user_id"`
	OrderID string         `json:"order_id"`
	Status  PurchaseStatus `json:"status"`
}
