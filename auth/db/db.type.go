package authDB

type CreateUserParams struct {
	Username string
	Password string
	Email    string
	Role     string
	IsActive bool
}
