package device

import (
	"genx-go/core/sensors"
	"genx-go/logger"
	"sync"
	"time"
)

//NewSensorState ...
func NewSensorState(_lastState *State, deviceSensors []sensors.ISensor) *State {
	hash := _lastState.State()
	for _, sensor := range deviceSensors {
		if sensor.Symbol() == "" {
			logger.Logger().WriteToLog(logger.Error, "[NewSensorState] Cant find symbol for sensor. Sensor: ", sensor.ToDTO())
			continue
		}
		hash[sensor.Symbol()] = sensor
	}
	return NewState(hash)
}

//NewState ...
func NewState(deviceSensors map[string]sensors.ISensor) *State {
	return &State{
		mutex:         &sync.Mutex{},
		deviceSensors: deviceSensors,
		updateTime:    time.Now().UTC(),
	}
}

//State ...
type State struct {
	mutex         *sync.Mutex
	deviceSensors map[string]sensors.ISensor
	updateTime    time.Time
}

//State ...
func (s *State) State() map[string]sensors.ISensor {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.deviceSensors
}
