package odometer

import (
	"fmt"
	"math"
	"testing"
)

func TestBasic(t *testing.T) {
	r := []int{1, 2, 3}
	repeat := 3
	o := New(r, repeat)
	i := 0
	for {
		v := o.Next()
		if v == nil {
			break
		}
		i++
		fmt.Printf("%d: %v\n", i, v)
	}
	expected := int(math.Pow(float64(len(r)), float64(repeat)))
	if i != expected {
		t.Errorf("expected %d, got %d", expected, i)
	}
}
