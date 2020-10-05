package dude

type SayMyName struct {
	Name string
}

func (s SayMyName) SayIt() string {
	return s.Name
}
