package configuration

import (
	"log"
	"io/ioutil"
	"encoding/xml"
	"genx-go/genxutils"
)

type xmlField struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
	Size int    `xml:"size,attr"`
}

type xmlReportConfiguration struct {
	Fields []xmlField `xml:"Fields>Field"`
}

type XmlProvider struct {
  file genxutils.IFile
}

func ConstructXmlProvider(file genxutils.IFile) (*XmlProvider){
   return &XmlProvider{
     file: file,
   }
}

func (provider *XmlProvider) Provide() ([]Field, error){
	//file := genxutils.FileUtils{Filename: provider.fileName}
	filePath := provider.file.Path()
	log.Println("Loading report configuration from:", filePath)
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
  for _, xmlField := range xmlConfiguration.Fields{
      field := Field{
        ID: xmlField.ID,
        Name: xmlField.Name,
        Size: xmlField.Size,
      }
      result = append(result, field)
  }
	return result, nil
}
