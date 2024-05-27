package utils

import "sync"

type ThreadManager struct {
	wg sync.WaitGroup
}

func (tm *ThreadManager) Go(f func()) {
	tm.wg.Add(1)
	go func() {
		defer tm.wg.Done()
		f()
	}()
}

func (tm *ThreadManager) Wait() {
	tm.wg.Wait()
}
