package adaptors

import (
	"genx-go/adaptors"
	"genx-go/repository/models"
	"testing"
)

func TestDTOAdaptorWithTemperatureSensor(t *testing.T) {
	model := &models.DeviceActivity{
		Identity:           "genx_000000804007",
		LastMessage:        "{\"sid\":77731078,\"Data\":{\"DevId\":\"genx_000000804031\",\"BatteryPercentage\":null,\"BatteryStatus\":null,\"FuelConsumption\":null,\"FuelLevel\":null,\"GPIO\":0,\"Reason\":2,\"Heading\":0.0,\"Latitude\":49.1666325,\"Longitude\":-122.98513277777778,\"Altitude\":2.0,\"Speed\":0.0,\"GpsTimeStamp\":\"2014-04-17T15:27:06Z\",\"TimeStamp\":\"2021-01-15T12:23:22.305533Z\",\"GpsValidity\":1,\"IBID\":0,\"IgnitionState\":0,\"Odometer\":6139109,\"Supply\":13184,\"Relay\":0,\"Rpm\":null,\"ReceivedTime\":\"2021-01-15T12:23:22.305533Z\",\"Satellites\":0,\"storage\":\"cloud\",\"PrevSourceId\":77731075},\"ts\":{\"Sensor1\":{\"Id\":\"9999999999999955\",\"TemperatureValue\":10.25,\"Name\":null,\"Event\":0,\"Events\":[],\"TemperatureThreshold\":0.0,\"LogId\":0},\"Sensor2\":{\"Id\":\"9999999999999956\",\"TemperatureValue\":9.56,\"Name\":null,\"Event\":0,\"Events\":[],\"TemperatureThreshold\":0.0,\"LogId\":0},\"Sensor3\":{\"Id\":\"0000000000000000\",\"TemperatureValue\":-255.0,\"Name\":null,\"Event\":0,\"Events\":[],\"TemperatureThreshold\":0.0,\"LogId\":0},\"Sensor4\":{\"Id\":\"0000000000000000\",\"TemperatureValue\":-255.0,\"Name\":null,\"Event\":0,\"Events\":[],\"TemperatureThreshold\":0.0,\"LogId\":0}},\"time\":\"2021-01-15T12:23:22.305533Z\",\"Events\":[]}",
		LastMessageID:      uint64(77731078),
		Serializedsoftware: "",
	}
	adaptor := adaptors.NewDeviceActivity(model)
	if adaptor == nil {
		t.Error("Adaptor is nil. Check logs")
	}
	adaptedSensors := adaptor.Adapt()
	if adaptedSensors == nil {
		t.Error("")
	}
}

func TestDtoAdaptorToSensors(t *testing.T) {
	model := &models.DeviceActivity{
		Identity:           "genx_000003870006",
		LastMessage:        "{\"DevId\":\"genx_000003870006\",\"Data\":{\"DevId\":\"genx_000003870006\",\"GpsValidity\":0,\"IgnitionState\":1,\"FMI\":\"Off\",\"PowerState\":\"Powered\",\"PrevSourceId\":77905300,\"ReceivedTime\":\"2021-01-14T13:56:28Z\",\"storage\":\"cloud\",\"Latitude\":48.74648666666667,\"Longitude\":37.590804444444444,\"Speed\":1.67,\"Heading\":47,\"TimeStamp\":\"2021-01-14T13:56:28Z\",\"Odometer\":349787,\"Reason\":6,\"LocId\":59772,\"Relay\":4,\"Supply\":12015,\"CSID\":25506,\"RSSI\":-57,\"Satellites\":0,\"GPSStat\":16,\"GPIO\":0}}",
		LastMessageID:      uint64(77905380),
		Serializedsoftware: "G602.06.80kX",
	}
	adaptor := adaptors.NewDeviceActivity(model)
	adaptedSensors := adaptor.Adapt()
	if adaptedSensors == nil || len(adaptedSensors) == 0 {
		t.Error("Unexpected sensors value")
	}
	_templates := make(map[string]ITemplate)
	_templates["GPS"] = &GPSTemplate{Latitude: float32(48.74648666666667), Longitude: float32(37.590804444444444), Speed: float32(1.67), Satellites: byte(0), GpsValidity: byte(16), Heading: float32(47)}
	_templates["Ignition"] = &IgnitionTemplate{State: byte(1)}
	_templates["GPIO"] = &GPIOTemplate{Switch1: byte(0), Switch2: byte(0), Switch3: byte(0), Switch4: byte(0)}
	_templates["Relay"] = &RelaysTemplate{Relay1: byte(0), Relay2: byte(0), Relay3: byte(1), Relay4: byte(0)}
	_templates["Power"] = &PowerTemplate{Power: int32(12015), State: "Powered"}
	_templates["LocId"] = &LocIDTemplate{LockID: uint32(59772)}
	_templates["Time"] = &TimeTemplate{TimeStamp: "2021-01-14T13:56:28Z"}
	_templates["Odometer"] = &TripTemplate{Odometer: int32(349787)}
	_templates["Network"] = &NetworkTemplate{RSSI: int8(-57), CSID: int32(25506)}
	_templates["Firmware"] = &FirmwareTemplate{Version: "G602.06.80kX"}
	result := checkSensorsState(adaptedSensors, &Template{Templates: _templates}, t)
	if len(result) != 0 {
		t.Error("Adapted sensors count should be 0. Current: ", len(adaptedSensors))
	}
}
