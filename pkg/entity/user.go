package entity

// User is the entity that represents a user.
type User struct {
	ID    int
	Name  string
	Email string
}

type CreateUserPayload struct {
	Name  string
	Email string
}

type UpdateUserPayload struct {
	ID    int
	Name  string
	Email string
}
