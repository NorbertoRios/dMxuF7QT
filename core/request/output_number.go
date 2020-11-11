package request

import (
	"fmt"
	"genx-go/logger"
)

//OutputNumber output number
type OutputNumber struct {
	Data string
}

//Index returns outout number
func (out *OutputNumber) Index() int {
	switch out.Data {
	case "OUT0":
		return 1
	case "OUT1":
		return 2
	case "OUT2":
		return 3
	case "OUT3":
		return 4
	case "OUT4":
		return 5
	default:
		{
			logger.Logger().WriteToLog(logger.Error, fmt.Sprintf("[OutputNumber | Command] Unexpected output value. Incoming port:%v", out.Data))
			return -1
		}
	}
}
