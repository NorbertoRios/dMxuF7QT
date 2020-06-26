package configuration

import (
	"encoding/json"
	"genx-go/utils"
	"io/ioutil"
	"log"
)

//CredentialsJSONProvider provider
type CredentialsJSONProvider struct {
	file utils.IFile
}

//ConstructCredentialsJSONProvider returns CredentialsJSONProvider
func ConstructCredentialsJSONProvider(file utils.IFile) *CredentialsJSONProvider {
	return &CredentialsJSONProvider{
		file: file,
	}
}

//ProvideCredentials provide
func (provider *XMLProvider) ProvideCredentials() (*ServiceCredentials, error) {
	filePath := provider.file.Path()
	log.Println("Loading credentials configuration from:", filePath)
	xmlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	credentials := &ServiceCredentials{}
	err = json.Unmarshal(xmlFile, credentials)
	if err != nil {
		return nil, err
	}
	return credentials, nil
}
