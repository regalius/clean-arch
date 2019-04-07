package shop

type (
	// MODELS

	// Shop : Model that represents Shop
	Shop struct {
		ID              int64
		Name            string
		PreviousNames   []string
		Type            int
		ReputationScore int64
		Images          []Image
		TimeCreated     int64
		TimeLastUpdated int64
		CreatedByID     int64
	}
	// Image : Model that represents Shop Image
	Image struct {
		Path string
		Name string
	}
	// END OF MODELS

	// STRUCTS

	// reputationBadgeMeta : Meta to determine shop's reputation badge
	reputationBadgeMeta struct {
		Lt   int64
		Gt   int64
		Name string
	}
	// END OF STRUCTS
)
