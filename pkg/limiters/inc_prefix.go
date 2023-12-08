package limiters

func incPrefix(p string) string {
	b := []byte(p)
	b[len(b)-1]++
	return string(b)
}
