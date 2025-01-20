package example

type TypeA struct {
	Name    string
	Surname string
	Address string
}

func example() {
	var leak1 = map[string]TypeA{}
	_ = leak1

	var ok1 = map[string]int{}
	var ok2 = map[string]*TypeA{}
	var ok3 = map[string]struct{}{}

	_ = ok1
	_ = ok2
	_ = ok3
}
