package shop

type (
	// Shop Model that represents Shop
	Shop struct {
		ID              int64
		ReputationScore int64
		TimeCreated     int64
		TimeLastUpdated int64
		CreatedByID     int64
		Type            int
		Name            string
		PreviousNames   []string
		Images          []Image
	}
	// Image Model that represents Shop Image
	Image struct {
		Path string
		Name string
	}
	// reputationBadgeMeta Meta to determine shop's reputation badge
	reputationBadgeMeta struct {
		GreaterThan int64
		LessThan    int64
		Name        string
	}
)
