package model

func appendIfNotExists(slice []string, element string) []string {
	for _, s := range slice {
		if s == element {
			return slice
		}
	}
	return append(slice, element)
}
