package models

type Product struct {
	ID           int64   `json:"id"`             // 商品ID
	Name         string  `json:"name"`           // 商品名
	Category     string  `json:"category"`       // 商品カテゴリ
	Description  string  `json:"description"`    // 商品概要
	PriceWithTax float64 `json:"price_with_tax"` // 商品価格(税込み)
	ImageURL     string  `json:"image_url"`      // 商品のサムネイル画像URL
	InitialStock uint32  `json:"initial_stock"`  // 元々の在庫数
	CurrentStock uint32  `json:"current_stock"`  // 現在の在庫数
}

type Dimensions struct {
	Weight float64 `json:"weight"` // 重量
	Height float64 `json:"height"` // 高さ
	Width  float64 `json:"width"`  // 幅
	Depth  float64 `json:"depth"`  // 奥行き
}

type Warranty struct {
	DurationMonth uint32 `json:"duration_month"` // 保障期間(月単位)
	Details       string `json:"details"`        // 保障内容の詳細
}

type Tag struct {
	ID   int32  `json:"id"`   // 商品のタグID
	Name string `json:"name"` // タグ名
}

type RelatedProduct struct {
	ID           int     `json:"id"`             // 関連商品のID
	Name         string  `json:"name"`           // 関連商品の名前
	PriceWithTax float64 `json:"price_with_tax"` // 関連商品の価格(税込み)
	Category     string  `json:"category"`       // 関連商品カテゴリ
	ImageURL     string  `json:"image_url"`      // 関連商品のサムネイル画像URL
}

type ProductListResponse struct {
	List []Product `json:"list"`
}

type ProductDetailResponse struct {
	Product         Product          `json:"product"`
	Dimensions      Dimensions       `json:"dimensions"`       // 商品の物理的な特徴
	Warranty        Warranty         `json:"warranty"`         // 保障情報
	Tags            []Tag            `json:"tags"`             // 商品に関連付けられたタグ
	RelatedProducts []RelatedProduct `json:"related_products"` // 関連商品
}
