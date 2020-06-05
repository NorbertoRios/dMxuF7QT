package configuration

import (
	"encoding/xml"
	"fmt"
	"genx-go/utils"
	"io/ioutil"
	"log"
)

//IReportConfiguration ReportConfiguration interface
type IReportConfiguration interface {
	GetField(id string) (*Field, error)
}

//ReportConfiguration root of report configuration
type ReportConfiguration struct {
	Fields []Field `xml:"Fields>Field"`
}

//GetField returns description for field by id
func (c *ReportConfiguration) GetField(id string) (*Field, error) {
	for _, f := range c.Fields {
		if f.ID == id {
			return &f, nil
		}
	}
	return nil, fmt.Errorf("Not found field with id:%v", id)
}

//ConstructReportConfiguration create report config instance
func ConstructReportConfiguration(fileName string) (IReportConfiguration, error) {
	file := utils.FileUtils{Filename: fileName}
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
