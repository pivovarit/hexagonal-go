package account

type AccountService struct {
	accountRepository        AccountRepository
	accountHistoryRepository AccountHistoryRepository
}

type AccountRepository interface {
}

type AccountHistoryRepository interface {
}
