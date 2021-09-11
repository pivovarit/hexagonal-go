package account

func ConvertToPublic(acc *account) *Account {
	return &Account{
		Id:      acc.id,
		Name:    acc.name,
		Surname: acc.surname,
	}
}

func ConvertToDomain(acc *Account) *account {
	return &account{
		id:      acc.Id,
		name:    acc.Name,
		surname: acc.Surname,
	}
}
