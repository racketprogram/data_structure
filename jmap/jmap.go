package jmap

type jmap struct {
	s           []*kv
	h           *hash
	valuesCount int
}

type kv struct {
	k string
	v string
}

func NewJmap(size int) (*jmap, error) {
	h, err := NewHash(size)
	if err != nil {
		return nil, err
	}

	return &jmap{
		s: make([]*kv, size),
		h: h,
	}, nil
}

func (m *jmap) Set(key, value string) {
	if m.valuesCount > len(m.s)-1 {
		panic("full")
	}
	hashValue := m.h.stringHash(key)
	for {
		if m.s[hashValue] == nil {
			m.s[hashValue] = &kv{k: key, v: value}
			m.valuesCount++
			return
		}
		hashValue++
		if hashValue > len(m.s)-1 {
			hashValue = 0
		}
	}
}

func (m *jmap) Get(key string) (value string, ok bool) {
	hashValue := m.h.stringHash(key)
	for {
		v := m.s[hashValue]
		if v != nil && v.k == key {
			return v.v, true
		}
		hashValue++
		if hashValue > len(m.s)-1 {
			hashValue = 0
		}
	}
}
