package account

type Account struct {
	Id   int
	Name string
	Role string
}

type CreateAccountRequest struct {
	Id   int
	Name string
	Role string
}
