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

	scenarioDef         = 0x00
	scenarioChangeState = 0x01
	scenarioRejected    = 0x02
	scenarioFinal       = 0x03
)

type CircuitBreaker struct {
	limF        int // Сколько подряд идущих ошибок размыкают цепь
	limK        int // Сколько успешных проб закрывают цепь
	limT        int // Сколько тиков цепь держится в Open
	failCounter atomic.Uint32
	tryCounter  atomic.Uint32
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
	cb.tryCounter.Store(0)
	cb.openTick.Store(0)

	return cb
}

func (cb *CircuitBreaker) Event(id string, tick int) {
	old, scenario := cb.changeState(id, tick)
	cb.print(scenario, tick, id, old)
}

func (cb *CircuitBreaker) changeState(eventID string, tick int) (uint32, int) {
	oldState := cb.state.Load()
	switch oldState {
	case stateOpen:
		if uint32(tick) <= cb.openTick.Load()+uint32(cb.limT) {
			return oldState, scenarioRejected
		}
		cb.state.Store(stateHalfOpen)
		oldState = cb.state.Load()
		fallthrough
	case stateHalfOpen:
		if eventID == eventFail {
			cb.state.Store(stateOpen)
			cb.tryCounter.Store(0)
			cb.openTick.Store(uint32(tick))

			return oldState, scenarioChangeState
		}
		cb.tryCounter.Add(1)
		if cb.tryCounter.Load() >= uint32(cb.limK) {
			cb.state.Store(stateClosed)
			cb.failCounter.Store(0)
			cb.tryCounter.Store(0)
			cb.openTick.Store(0)
			return oldState, scenarioChangeState
		}
		// попытки в рамках лимита open event OK
		// если зашли из состояния open
		return oldState, scenarioDef
	case stateClosed:
		if eventID == eventOk {
			cb.clearFC()
			return oldState, scenarioDef
		}
		cb.failCounter.Add(1)
		if cb.failCounter.Load() >= uint32(cb.limF) {
			cb.state.Store(stateOpen)
			cb.openTick.Store(uint32(tick))
			return oldState, scenarioChangeState
		}

		return oldState, scenarioDef
	}

	return oldState, scenarioDef
}

func (cb CircuitBreaker) PrintFinal() {
	cb.print(scenarioFinal, 0, "", 0)
}

func (cb CircuitBreaker) print(scenario int, tick int, eventId string, oldState uint32) {
	switch scenario {
	default:
		fmt.Printf("%d %s %s\n", tick, cb.printState(cb.state.Load()), eventId)
	case scenarioChangeState:
		// меняем состояние
		fmt.Printf("%d %s %s -> %s\n", tick, cb.printState(oldState), eventId, cb.printState(cb.state.Load()))
	case scenarioRejected:
		// отклонение запросов
		fmt.Printf("%d %s REJECTED\n", tick, cb.printState(cb.state.Load()))
	case scenarioFinal:
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
