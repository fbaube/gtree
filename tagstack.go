package gtree

type Tagentry struct {
	tag   string
	index int
}

func NewTagentry(aTag string, anIndex int) Tagentry {
	p := new(Tagentry)
	p.tag = aTag
	p.index = anIndex
	// println("Pushed:", aTag)
	return *p
}

func (pTE *Tagentry) Tag() string {
	return pTE.tag
}

func (pTE *Tagentry) Index() int {
	return pTE.index
}

// gagstack is a LIFO stack for GTags.
type Tagstack []Tagentry

// func NewTagStack() tagstack { return make([]tagentry, 10) }

// IsEmpty is a no-brainer.
func (ts Tagstack) IsEmpty() bool { return len(ts) == 0 }

// Peek will barf on an empty stack.
func (ts Tagstack) Peek() Tagentry { return ts[len(ts)-1] }

// Push reslices the stack.
func (ts *Tagstack) Push(te Tagentry) { (*ts) = append((*ts), te) }

// Pop will barf on an empty stack.
func (ts *Tagstack) Pop() Tagentry {
	d := (*ts)[len(*ts)-1]
	(*ts) = (*ts)[:len(*ts)-1]
	return d
}
