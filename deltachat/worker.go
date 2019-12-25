package deltachat

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type worker struct {
	Name      string
	work      func()
	interrupt func()
	quit      chan bool
	hasQuit   chan bool
	running   bool
}

func newWorker(name string, work func(), interrupt func()) worker {
	return worker{
		Name:      name,
		work:      work,
		interrupt: interrupt,
		quit:      make(chan bool, 1),
		hasQuit:   make(chan bool, 1),
		running:   false,
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
				log.Printf("Quitting %s worker\n", w.Name)
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
			log.Printf("Interrupting %s worker", w.Name)
			w.interrupt()
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (w *worker) IsRunning() bool {
	return w.running
}
