package multicoretablet

type fifo struct {
	buf      map[uint]int
	nextPush uint
	nextPull uint
}

func (f *fifo) push(val int) {
	if f.buf == nil {
		f.buf = map[uint]int{}
	}
	f.buf[f.nextPush] = val
	f.nextPush++
}

func (f *fifo) pull() (int, bool) {
	if f.buf == nil {
		f.buf = map[uint]int{}
	}
	val, found := f.buf[f.nextPull]
	if !found {
		return 0, false
	}
	delete(f.buf, f.nextPull)
	f.nextPull++
	return val, true
}
