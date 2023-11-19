package interfaces

type Controller interface {
	RegisterDeviceIds(deviceIds []string)
	ListRegisteredDeviceIds()
	TriggerAlarmForAllDetectors()
}

type ControllerData struct {
	DeviceIds []string
}
