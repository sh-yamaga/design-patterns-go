package decorator

type IStream interface {
	Read() string
}

type Stream struct {
	Data string
}

func (s *Stream) Read() string {
	return s.Data
}
