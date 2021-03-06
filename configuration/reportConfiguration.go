package configuration

import (
	"fmt"
	"genx-go/logger"
)

//ReportConfiguration represents report config
type ReportConfiguration struct {
	Fields []Field
}

//ConstructReportConfiguration create report config instance
func ConstructReportConfiguration(provider IReportConfigProvider) *ReportConfiguration {
	fields, err := provider.Provide()
	if err != nil {
		logger.Logger().WriteToLog(logger.Fatal, "[ReportConfiguration | ConstructReportConfiguration] Error while constructing report configuration. Error: ", err)
	}
	configuration := &ReportConfiguration{
		Fields: fields,
	}
	return configuration
}

//GetFieldByID returns description for field by id
func (reportConfiguration *ReportConfiguration) GetFieldByID(id string) (*Field, error) {
	for _, reportField := range reportConfiguration.Fields {
		if reportField.ID == id {
			return &reportField, nil
		}
	}
	return nil, fmt.Errorf("Not found field with id:%v", id)
}

//GetFieldsByIds returns fields array by ids
func (reportConfiguration *ReportConfiguration) GetFieldsByIds(ids []string) []*Field {
	result := make([]*Field, 0)
	for _, id := range ids {
		if reportField, err := reportConfiguration.GetFieldByID(id); err == nil {
			result = append(result, reportField)
		} else {
			logger.Logger().WriteToLog(logger.Error, "[GetReportColumnsByIds] ", err)
		}
	}
	return result
}
