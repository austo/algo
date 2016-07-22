package odometer

type O struct {
	limit   int
	curr    int
	r       []int
	clock   []int
	started bool
}

func New(r []int, repeat int) *O {
	o := &O{
		r:     r,
		limit: len(r),
		clock: make([]int, repeat),
		curr:  repeat - 1,
	}
	return o
}

func (o *O) Next() []int {
	if len(o.r) == 0 {
		return nil
	}
	if o.started {
		if o.clock[o.curr] == o.limit-1 {
			for o.curr >= 0 && o.clock[o.curr] == o.limit-1 {
				o.curr--
			}
			if o.curr < 0 {
				return nil
			}
			o.clock[o.curr]++
			for i := o.curr + 1; i < len(o.clock); i++ {
				o.clock[i] = 0
			}
			o.curr = len(o.clock) - 1
		} else {
			o.clock[o.curr]++
		}
	} else {
		o.started = true
	}
	result := make([]int, 0, len(o.clock))
	for _, v := range o.clock {
		result = append(result, o.r[v])
	}
	return result
}
