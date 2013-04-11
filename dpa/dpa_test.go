package dpa

import (
	"fmt"
	"testing"
)

func ExampleProperDivisors() {
	ProperDivisors(6)
	// Return: [1, 2, 3]
}

func o(in interface{}) string {
	return fmt.Sprintf("%v", in)
}

func tPd(in int, out []int, t *testing.T) {
	f := ProperDivisors
	if x, y := o(f(in)), o(out); x != y {
		t.Errorf("ProperDivisors(%v) = %v, want %v", in, f(in), out)
	}
}

func tSum(in []int, out int, t *testing.T) {
	f := Sum
	if f(in) != out {
		t.Errorf("got Sum(%v) = %v; want %v", in, f(in), out)
	}
}

func tDpa(in int, out DPA, t *testing.T) {
    f := IsDPA
    if f(in) != out {
		t.Errorf("got %v is %v; want %v", in, f(in), out)
    }
}

func TestSum1(t *testing.T) {
    tSum([]int{1,2,3}, 6, t)
}
func TestSum2(t *testing.T) {
    tSum([]int{1,2,3,8}, 14, t)
}
func TestSum3(t *testing.T) {
    tSum([]int{2,4,6}, 12, t)
}

func TestProperDivisor6(t *testing.T) {
	tPd(6, []int{1, 2, 3}, t)
}

func TestProperDivisor4(t *testing.T) {
	tPd(4, []int{1, 2}, t)
}

func TestProperDivisor12(t *testing.T) {
	tPd(12, []int{1, 2, 3, 4, 6}, t)
}

func TestDpa6(t *testing.T) {
    tDpa(6, Perfect, t)
}
func TestDpa4(t *testing.T) {
    tDpa(4, Deficient, t)
}
func TestDpa12(t *testing.T) {
    tDpa(12, Abundant, t)
}
