package main

import "fmt"

func main() {

	testFirewall := firewallFromInput(testInput)
	testFirewall.reset()
	testPenalty := testFirewall.sendPacket()
	fmt.Println("Test:", testPenalty)

	part1Firewall := firewallFromInput(input)
	part1Firewall.reset()
	part1Penalty := part1Firewall.sendPacket()
	fmt.Println("Part1:", part1Penalty)

	fmt.Println("Test2:", testFirewall.findDelayForNoPenalty(10000000))
	fmt.Println("Part2:", part1Firewall.findDelayForNoPenalty(10000000))

}

func firewallFromInput(in map[int]int) firewall {
	layers := []layer{}
	for d, r := range in {
		layers = append(layers, layer{
			depth:      d,
			rainge:     r,
			scannerPos: 0,
		})
	}
	return makeFirewall(layers)
}

type firewall struct {
	layers   map[int]layer
	endLayer int
}

func makeFirewall(layers []layer) firewall {
	fw := firewall{}
	fw.layers = map[int]layer{}
	fw.endLayer = -1

	for _, l := range layers {
		l.reset()
		fw.layers[l.depth] = l
		if l.depth > fw.endLayer {
			fw.endLayer = l.depth
		}
	}

	return fw
}

func (f *firewall) reset() {
	for d, l := range f.layers {
		l.reset()
		f.layers[d] = l
	}
}
func (f *firewall) resetTo(delay int) {
	f.reset()
	for delay > 0 {
		f.advance()
		delay--
	}
}
func (f *firewall) sendPacket() int {
	packetPos, penalty := -1, 0
	for {
		// advance packet
		packetPos++
		// check if we've exited
		if packetPos > f.endLayer {
			return penalty
		}
		// check if we're actually on an active layer
		thisLayer, exists := f.layers[packetPos]
		if exists {
			// check scanner position at layer
			if thisLayer.scannerPos == 0 {
				penalty += (packetPos * thisLayer.rainge)
			}
		}
		// advance all scanners
		f.advance()
	}
}
func (f *firewall) sendPacketNoPenalty() bool {
	packetPos := -1
	for {
		// advance packet
		packetPos++
		// check if we've exited
		if packetPos > f.endLayer {
			return true
		}
		// check if we're actually on an active layer
		// and if we're caught
		if thisLayer, exists := f.layers[packetPos]; exists && thisLayer.scannerPos == 0 {
			return false
		}
		// advance all scanners
		f.advance()
	}
}

func (f *firewall) findDelayForNoPenalty(maxDelay int) int {
	// work out delays which will never work
	sieve := make([]bool, maxDelay)
	for _, l := range f.layers {
		a := -l.depth
		for a < 0 {
			a += l.cyclocity
		}
		for a < maxDelay {
			sieve[a] = true
			a += l.cyclocity
		}
	}
	// now find first delay that asn't been set
	for d, set := range sieve {
		if !set {
			return d
		}
	}
	return -1
	// testDelay := 0
	// for {
	// 	f.resetTo(testDelay)
	// 	if notcaught := f.sendPacketNoPenalty(); notcaught {
	// 		break
	// 	}
	// 	testDelay++
	// }
	// return testDelay
}

func (f *firewall) advance() {
	for d, l := range f.layers {
		l.advanceScanner()
		f.layers[d] = l
	}
}

type layer struct {
	depth      int
	rainge     int
	cyclocity  int
	scannerPos int
	scannerRev bool
}

func (l *layer) reset() {
	l.scannerPos = 0
	l.scannerRev = false
	l.cyclocity = (2 * l.rainge) - 2
}

func (l *layer) advanceScanner() {
	if l.scannerRev {
		l.scannerPos--
		l.scannerRev = l.scannerPos > 0
	} else {
		l.scannerPos++
		l.scannerRev = l.scannerPos == (l.rainge - 1)
	}
}
