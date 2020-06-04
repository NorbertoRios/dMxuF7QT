package configuration

import (
	"encoding/xml"
	"fmt"
	"genx-go/genxutils"
	"io/ioutil"
	"log"
)

//ReportConfiguration root of report configuration
type ReportConfiguration struct {
	Fields []Field `xml:"Fields>Field"`
}

//ConstructReportConfiguration create report config instance
func ConstructReportConfiguration(fileName string) (*ReportConfiguration, error) {
	file := genxutils.FileUtils{Filename: fileName}
	filePath := file.Path()
	log.Println("Loading report configuration from:", filePath)
	configXML, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	configInstance := &ReportConfiguration{}
	err = xml.Unmarshal(configXML, configInstance)
	if err != nil {
		return nil, err
	}
	return configInstance, nil
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
	for _, id := range ids{
		if reportField, err := reportConfiguration.GetFieldById(id); err == nil {
			result = append(result, reportField)
		} else {
			log.Println("[GetReportColumnsByIds] ", err)
		}
	}
	return result
}
