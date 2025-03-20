package bridge

type Radio struct{}

func (r *Radio) On() string {
	return "Radio is turned on"
}

func (r *Radio) Off() string {
	return "Radio is turned off"
}
