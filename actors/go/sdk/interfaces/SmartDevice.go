package interfaces

import "fmt"

type SmartDevice interface {
	SetData(data SmartDeviceData)
	GetData() SmartDeviceData
	DetectSmoke()
	SoundAlarm()
	ClearAlarm()
}

type SmartDeviceData struct {
	status   string
	location string
}

// Override the default string method
func (data SmartDeviceData) String() string {
	return fmt.Sprintf("location: {%s}, Status: {%s}", data.location, data.status)
}
