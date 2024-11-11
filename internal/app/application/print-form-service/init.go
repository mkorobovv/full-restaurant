package print_form_service

type PrintFormService struct {
	printFormsRepository printFormsRepository
}

type printFormsRepository interface {
}

func New(repo *printFormsRepository) *PrintFormService {
	return &PrintFormService{
		printFormsRepository: repo,
	}
}
