package main

import (
	"errors"
	"flag"
	"fmt"
)

var n int
var r bool

func init() {
	flag.IntVar(&n, "n", 10, "number of disks")
	flag.BoolVar(&r, "r", false, "use recursive solution")
}

func main() {
	flag.Parse()
	towers := [...]Stack{
		make(Stack, 0, n),
		make(Stack, 0, n),
		make(Stack, 0, n),
	}

	for i := n; i > 0; i-- {
		towers[0].push(i)
	}

	var f func(int, *Stack, *Stack, *Stack) error
	if r {
		f = moveDisks
	} else {
		f = moveDisksNR
	}

	if err := f(n, &towers[0], &towers[2], &towers[1]); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(towers[2])
	}
}

// iterative version
func moveDisksNR(n int, from, to, buffer *Stack) error {
	// totalMoves := int(math.Pow(2, float64(n))) - 1

	var t, b *Stack
	if n%2 == 0 {
		t, b = buffer, to
	} else {
		t, b = to, buffer
	}

	var err error
	// for i := 1; i <= totalMoves; i++ {
	i := 1
	for len(*to) < n {
		if err != nil {
			break
		}
		switch i % 3 {
		case 0:
			err = moveDiskLegal(b, t)
		case 1:
			err = moveDiskLegal(from, t)
		case 2:
			err = moveDiskLegal(from, b)
		}
		i++
	}
	return err
}

func moveDiskLegal(from, to *Stack) (err error) {
	defer func() {
		if v := recover(); v != nil {
			err = errors.New(v.(string))
		}
	}()
	if len(*from) == 0 {
		from.push(to.pop())
	} else if len(*to) == 0 {
		to.push(from.pop())
	} else if from.peek() > to.peek() {
		from.push(to.pop())
	} else {
		to.push(from.pop())
	}
	return err
}

// recursive version
func moveDisks(k int, from, to, buffer *Stack) error {
	if k <= 0 {
		return nil
	}
	if err := moveDisks(k-1, from, buffer, to); err != nil {
		return err
	}
	if err := moveDisk(from, to); err != nil {
		return err
	}
	return moveDisks(k-1, buffer, to, from)
}

func moveDisk(from, to *Stack) error {
	v := from.pop()
	if len(*to) == 0 || to.peek() > v {
		to.push(v)
		return nil
	}
	return fmt.Errorf("cannot move %d above %d\n", v, to.peek())
}

type Stack []int

func (s *Stack) push(v int) {
	*s = append([]int{v}, *s...)
}

func (s *Stack) pop() int {
	if len(*s) == 0 {
		panic("pop: empty")
	}
	v := (*s)[0]
	*s = (*s)[1:]
	return v
}

func (s Stack) peek() int {
	if len(s) == 0 {
		panic("peek: empty")
	}
	return s[0]
}
