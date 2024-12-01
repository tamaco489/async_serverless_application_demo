package main

import (
	"log"
	"net/http"

	"github.com/tamaco489/async_serverless_application_demo/api/nautilus/handlers"
)

// 発注処理サーバの実装
func main() {
	log.Println("Nautilus API Server started on 8080")

	// ヘルスチェックAPI
	http.HandleFunc("/nautilus/v1/healthcheck", handlers.HealthCheckHandler)

	// 商品一覧取得API
	http.HandleFunc("/nautilus/v1/products", handlers.GetProductList)

	// 商品詳細取得API
	http.HandleFunc("/nautilus/v1/products/", handlers.GetProductByID)

	// 在庫確保API
	http.HandleFunc("/nautilus/v1/products/reserve", handlers.ReserveHandler)

	// 商品購入API
	http.HandleFunc("/nautilus/v1/products/purchase", handlers.PurchaseHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start Nautilus API Server: %v", err)
	}
}
