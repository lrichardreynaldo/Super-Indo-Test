package utils

import models "superIndo/model"

var (
	CategoriesProduct = []models.Category{
		{
			ID:           1,
			CategoryName: "Elektronik",
		},
		{
			ID:           2,
			CategoryName: "Makanan",
		},
		{
			ID:           3,
			CategoryName: "Minuman",
		},
		{
			ID:           4,
			CategoryName: "Kebersihan",
		},
		{
			ID:           5,
			CategoryName: "Rumah Tangga",
		},
		{
			ID:           6,
			CategoryName: "AI",
		},
	}

	Products = []models.Product{
		{
			ID:          1,
			ProductName: "Smartphone",
			CategoryID:  1,
		},
		{
			ID:          2,
			ProductName: "Laptop",
			CategoryID:  1,
		},
		{
			ID:          3,
			ProductName: "Nasi Goreng",
			CategoryID:  2,
		},
		{
			ID:          4,
			ProductName: "Jus Apel",
			CategoryID:  3,
		},
		{
			ID:          5,
			ProductName: "Soda",
			CategoryID:  3,
		},
		{
			ID:          6,
			ProductName: "Sabun Cuci",
			CategoryID:  4,
		},
		{
			ID:          7,
			ProductName: "Sapu",
			CategoryID:  5,
		},
	}

	ProductDetails = []models.ProductDetail{
		{
			ProductID:   1,
			Description: "HP tercanggih baru merk XXX",
			Price:       2349999.99,
			Stock:       50,
		},
		{
			ProductID:   2,
			Description: "Laptop gaming terbaru",
			Price:       12399479.99,
			Stock:       30,
		},
		{
			ProductID:   3,
			Description: "Makanan tradisional Indonesia",
			Price:       24000,
			Stock:       200,
		},
		{
			ProductID:   4,
			Description: "Jus apel yang baru dibikin",
			Price:       8999.99,
			Stock:       150,
		},
		{
			ProductID:   5,
			Description: "Minuman soda yang refreshing",
			Price:       6000,
			Stock:       300,
		},
		{
			ProductID:   6,
			Description: "Sabun untuk mencuci piring",
			Price:       5500,
			Stock:       500,
		},
		{
			ProductID:   7,
			Description: "Sapu tahan banting",
			Price:       56799.99,
			Stock:       100,
		},
	}
)
