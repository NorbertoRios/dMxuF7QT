package adaptors

import (
	"encoding/json"
	"genx-go/adaptors/dto"
	"genx-go/adaptors/observers"
	"genx-go/core/sensors"
	"genx-go/logger"
	"genx-go/repository/models"
)

//NewDeviceActivity ...
func NewDeviceActivity(_activity *models.DeviceActivity) *DeviceActivity {
	adaptorObservers := []observers.IAdaptorObserver{}
	adaptorObservers = append(adaptorObservers, observers.NewDTCCodeObserver())
	adaptorObservers = append(adaptorObservers, observers.NewFuelObserver())
	adaptorObservers = append(adaptorObservers, observers.NewGPSObserver())
	adaptorObservers = append(adaptorObservers, observers.NewIBIDObserver())
	adaptorObservers = append(adaptorObservers, observers.NewIgnitionObserver())
	adaptorObservers = append(adaptorObservers, observers.NewInputsObserver())
	adaptorObservers = append(adaptorObservers, observers.NewNetworkObserver())
	adaptorObservers = append(adaptorObservers, observers.NewOutputsObserver())
	adaptorObservers = append(adaptorObservers, observers.NewPowerSensorObserver())
	adaptorObservers = append(adaptorObservers, observers.NewQueueObserver())
	adaptorObservers = append(adaptorObservers, observers.NewTemperatureSensorsObserver())
	adaptorObservers = append(adaptorObservers, observers.NewTimeObserver())
	adaptorObservers = append(adaptorObservers, observers.NewTripObserver())
	adaptorObservers = append(adaptorObservers, observers.NewVINObserver())
	message := &dto.DtoMessage{}
	err := json.Unmarshal([]byte(_activity.LastMessage), message)
	if err != nil {
		logger.Logger().WriteToLog(logger.Error, "[DeviceActivity | Adapt] Error while unmarshaling last known device state. Error: ", err, ". LastActivity : ", _activity)
		return nil
	}
	return &DeviceActivity{
		activity:         _activity,
		adaptorObservers: adaptorObservers,
		message:          message,
	}
}

//DeviceActivity ...
type DeviceActivity struct {
	activity         *models.DeviceActivity
	adaptorObservers []observers.IAdaptorObserver
	message          *dto.DtoMessage
}

//DTO ...
func (da *DeviceActivity) DTO() *dto.DtoMessage {
	return da.message
}

//Adapt ...
func (da *DeviceActivity) Adapt() map[string]sensors.ISensor {
	_hash := da.adaptLastMessageToSensors()
	fw := da.adaptSoftware()
	_hash[fw.Symbol()] = fw
	return _hash
}

func (da *DeviceActivity) adaptLastMessageToSensors() map[string]sensors.ISensor {
	_hash := make(map[string]sensors.ISensor)
	for _, observer := range da.adaptorObservers {
		if sensor := observer.Notify(da.message); sensor != nil {
			_hash[sensor.Symbol()] = sensor
		}
	}
	return _hash
}

func (da *DeviceActivity) adaptSoftware() sensors.ISensor {
	return sensors.BuildFirmwareSensor(da.activity.Serializedsoftware)
}
