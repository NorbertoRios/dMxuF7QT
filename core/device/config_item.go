package device

import (
	"fmt"
	"strings"
	"time"
)

//ConfigItem represents configuration item
type ConfigItem struct {
	Name    string
	Value   string
	SendtAt time.Time
	State   byte //0-Created; 1-Sended; 2-Acked;
}

//Parameter returns config for device
func (item *ConfigItem) Parameter() string {
	if strings.ToUpper(item.Name) == "SETBOUNDARY" {
		return fmt.Sprintf("%vBACKUPNVRAM;", item.Value)
	}
	return fmt.Sprintf("SETPARAM;%vENDPARAM;BACKUPNVRAM;", item.Value)
}
