package templates_repository

func (repo *TemplateRepository) GetTemplate(id int) ([]byte, bool) {
	template, exists := repo.templates[id]
	if !exists {
		return nil, exists
	}

	return template, true
}
