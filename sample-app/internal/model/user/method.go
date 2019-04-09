package user

import "fmt"

// GetFullName get user's full name
func (u User) GetFullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}
