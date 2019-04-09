package shop

// ENUMS //
var (
	// shopTypesEnum Enum that can be used to infer shop's type to human readable string
	shopTypesEnum = map[int]string{
		1:  "Normal",
		2:  "Silver Member",
		3:  "Gold Member",
		4:  "Platinum Member",
		50: "Top Seller",
		90: "Ads User",
		91: "First Priority",
	}

	// shopReputationEnum Enum that can be used to infer shop's reputation score to human readable string
	shopReputationEnum = []reputationBadgeMeta{
		{
			GreaterThan: 0,
			LessThan:    100,
			Name:        "Newbs",
		},
		{
			GreaterThan: 101,
			LessThan:    1000,
			Name:        "Startups",
		},
		{
			GreaterThan: 1001,
			LessThan:    10000,
			Name:        "Corporations",
		},
		{
			GreaterThan: 10001,
			Name:        "Giants",
		},
	}
)
