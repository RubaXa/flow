package flow

import "sync"

func CreateWorkerPool(size int, worker func()) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(size)

	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			worker()
		}()
	}

	return wg
}
