package decorator

type Interface interface {
	Read() string
}

type Stream struct {
	Data string
}

func (s *Stream) Read() string {
	return s.Data
}
