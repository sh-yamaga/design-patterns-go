package bridge

type Device interface {
	On() string
	Off() string
}

type RemoteController struct {
	device Device
}

func NewRemoteController(device Device) *RemoteController {
	return &RemoteController{device: device}
}

func (rc RemoteController) TurnOn() string {
	return rc.device.On()
}

func (rc RemoteController) TurnOff() string {
	return rc.device.Off()
}
