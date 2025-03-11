package flight

type Flight []string

func (f Flight) valid() bool {
	return len(f) == 2
}

func (f Flight) source() string {
	if !f.valid() {
		return ""
	}

	return f[0]
}

func (f Flight) destination() string {
	if !f.valid() {
		return ""
	}

	return f[1]
}
