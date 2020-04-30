package jmap

type jmap struct {
	s []string
	h hash
}

func NewJmap(size int) (*jmap, error) {
	h, err := NewHash(size)
	if err != nil {
		return nil, err
	}

	return &jmap{
		s: make([]string, 0, size),
		h: h,
	}
}
