package configuration

import (
	"container/list"
	"encoding/json"
	"genx-go/logger"
	"io/ioutil"
	"net/http"
	"time"
)

//NewFacadeClient ...
func NewFacadeClient() *FacadeClient {
	return &FacadeClient{
		client: http.Client{
			Timeout: time.Duration(30 * time.Second),
		},
	}
}

//FacadeClient ...
type FacadeClient struct {
	client http.Client
}

//Execute ..
func (facadeClient *FacadeClient) Execute(_request *FacadeRequest) interface{} {
	configs := list.New()
	response, err := facadeClient.client.Do(_request.Request())
	defer response.Body.Close()
	if err != nil {
		logger.Logger().WriteToLog(logger.Error, "[FacadeClient | Execute] Error while executing request to facade service. Error: ", err)
		return configs
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Logger().WriteToLog(logger.Error, "[FacadeClient | Execute] Error while reading byte response. Error: ", err)
		return configs
	}
	err = json.Unmarshal(body, &configs)
	if err != nil {
		logger.Logger().WriteToLog(logger.Error, "[FacadeClient | Execute] Error while unmarshaling byte response to string array. Error: ", err)
		return configs
	}
	return configs
}
