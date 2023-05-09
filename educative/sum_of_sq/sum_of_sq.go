package sumofsq

import (
	"sync"
)

//	SumOfSquares function which takes an integer, c and
// returns the sum of all squares between 1 and c.
// You’ll need to use select statements, goroutines, and channels.
func SumOfSquares(c int) int {
	var total int
	var wg sync.WaitGroup

	sum := make(chan int, c)
	done := make(chan struct{})

	wg.Add(c)
	for i := 1; i <= c; i++ {
		go func(i int) {
			sum <- i * i
			wg.Done()
		}(i)
	}

	go func(wg *sync.WaitGroup, done chan<- struct{}) {
		wg.Wait()
		close(done) // sends done signal
	}(&wg, done)

loop:
	for {
		select {
		case s := <-sum:
			total += s
		case <-done:
			close(sum) // it is better to close channel on the sending side, but here I am sure no more sendings will ever occur
			break loop
		}
	}
	return total
}
