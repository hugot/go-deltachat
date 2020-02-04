package deltachat

import (
	"errors"
	"fmt"
	"time"
)

type worker struct {
	Name      string
	work      func()
	interrupt func()
	quit      chan bool
	hasQuit   chan bool
	running   bool
	logger    Logger
}

func newWorker(
	name string,
	work func(),
	interrupt func(),
	logger Logger,
) worker {
	return worker{
		Name:      name,
		work:      work,
		interrupt: interrupt,
		quit:      make(chan bool, 1),
		hasQuit:   make(chan bool, 1),
		running:   false,
		logger:    logger,
	}
}

func (w *worker) Start() error {
	if w.running {
		return errors.New(fmt.Sprintf("Worker %s has already started", w.Name))
	}

	w.running = true

	go func() {
		for {
			select {
			case <-w.quit:
				w.logger.Printf("Quitting %s worker\n", w.Name)
				w.hasQuit <- true
				return
			default:
				w.work()
			}
		}
	}()

	return nil
}

func (w *worker) Stop() {
	w.quit <- true

	for {
		select {
		case <-w.hasQuit:
			w.running = false
			return
		default:
			w.logger.Printf("Interrupting %s worker", w.Name)
			w.interrupt()
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (w *worker) IsRunning() bool {
	return w.running
}
