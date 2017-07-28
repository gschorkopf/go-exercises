package main

type stateMachine struct {
	st     string
	sendc  chan int
	statec chan string
}

func newStateMachine() *stateMachine {
	sm := &stateMachine{
		st:     "A",
		sendc:  make(chan int),
		statec: make(chan string),
	}
	go sm.loop()
	return sm
}

func (sm *stateMachine) loop() {
	for {
		select {
		case sm.statec <- sm.st:
		case v := <-sm.sendc:
			switch {
			case sm.st == "A" && v == 0:
				sm.st = "A"
			case sm.st == "A" && v == 1:
				sm.st = "B"
			case sm.st == "B" && v == 0:
				sm.st = "C"
			case sm.st == "B" && v == 1:
				sm.st = "A"
			case sm.st == "C" && v == 0:
				sm.st = "B"
			case sm.st == "C" && v == 1:
				sm.st = "A"
			}
		}
	}
}

func (sm *stateMachine) send(i int)    { sm.sendc <- i }
func (sm *stateMachine) state() string { return <-sm.statec }

func main() {
	sm := newStateMachine()
	sm.send(1)          // state A + 1 => state B
	sm.send(0)          // state B + 0 => state C
	println(sm.state()) // "state C"
}
