package print_form_service

import (
	"context"
	"errors"
	"log"
)

const reportTemplateID = 1

func (svc *PrintFormService) CreateReport(ctx context.Context, request CreateReportRequest) ([]byte, error) {
	template, exists := svc.templatesRepository.GetTemplate(reportTemplateID)
	if !exists {
		err := errors.New("cannot found report template")

		log.Println(err)

		return nil, err
	}

	var report Report

	reportBytes, err := svc.makeReport(template, report)
	if err != nil {
		return nil, err
	}

	return reportBytes, nil
}

func (svc *PrintFormService) makeReport(template []byte, report Report) ([]byte, error) {
	return nil, nil
}
