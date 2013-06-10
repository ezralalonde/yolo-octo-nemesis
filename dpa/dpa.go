package dpa

type DPA int

const (
	Deficient DPA = iota
	Perfect
	Abundant
)

//Sum calculates the sum of all elements of parameter `xs`,
//and returns the value.
func Sum(xs []int) (sum int) {
	for _, ii := range xs {
		sum = sum + ii
	}
	return sum
}

//ProperDivisors calculates all of the proper divisors of
//parameter `xx` and returns a slice of them.
//Proper divisors are all numbers Y such that xx % Y === 0
func ProperDivisors(xx int) (divisors []int) {
	for ii := 1; ii <= xx/2; ii++ {
		if xx%ii == 0 {
			divisors = append(divisors, ii)
		}
	}
	return divisors
}

//IsDPA akes a number `in`, and returns whether it is
//Deficient, Perfect, or Abundant.
func IsDPA(in int) DPA {
	pd := ProperDivisors
	sum := Sum(pd(in))
	if in < sum {
		return Abundant
	} else if in > sum {
		return Deficient
	}
	return Perfect
}

//Text takes a DPA value as a parameter, and outputs an appropriate
//format string relating to the value of the parameter.
func Text(in DPA) string {
	if in == Deficient {
		return "%v is a deficient number.\n\n"
	} else if in == Abundant {
		return "%v is an abundant number.\n\n"
	}
	return "%v is a perfect number.\n\n"
}

//Calc takes an int, and returns a piece of text regarding the
//Abundance, Deficiency, or Perfection of a number as calculated
//by `IsDPA` and `Text`.
func Calc(in int) string {
    return Text(IsDPA(in))
}
