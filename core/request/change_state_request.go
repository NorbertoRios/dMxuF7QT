package request

//ChangeStateRequest ...
type ChangeStateRequest struct {
	BaseRequest
	Port string `json:"PortNumber"`
}
