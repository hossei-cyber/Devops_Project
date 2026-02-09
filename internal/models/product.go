package models

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
}

var Products = []Product{
	{
		ID:          1,
		Name:        "Luxury Face Serum with Vitamin C",
		Price:       34.99,
		Category:    "Skincare",
		Description: "Brightening vitamin C serum with hyaluronic acid for glowing skin",
	},
	{
		ID:          2,
		Name:        "Essential Oil Diffuser Set",
		Price:       49.99,
		Category:    "Aromatherapy",
		Description: "Ultrasonic diffuser with lavender, eucalyptus, and peppermint oils",
	},
	{
		ID:          3,
		Name:        "Jade Roller & Gua Sha Set",
		Price:       24.99,
		Category:    "Wellness Tools",
		Description: "Natural jade stone tools for face massage and lymphatic drainage",
	},
	{
		ID:          4,
		Name:        "Organic Bath Bomb Collection",
		Price:       32.99,
		Category:    "Bath & Body",
		Description: "Set of 6 handmade bath bombs with natural ingredients and essential oils",
	},
	{
		ID:          5,
		Name:        "Silk Pillowcase Premium",
		Price:       39.99,
		Category:    "Sleep & Relaxation",
		Description: "100% mulberry silk pillowcase for hair and skin protection",
	},
	{
		ID:          6,
		Name:        "Mindfulness Journal & Pen Set",
		Price:       22.99,
		Category:    "Mental Wellness",
		Description: "Guided gratitude journal with prompts for daily reflection and mindfulness",
	},
	{
		ID:          7,
		Name:        "LED Light Therapy Mask",
		Price:       89.99,
		Category:    "Skincare",
		Description: "Professional-grade LED mask for anti-aging and acne treatment",
	},
	{
		ID:          8,
		Name:        "Himalayan Salt Lamp",
		Price:       28.99,
		Category:    "Aromatherapy",
		Description: "Hand-carved pink salt lamp for air purification and ambient lighting",
	},
}
