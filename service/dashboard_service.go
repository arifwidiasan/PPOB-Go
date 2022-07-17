package service

func (s *svc) GetCountTransactionService() (int64, error) {
	return s.repo.GetCountTransaction()
}

func (s *svc) GetCountTransactionSuccessService() (int64, error) {
	return s.repo.GetCountTransactionSuccess()
}

func (s *svc) GetCountProductService() (int64, error) {
	return s.repo.GetCountProduct()
}

func (s *svc) GetCountUserService() (int64, error) {
	return s.repo.GetCountUser()
}
