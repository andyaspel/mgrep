package worklist

//* entry into channel

type Entry struct {
	Path string
}
type Worklist struct {
	jobs chan Entry
}

func (w *Worklist) Add(work Entry) {
	w.jobs <- work
}

func (w *Worklist) Next() Entry {
	j := <-w.jobs
	return j
}

func New(bufSize int) Worklist {
	return Worklist{make(chan Entry, bufSize)}
}

func NewJob(path string) Entry {
	return Entry{path}
}

//* close chanels

func (w *Worklist) Finalize(nWorkers int) {
	for i := 0; i < nWorkers; i++ {
		w.Add(Entry{""})
	}
}
