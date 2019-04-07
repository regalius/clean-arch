package user

type (
	// User model to represents a user
	User struct {
		ID        int64
		FirstName string
		LastName  string
		Gender    float32
		Address   string
	}
)
