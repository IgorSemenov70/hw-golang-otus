package hw05parallelexecution

import (
	"errors"
	"sync"
)

var (
	ErrErrorsLimitExceeded       = errors.New("errors limit exceeded")
	ErrEmptyArrayTasks           = errors.New("tasks array is empty")
	ErrNumberOfWorkersCantBeZero = errors.New("the number of workers can't be 0")
)

type Task func() error

func Run(tasks []Task, n, m int) error {
	if len(tasks) == 0 {
		return ErrEmptyArrayTasks
	}
	if n <= 0 {
		return ErrNumberOfWorkersCantBeZero
	}
	wg := &sync.WaitGroup{}
	taskChannel := make(chan Task)
	errChannel := make(chan error)
	done := make(chan struct{})

	go func() {
		for _, err := range tasks {
			select {
			case <-done:
				break
			case taskChannel <- err:
			}
		}
		close(taskChannel)
		wg.Wait()
		close(errChannel)
	}()

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for task := range taskChannel {
				select {
				case <-done:
					return
				case errChannel <- task():
				}
			}
		}()
	}

	errCount := 0
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
