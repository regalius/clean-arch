package product

type (
	// Product Model that represents Product
	Product struct {
		ID              int64
		Price           int64
		ShopID          int64
		TimeCreated     int64
		TimeLastUpdated int64
		CreatedByID     int64
		Name            string
		Images          []Image
	}
	// Image Model that represents Product Image
	Image struct {
		Path string
		Name string
	}
)
