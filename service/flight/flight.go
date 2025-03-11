package flight

type Flight []string

func (f Flight) valid() bool {
	return true
}

func (f Flight) source() string {
	return ""
}

func (f Flight) destination() string {
	return ""
}
