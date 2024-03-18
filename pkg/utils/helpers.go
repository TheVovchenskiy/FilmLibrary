package utils

func In(elem any, elemets ...any) bool {
	for _, element := range elemets {
		if element == elem {
			return true
		}
	}
	return false
}
