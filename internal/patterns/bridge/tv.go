package bridge

type TV struct{}

func (t *TV) On() string {
	return "TV is turned on"
}

func (t *TV) Off() string {
	return "TV is turned off"
}
