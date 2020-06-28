package configuration

import (
	"encoding/xml"
	"genx-go/logger"
	"genx-go/utils"
	"io/ioutil"
)

type xmlField struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
	Size int    `xml:"size,attr"`
}

type xmlReportConfiguration struct {
	Fields []xmlField `xml:"Fields>Field"`
}

//XMLProvider represent provider for XMl
type XMLProvider struct {
	file utils.IFile
}

//ConstructXMLProvider returns xml provider
func ConstructXMLProvider(file utils.IFile) *XMLProvider {
	return &XMLProvider{
		file: file,
	}
}

//Provide provide
func (provider *XMLProvider) Provide() ([]Field, error) {
	filePath := provider.file.Path()
	logger.Info("Loading report configuration from:", filePath)
	xmlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	xmlConfiguration := &xmlReportConfiguration{}
	err = xml.Unmarshal(xmlFile, xmlConfiguration)
	if err != nil {
		return nil, err
	}

	result := make([]Field, 0)
	for _, xmlField := range xmlConfiguration.Fields {
		field := Field{
			ID:   xmlField.ID,
			Name: xmlField.Name,
			Size: xmlField.Size,
		}
		result = append(result, field)
	}
	return result, nil
}
