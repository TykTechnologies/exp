package example

type logger struct {
	out []string
}

func (l *logger) Log(s string) {
	l.out = append(l.out, s)
}
