package utils

func Include[T comparable](list []T, test T) bool {
	for _, item := range list {
		if item == test {
			return true
		}
	}
	return false
}
