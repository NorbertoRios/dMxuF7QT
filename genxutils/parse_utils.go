package genxutils

func invalidateSpeed(value interface{}) interface{} {
	if v, valid := value.(int32); valid {
		return float32(float32(v) / 3.6)
	}
	return value
}
func invalidateCoordinate(value interface{}) interface{} {
	if v, valid := value.(int32); valid {
		return float32(v / 3600000.0)
	}
	return value
}
func invalidateTenth(value interface{}) interface{} {
	if v, valid := value.(int32); valid {
		return int32(v * 10)
	}
	return value
}
func invalidateHundredth(value interface{}) interface{} {
	if v, valid := value.(int32); valid {
		return float32(v / 100.0)
	}
	return value
}
func invalidateHours(value interface{}) interface{} {
	if v, valid := value.(int32); valid {
		return int32(v * 36)
	}
	return value
}
func invalidateIO(value interface{}) interface{} { //switches and relay
	if v, valid := value.(byte); valid {
		return v
	}
	return value
}
