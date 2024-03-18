package utils

func In[T comparable](elem T, elemets []T) bool {
	for _, element := range elemets {
		if element == elem {
			return true
		}
	}
	return false
}
