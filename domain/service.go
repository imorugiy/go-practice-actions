package domain

type Service interface {
	Find(name string) (*Metadata, error)
	Store() error
}

type Repository interface {
	Find(name string) (*Metadata, error)
	Store() error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Find(name string) (*Metadata, error) {
	return s.repo.Find(name)
}

func (s *service) Store() error {
	return s.repo.Store()
}
