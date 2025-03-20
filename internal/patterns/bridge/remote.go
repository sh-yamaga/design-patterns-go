package bridge

type RemoteController struct {
	device IDevice
}

func NewRemoteController(device IDevice) *RemoteController {
	return &RemoteController{device: device}
}

func (rc RemoteController) TurnOn() string {
	return rc.device.On()
}

func (rc RemoteController) TurnOff() string {
	return rc.device.Off()
}
