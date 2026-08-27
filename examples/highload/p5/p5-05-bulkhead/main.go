package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	stateStart    = "START"
	stateQueued   = "QUEUED"
	stateRejected = "REJECTED"
	stateDone     = "DONE"

	eventArrive = "ARRIVE"
	eventDone   = "DONE"
)

type EventDispatcher struct {
	Rejected                int
	limC, limQ, activeCount int
	active                  map[string]int
	queue                   []string
}

func New(n int, limC, limQ int) *EventDispatcher {
	return &EventDispatcher{
		active: make(map[string]int, n),
		queue:  make([]string, 0, n),
		limC:   limC,
		limQ:   limQ,
	}
}

func (ed *EventDispatcher) addActive(id string) {
	ed.activeCount++
	ed.active[id] += 1
}
func (ed *EventDispatcher) doneActive(id string) {
	ed.activeCount--
	if ed.activeCount < 0 {
		ed.activeCount = 0
	}
	val, ok := ed.active[id]
	if !ok {
		return
	}
	if val < 2 {
		delete(ed.active, id)
		return
	}
	ed.active[id] -= 1
}

func (ed *EventDispatcher) printing(printer func(eventId, state string), eventId, state string) {
	printer(eventId, state)
}

func (ed *EventDispatcher) Stat() (int, int, int) {
	return ed.activeCount, len(ed.queue), ed.Rejected
}
func (ed *EventDispatcher) Dispatch(event, id string) string {
	if event != eventArrive && event != eventDone {
		return ""
	}
	if event == eventArrive {
		if ed.activeCount < ed.limC {
			ed.addActive(id)
			return stateStart
		}
		// если пул закончился
		if len(ed.queue) < ed.limQ {
			ed.queue = append(ed.queue, id)
			return stateQueued
		}
		ed.Rejected++
		return stateRejected
	}
	ed.doneActive(id)
	if len(ed.queue) == 0 {
		return stateDone
	}
	// нужно взять из очереди
	id2 := ed.queue[0]
	ed.queue = ed.queue[1:]
	ed.addActive(id2)

	ed.printing(printEvent, id, stateDone)
	ed.printing(printEvent, id2, stateStart)
	return ""
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var (
		c, q, n int
	)

	_, err := fmt.Fscan(r, &c, &q)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}

	_, err = fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(2)
	}
	ed := New(n, c, q)
	for i := 0; i < n; i++ {
		var (
			event, id string
		)

		_, err = fmt.Fscan(r, &event, &id)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(2)
		}
		state := ed.Dispatch(event, id)
		if state != "" {
			printEvent(id, state)
		}
	}
	a, q, rj := ed.Stat()
	printStat(a, q, rj)
}

func printStat(aCount, qCount, rCount int) {
	fmt.Printf("ACTIVE %d QUEUED %d REJECTED %d\n", aCount, qCount, rCount)
}

func printEvent(eventId, state string) {
	fmt.Printf("%s %s\n", state, eventId)
}
