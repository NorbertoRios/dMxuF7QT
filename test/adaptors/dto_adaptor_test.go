package adaptors

import (
	"genx-go/adaptors"
	"genx-go/repository/models"
	"testing"
)

func TestDtoAdaptorToSensors(t *testing.T) {
	model := &models.DeviceActivity{
		Identity:           "genx_000003870006",
		LastMessage:        "{\"DevId\":\"genx_000003870006\",\"Data\":{\"DevId\":\"genx_000003870006\",\"GpsValidity\":0,\"IgnitionState\":1,\"FMI\":\"Off\",\"PowerState\":\"Powered\",\"PrevSourceId\":77905300,\"ReceivedTime\":\"2021-01-14T13:56:28Z\",\"storage\":\"cloud\",\"Latitude\":48.74648666666667,\"Longitude\":37.590804444444444,\"Speed\":1.67,\"Heading\":47,\"TimeStamp\":\"2021-01-14T13:56:28Z\",\"Odometer\":349787,\"Reason\":6,\"LocId\":59772,\"Relay\":4,\"Supply\":12015,\"CSID\":25506,\"RSSI\":-57,\"Satellites\":0,\"GPSStat\":16,\"GPIO\":0}}",
		LastMessageID:      uint64(77905380),
		Serializedsoftware: "G602.06.80kX",
	}
	adaptor := adaptors.NewDeviceActivity(model)
	sensors := adaptor.Adapt()
	if sensors == nil {
		t.Error("Unexpected sensors value")
	}
}
