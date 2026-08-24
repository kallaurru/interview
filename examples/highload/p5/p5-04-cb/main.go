package main

import (
	"bufio"
	"fmt"
	"os"
	"sync/atomic"
)

const (
	stateClosed   uint32 = 0x00
	stateOpen     uint32 = 0x02
	stateHalfOpen uint32 = 0x01
	eventOk              = "OK"
	eventFail            = "FAIL"
)

type CircuitBreaker struct {
	limF        int // Сколько подряд идущих ошибок размыкают цепь
	limK        int // Сколько успешных проб закрывают цепь
	limT        int // Сколько тиков цепь держится в Open
	failCounter atomic.Uint32
	openTick    atomic.Uint32 // Момент перехода в состояние open
	state       atomic.Uint32
}

func New(f, k, t int) *CircuitBreaker {
	cb := &CircuitBreaker{
		limF: f,
		limK: k,
		limT: t,
	}

	cb.state.Store(stateClosed)
	cb.failCounter.Store(0)
	cb.openTick.Store(0)

	return cb
}

func (cb *CircuitBreaker) Event(id string, tick int) {
	old, ok := cb.changeState(id)
	if !ok {
		cb.print(0, tick, id, 0)
		return
	}
	cb.print(1, tick, id, old)
}

func (cb *CircuitBreaker) changeState(eventID string) (uint32, bool) {
	oldState := cb.state.Load()
	switch oldState {
	case stateOpen:
	case stateClosed:
		if eventID == eventOk {
			cb.clearFC()
			return oldState, false
		}
		cb.failCounter.Add(1)
		if cb.failCounter.Load() >= uint32(cb.limF) {
			cb.state.Store(stateOpen)
			return oldState, true
		}

		return oldState, false
	case stateHalfOpen:

	}
}

func (cb CircuitBreaker) PrintFinal() {
	cb.print(2, 0, "", 0)
}

func (cb CircuitBreaker) print(scenario int, tick int, eventId string, oldState uint32) {
	switch scenario {
	default:
		fmt.Printf("%d %s %s\n", tick, cb.printState(cb.state.Load()), eventId)
	case 1:
		// меняем состояние
		fmt.Printf("%d %s %s -> %s\n", tick, cb.printState(oldState), eventId, cb.printState(cb.state.Load()))
	case 2:
		// финальное состояние
		fmt.Printf("FINAL %s\n", cb.printState(cb.state.Load()))
	}
}

func (cb CircuitBreaker) printState(st uint32) string {
	switch st {
	case stateOpen:
		return "OPEN"
	case stateClosed:
		return "CLOSED"
	case stateHalfOpen:
		return "HALF_OPEN"
	}
	return ""
}

func (cb *CircuitBreaker) clearFC() {
	cb.failCounter.Store(0)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var (
		f, t, k, n int
	)

	_, err := fmt.Fscan(r, &f, &t, &k)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}

	_, err = fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(2)
	}

	cb := New(f, k, t)
	for i := 0; i < n; i++ {
		var (
			res string
		)

		_, err = fmt.Fscan(r, &res)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(2)
		}
		cb.Event(res, i+1)
	}

	cb.PrintFinal()
}
