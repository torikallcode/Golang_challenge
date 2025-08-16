package models

type User struct {
	ID    int
	Name  string
	Email string
}

func (u *User) UpdateEmail(newEmail string) {
	u.Email = newEmail
}
