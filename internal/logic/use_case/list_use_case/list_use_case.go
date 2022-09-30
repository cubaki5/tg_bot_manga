package list_use_case

type Database interface {
	List() (interface{}, error)
}

type ListUseCase struct {
	db Database
}

func NewGetUseCase(db Database) *ListUseCase {
	return &ListUseCase{db: db}
}

func (s *ListUseCase) List() (interface{}, error) {
	return nil, nil
}
