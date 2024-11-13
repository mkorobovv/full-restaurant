package templates_repository

import (
	"log"
	"os"
)

type TemplateRepository struct {
	templates map[int][]byte
}

func New() *TemplateRepository {
	templates, err := initTemplates()
	if err != nil {
		log.Fatal(err.Error())
	}

	return &TemplateRepository{
		templates: templates,
	}
}

const reportTemplateID = 1

func initTemplates() (map[int][]byte, error) {

	reportTemplate, err := os.ReadFile("internal/app/adapters/secondary/repositories/templates-repository/templates/report.xlsx")
	if err != nil {
		return nil, err
	}

	templates := make(map[int][]byte)

	templates[reportTemplateID] = reportTemplate

	return templates, nil
}
