package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {

	if n < 1 {
		n = 1
	}

	if m <= 0 {
		m = 1
	}

	taskCh := make(chan Task)

	var errCounter int32
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range taskCh {
				if err := task(); err != nil {
					atomic.AddInt32(&errCounter, 1)
				}
			}
		}()
	}

	for _, task := range tasks {
		if atomic.LoadInt32(&errCounter) >= int32(m) {
			break
		}
		taskCh <- task
	}
	close(taskCh)
	wg.Wait()

	if errCounter >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
