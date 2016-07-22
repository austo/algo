package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// find the difference in height between two nodes in a binary tree

type Node struct {
	Left  *Node
	Right *Node
	Data  interface{}
}

func main() {
	r := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	root := buildTree(r)
	printTree(root)
	prettyPrintTree(root)
	n1, n2 := root.Right, root.Left.Left.Right
	if diff, err := diffHeight(root, n1, n2); err != nil {
		fmt.Println(err)
		fmt.Printf("difference between %v and %v: %d\n", n1.Data, n2.Data, diff)
	} else {
		fmt.Printf("difference between %v and %v: %d\n", n1.Data, n2.Data, diff)
	}
}

func diffHeight(root, n1, n2 *Node) (int, error) {
	l1, err := level(root, n1)
	if err != nil {
		return 0, err
	}
	l2, err := level(root, n2)
	if err != nil {
		return 0, err
	}
	if l1 > l2 {
		return l1 - l2, nil
	}
	return l2 - l1, nil
}

func level(root, target *Node) (int, error) {
	if root == nil {
		return 0, errors.New("height: bottom of tree")
	}
	if root == target {
		return 0, nil
	}

	lh, lerr := level(root.Left, target)
	rh, rerr := level(root.Right, target)

	if lerr != nil && rerr != nil {
		return 0, lerr
	}
	if lerr != nil {
		return rh + 1, nil
	}
	if rerr != nil {
		return lh + 1, nil
	}
	if lh > rh {
		return lh + 1, nil
	}
	return rh + 1, nil
}

func buildTree(r []interface{}) *Node {
	root := &Node{
		Data: r[0],
	}
	r = r[1:]
	for i, x := range r {
		insert(root, getMapBits(i), x)
	}
	return root
}

func insert(root *Node, c []int, x interface{}) {
	np := root
	var last int
	if c[len(c)-1] == 1 {
		last = 1
	}
	c = c[:len(c)-1]
	for _, dir := range c {
		if dir == 0 {
			np = np.Left
		} else {
			np = np.Right
		}
	}
	newNode := &Node{
		Data: x,
	}
	if last == 0 {
		np.Left = newNode
	} else {
		np.Right = newNode
	}
}

func getMap(index int) (r []int) {
	index += 2
	base := fmt.Sprintf("%b", index)
	baseMap := strings.Split(base, "")
	for i, v := range baseMap {
		if i == 0 {
			continue
		}
		b, _ := strconv.Atoi(v)
		r = append(r, b)
	}
	return
}

// only handles 32-bit integers
func getMapBits(index int) (r []int) {
	index += 2
	highest := false
	for _, f := range []int{32, 16, 8, 4, 2, 1} {
		b := index & f
		if !highest && b == 0 {
			continue
		}
		if b > 0 {
			highest = true
			r = append(r, 1)
		} else {
			r = append(r, 0)
		}
	}
	return r[1:]
}

// iterative level order
func printTree(root *Node) {
	q := []*Node{root}
	for len(q) > 0 {
		np := q[0]
		q = q[1:]
		fmt.Printf("%v ", np.Data)
		if np.Left != nil {
			q = append(q, np.Left)
		}
		if np.Right != nil {
			q = append(q, np.Right)
		}
	}
	fmt.Println()
}

func height(root *Node) int {
	if root == nil {
		return 0
	}

	lh := height(root.Left)
	rh := height(root.Right)
	if lh > rh {
		return lh + 1
	}
	return rh + 1
}

func printLevel(root *Node, h int) {
	if root == nil {
		return
	}
	if h == 1 {
		fmt.Printf("%v ", root.Data)
	}
	h--
	printLevel(root.Left, h)
	printLevel(root.Right, h)
}

// recursive level order
func prettyPrintTree(root *Node) {
	h := height(root)

	for i := 1; i <= h; i++ {
		printLevel(root, i)
		fmt.Println()
	}
}
