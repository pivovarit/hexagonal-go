package account

func ConvertToPublic(acc *account, role role) *Account {
	return &Account{
		Id:      acc.id,
		Name:    acc.name,
		Role:    string(role),
	}
}

func ConvertToDomain(acc *Account) *account {
	return &account{
		id:      acc.Id,
		name:    acc.Name,
	}
}
