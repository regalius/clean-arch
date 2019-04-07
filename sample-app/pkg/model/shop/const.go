package shop

// ENUMS //
var (
	// shopTypesEnum: To infer shop's type to human readable string
	shopTypesEnum = map[int]string{
		1:  "Normal",
		2:  "Silver Member",
		3:  "Gold Member",
		4:  "Platinum Member",
		50: "Top Seller",
		90: "Ads User",
		91: "First Priority",
	}

	// shopReputationEnum: To infer shop's reputation score to human readable string
	shopReputationEnum = []reputationBadgeMeta{
		{
			Gt:   0,
			Lt:   100,
			Name: "Newbs",
		},
		{
			Gt:   101,
			Lt:   1000,
			Name: "Startups",
		},
		{
			Gt:   1001,
			Lt:   10000,
			Name: "Corporations",
		},
		{
			Gt:   10001,
			Name: "Giants",
		},
	}
)
