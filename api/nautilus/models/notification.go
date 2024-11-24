package models

type ProductOrderCompleteRequest struct {
	OrderID string `json:"order_id"` // 商品購入処理が正常に行われた際に発行される注文ID
}
