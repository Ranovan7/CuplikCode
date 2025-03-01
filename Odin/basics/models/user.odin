package models

import "users"

User :: struct {
	id:     int,
	name:   string,
	status: users.UserStatus,
}

new_user :: proc() -> User {
	u := User {
		name   = "Bob",
		id     = 1,
		status = users.UserStatus.Active,
	}

	return u
}

