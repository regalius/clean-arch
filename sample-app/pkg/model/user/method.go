package user

import "fmt"

// getFullName get user's full name
func (u User) getFullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}
