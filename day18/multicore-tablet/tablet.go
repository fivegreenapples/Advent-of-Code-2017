package multicoretablet

import (
	"errors"
	"fmt"
	"sync"
)

// Tablet describes a 2 core programmable device
type Tablet struct {
	cores     []singleCore
	commsCh   chan commsEvent
	sendStats []int
}

type commsEvent struct {
	typ      string
	sendVal  int
	sendTo   int
	recvTo   int
	recvChan chan int
}

// New constructs a Tablet with 2 cores and using the given
// set of programming instructions
func New(rawInstructions []string) *Tablet {

	tab := Tablet{}

	core0 := makeCore(&tab, 0, rawInstructions)
	core1 := makeCore(&tab, 1, rawInstructions)

	tab.cores = append(tab.cores, core0)
	tab.cores = append(tab.cores, core1)

	return &tab
}

// reset initialises each Tablet core to a starting state where all
// internal registers are zero.
func (t *Tablet) reset() {
	t.sendStats = []int{}
	for i := range t.cores {
		t.cores[i].reset()
		t.sendStats = append(t.sendStats, 0)
	}
}

func (t *Tablet) coordinateInterprocessComs(wg *sync.WaitGroup) {

	// designed to be run in a goroutine and stopped by closing the
	// comms channel

	rcvBuffers := make([]fifo, 2)
	waitingRecvChans := make([]chan int, 2)

	for ev := range t.commsCh {
		switch ev.typ {
		case "send":
			if waitingRecvChans[ev.sendTo] != nil {
				waitingRecvChans[ev.sendTo] <- ev.sendVal
				waitingRecvChans[ev.sendTo] = nil
			} else {
				rcvBuffers[ev.sendTo].push(ev.sendVal)
			}
		case "recv":
			val, valid := rcvBuffers[ev.recvTo].pull()
			if !valid {
				if waitingRecvChans[ev.recvTo] != nil {
					panic("can't receive to same core more than once")
				}
				waitingRecvChans[ev.recvTo] = ev.recvChan
			} else {
				ev.recvChan <- val
			}
		default:
			panic("unhandled event type: " + ev.typ)
		}

		allWaiting := true
		for _, c := range waitingRecvChans {
			if c == nil {
				allWaiting = false
			}
		}
		if allWaiting {
			for id, c := range waitingRecvChans {
				close(c)
				waitingRecvChans[id] = nil
			}
		}
	}

	// comms closed
	wg.Done()

}

func (t *Tablet) send(coreID int, val int) {
	t.sendStats[coreID]++
	t.commsCh <- commsEvent{
		typ:     "send",
		sendVal: val,
		sendTo:  (coreID + 1) % len(t.cores), // send to next core in list
	}
}
func (t *Tablet) receive(coreID int) (int, error) {
	recvChan := make(chan int)
	t.commsCh <- commsEvent{
		typ:      "recv",
		recvTo:   coreID,
		recvChan: recvChan,
	}

	val, notclosed := <-recvChan
	if !notclosed {
		return 0, errors.New("deadlock")
	}
	return val, nil
}

// Run executes the Tablet's instructions on all cores in parallel
func (t *Tablet) Run() {

	t.reset()

	wgComms := sync.WaitGroup{}
	wgCores := sync.WaitGroup{}

	wgComms.Add(1)
	t.commsCh = make(chan commsEvent)
	go t.coordinateInterprocessComs(&wgComms)

	for i := range t.cores {
		wgCores.Add(1)
		thisRef := i
		go func() {
			t.cores[thisRef].run()
			wgCores.Done()
		}()
	}

	wgCores.Wait()
	close(t.commsCh)
	wgComms.Wait()

	fmt.Printf("Program terminated\n")
	for i, sends := range t.sendStats {
		fmt.Printf("Sends from core %d: %d\n", i, sends)
	}
}
