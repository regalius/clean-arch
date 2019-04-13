package config

type (
	// NewConfigOptions options for initializing new config instance
	NewConfigOptions struct {
		Defaults map[string]interface{}
		IsWatch  bool
	}
)
