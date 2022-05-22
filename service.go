package citizen

//service need repo to works
type Service struct {
	repo ProfileRepo
}

//instantiate Default Service
func Default(pr ProfileRepo) *Service {
	return &Service{repo: pr}
}

//Find will Get All Profiles data as JSON Format
func (s Service) Find() []byte {
	return s.repo.Find()
}
