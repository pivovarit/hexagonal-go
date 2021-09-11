package account

type Service struct {
	accountRepository repository
	roleRepository    roleRepository
}

type repository interface {
	save(*account) (bool, error)
	findOne(id int) (*account, error)
	findAllByRole(role string) ([]*account, error)
}

type roleRepository interface {
	save(id int, role role) (bool, error)
	findOne(id int) (role, error)
}

func (s *Service) FindOne(id int) (*Account, error) {
	account, err := s.accountRepository.findOne(id)
	if err != nil {
		return nil, err
	}

	role, err := s.roleRepository.findOne(account.id)
	if err != nil {
		return nil, err
	}

	return ConvertToPublic(account, role), nil
}
