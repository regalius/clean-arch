package http

type (
	GetProductRecomResponse struct {
		Meta Meta              `json:"meta"`
		Data []ProductResponse `json:"data"`
	}

	Meta struct {
		OverallAffinities float32 `json:"overal_affinity"`
	}

	ProductResponse struct {
		ID              int64    `json:"id"`
		Price           int64    `json:"price"`
		ShopID          int64    `json:"shop_id"`
		TimeCreated     int64    `json:"created"`
		TimeLastUpdated int64    `json:"-"`
		CreatedByID     int64    `json:"-"`
		URL             string   `json:"url"`
		Name            string   `json:"name"`
		Images          []string `json:"images"`
		AffinityScore   float32  `json:"affinity"`
	}
)
