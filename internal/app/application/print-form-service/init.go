package print_form_service

type PrintFormService struct {
	printFormsRepository printFormsRepository
	templatesRepository  templatesRepository
}

type printFormsRepository interface {
}

type templatesRepository interface {
	GetTemplate(id int64) ([]byte, bool)
}

func New(repo printFormsRepository) *PrintFormService {
	return &PrintFormService{
		printFormsRepository: repo,
	}
}
