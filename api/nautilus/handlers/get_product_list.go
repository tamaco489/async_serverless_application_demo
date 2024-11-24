package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/takeuchima0/async_serverless_application_demo/api/nautilus/models"
)

// 商品一覧取得API
func GetProductList(w http.ResponseWriter, r *http.Request) {

	// Validate HTTP Method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Mock response
	response := models.ProductListResponse{
		List: []models.Product{
			{
				ID:           10010001,
				Name:         "商品A",
				Category:     "カテゴリー1",
				Description:  "",
				PriceWithTax: 4980,
				ImageURL:     "https://<domain>/images/thumbnail_image_path_1",
				InitialStock: 8,
				CurrentStock: 30,
			},
			{
				ID:           10010002,
				Name:         "商品B",
				Category:     "カテゴリー1",
				Description:  "",
				PriceWithTax: 4980,
				ImageURL:     "https://<domain>/images/thumbnail_image_path_2",
				InitialStock: 0,
				CurrentStock: 30,
			},
			{
				ID:           10010003,
				Name:         "商品C",
				Category:     "カテゴリー1",
				Description:  "",
				PriceWithTax: 4980,
				ImageURL:     "https://<domain>/images/thumbnail_image_path_3",
				InitialStock: 30,
				CurrentStock: 30,
			},
		},
	}

	//  Encode response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
