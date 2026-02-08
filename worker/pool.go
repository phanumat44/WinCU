package worker

import (
	"sync"
)

// Task represents a unit of work
type Task func()

// Pool represents a worker pool
type Pool struct {
	workers int
	taskCh  chan Task
	wg      sync.WaitGroup
}

// NewPool creates a new worker pool
func NewPool(workers int) *Pool {
	return &Pool{
		workers: workers,
		taskCh:  make(chan Task),
	}
}

// Start initializes the workers
func (p *Pool) Start() {
	for i := 0; i < p.workers; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for task := range p.taskCh {
				task()
			}
		}()
	}
}

// Submit adds a task to the pool
func (p *Pool) Submit(task Task) {
	p.taskCh <- task
}

// Stop closes the task channel and waits for all workers to finish
func (p *Pool) Stop() {
	close(p.taskCh)
	p.wg.Wait()
}
