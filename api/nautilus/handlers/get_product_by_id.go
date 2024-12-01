package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/tamaco489/async_serverless_application_demo/api/nautilus/models"
)

// 商品詳細取得API
func GetProductByID(w http.ResponseWriter, r *http.Request) {

	// Validate HTTP Method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Retrieve product id from path parameter
	path := strings.TrimPrefix(r.URL.Path, "/nautilus/v1/products/")
	productID, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mock response
	response := models.ProductDetailResponse{
		Product: models.Product{
			ID:           int64(productID),
			Name:         "商品A",
			Category:     "カテゴリー1",
			Description:  "",
			PriceWithTax: 4980,
			ImageURL:     "https://<domain>/images/thumbnail_image_path_1",
			InitialStock: 8,
			CurrentStock: 30,
		},
		Dimensions: models.Dimensions{
			Weight: 24.4,
			Height: 16.5,
			Width:  5.0,
			Depth:  5.0,
		},
		Warranty: models.Warranty{
			DurationMonth: 24,
			Details:       "Warranty Details",
		},
		Tags: []models.Tag{
			{
				ID:   4,
				Name: "タグ4",
			},
			{
				ID:   23,
				Name: "タグ23",
			},
		},
		RelatedProducts: []models.RelatedProduct{
			{
				ID:           10010024,
				Name:         "商品X",
				PriceWithTax: 5980,
				Category:     "カテゴリー",
				ImageURL:     "https://<domain>/images/thumbnail_image_path_24",
			},
			{
				ID:           10010025,
				Name:         "商品Y",
				PriceWithTax: 5980,
				Category:     "カテゴリー",
				ImageURL:     "https://<domain>/images/thumbnail_image_path_25",
			},
			{
				ID:           10010026,
				Name:         "商品Z",
				PriceWithTax: 5980,
				Category:     "カテゴリー",
				ImageURL:     "https://<domain>/images/thumbnail_image_path_26",
			},
		},
	}

	// Encode response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
