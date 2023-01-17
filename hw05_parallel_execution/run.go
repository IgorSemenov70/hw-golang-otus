package hw05parallelexecution

import (
	"errors"
	"sync"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrEmptyArrayTasks     = errors.New("tasks array is empty")
)

type Task func() error

func Run(tasks []Task, n, m int) error {
	if len(tasks) == 0 {
		return ErrEmptyArrayTasks
	}
	wg := &sync.WaitGroup{}
	taskChannel := make(chan Task)
	errChannel := make(chan error)
	done := make(chan struct{})
	errCount := 0

	go func(done <-chan struct{}, wg *sync.WaitGroup) {
		for _, err := range tasks {
			select {
			case taskChannel <- err:
			case <-done:
				break
			}
		}
		close(taskChannel)
		wg.Wait()
		close(errChannel)
	}(done, wg)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(done <-chan struct{}, wg *sync.WaitGroup) {
			defer wg.Done()

			for task := range taskChannel {
				select {
				case errChannel <- task():
				case <-done:
					return
				}
			}
		}(done, wg)
	}

	for err := range errChannel {
		if err != nil {
			errCount++
		}
		if errCount == m && m > 0 {
			close(done)
			wg.Wait()
			return ErrErrorsLimitExceeded
		}
	}
	return nil
}
