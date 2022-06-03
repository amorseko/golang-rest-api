package simple

type SimpleRepository struct {
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

type SimpleServices struct {
	*SimpleRepository
}

func NewSimpleServices(repository *SimpleRepository) *SimpleServices {
	return &SimpleServices{
		SimpleRepository: repository,
	}
}
