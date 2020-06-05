package configuration

import (
	"fmt"
	"log"
)

type IProvider interface {
	Provide() ([]Field, error)
}

type ReportConfiguration struct {
	Fields []Field
}

//ConstructReportConfiguration create report config instance
func ConstructReportConfiguration(provider IProvider) (*ReportConfiguration, error) {
	fields, err := provider.Provide()
	if err != nil {
		return nil, err
	}
	configuration := &ReportConfiguration{
		Fields: fields,
	}
	return configuration, nil
}

//GetField returns description for field by id
func (reportConfiguration *ReportConfiguration) GetFieldById(id string) (*Field, error) {
	for _, reportField := range reportConfiguration.Fields {
		if reportField.ID == id {
			return &reportField, nil
		}
	}
	return nil, fmt.Errorf("Not found field with id:%v", id)
}

func (reportConfiguration *ReportConfiguration) GetFieldsByIds(ids []string) []*Field {
	result := make([]*Field, 0)
	for _, id := range ids {
		if reportField, err := reportConfiguration.GetFieldById(id); err == nil {
			result = append(result, reportField)
		} else {
			log.Println("[GetReportColumnsByIds] ", err)
		}
	}
	return result
}
