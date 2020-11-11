package message

//ParametersMessage represent parameters message
type ParametersMessage struct {
	Identity    string
	MessageType string
	Parameters  map[string]string
}
