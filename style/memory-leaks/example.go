package example

type TypeA struct {
	Name    string
	Surname string
	Address string
}

var leak1 = map[string]TypeA{}

var ok1 = map[string]int{}
var ok2 = map[string]*TypeA{}
var ok3 = map[string]struct{}{}
