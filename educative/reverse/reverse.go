package reverse

// Reverse that takes a slice of any and reverses the slice in place without using a temporary slice
func Reverse[T any](s *[]T) {
	for i := 0; i < len(*s)/2; i++ {
		(*s)[i], (*s)[len(*s)-1-i] = (*s)[len(*s)-1-i], (*s)[i]
	}
}
