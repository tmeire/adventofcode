package collection

type StringSet map[string]struct{}

func NewStringSet() StringSet {
	return make(StringSet)
}

func (s StringSet) Contains(x string) bool {
	_, ok := s[x]
	return ok
}

func (s StringSet) Put(x string) bool {
	if _, ok := s[x]; ok {
		return false
	}

	s[x] = struct{}{}
	return true
}

func (s StringSet) Delete(x string) {
	delete(s, x)
}
