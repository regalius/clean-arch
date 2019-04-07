package product

type (
	// Filter List of filter that can be used to getByFilter in Product repository
	Filter struct {
		MinPrice         int64
		MaxPrice         int64
		NameContains     string
		ShopNameContains string
	}
)
