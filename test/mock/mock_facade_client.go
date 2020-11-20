package mock

import (
	"genx-go/core/interfaces"
	"genx-go/types"
)

//FacadeClient ...
type FacadeClient struct {
}

//Execute ...
func (client *FacadeClient) Execute(request interfaces.IRequest) interface{} {
	cfg := []string{
		"4=FFFour;",
		"5=FFFive;",
		"6=SSSix;",
		"7=SSSeven;",
		"8=EEEight;",
		"9=NNNine;",
		"10=TTTen;",
	}
	sType := types.StringArray{Data: cfg}
	return sType.List()
}
