package request

import "strings"

//BaseRequest ...
type BaseRequest struct {
	Identity         string `json:"Identity"`
	FacadeCallbackID string `json:"FacadeCallbackID"`
	TTL              int    `json:"TTL"`
}

//Serial ...
func (data *BaseRequest) Serial() string {
	return strings.Replace(data.Identity, "genx_", "", 1)
}

//CallbackID ...
func (data *BaseRequest) CallbackID() string {
	return data.FacadeCallbackID
}
