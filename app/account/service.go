package account

type Service struct {
	accountRepository        repository
	accountHistoryRepository historyRepository
}

type repository interface {
	FindOne(id int) (*account, error)
}

type historyRepository interface {
	FindOne(id int) (*history, error)
}

func (s *Service) FindOne(id int) (*Account, error) {
	result, err := s.accountRepository.FindOne(id)
	if err != nil {
		return nil, err
	}
	return ConvertToPublic(result), nil
}
