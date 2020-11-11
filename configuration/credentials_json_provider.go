package configuration

import (
	"encoding/json"
	"genx-go/logger"
	"genx-go/types"
	"io/ioutil"
)

//CredentialsJSONProvider provider
type CredentialsJSONProvider struct {
	file types.IFile
}

//ConstructCredentialsJSONProvider returns CredentialsJSONProvider
func ConstructCredentialsJSONProvider(file types.IFile) *CredentialsJSONProvider {
	return &CredentialsJSONProvider{
		file: file,
	}
}

//ProvideCredentials provide
func (provider *XMLProvider) ProvideCredentials() (*ServiceCredentials, error) {
	filePath := provider.file.Path()
	logger.Logger().WriteToLog(logger.Info, "Loading credentials configuration from:", filePath)
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
