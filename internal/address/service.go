package address

type DatabaseLoader interface {
	Load() (zipDatabase ZipDatabase, err0 error)
}

type Service struct {
	ZipDatabase ZipDatabase
}

func NewService(databaseLoader *JsonZipDatabaseLoader) (*Service, error) {
	zipDatabase, err := databaseLoader.Load()
	if err != nil {
		return nil, err
	}
	return &Service{
		ZipDatabase: zipDatabase,
	}, nil
}

func (s *Service) Fetch(zipCode string) (*Address, bool) {
	address, ok := s.ZipDatabase[zipCode]

	return &address, ok
}
