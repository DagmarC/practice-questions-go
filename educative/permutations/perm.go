package permutations

import "fmt"

func Perm(word []rune) {
	perm(word, 0)
}

// perm function that accepts a slice or string and prints
// all possible combinations of characters
func perm(w []rune, i int) {
	if i >= len(w) {
		fmt.Println(string(w))
		return
	}
	for j := i; j < len(w); j++ {
		w[i], w[j] = w[j], w[i]
		perm(w, i+1)
		w[i], w[j] = w[j], w[i]
	}
}
