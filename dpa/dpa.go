package dpa

type DPA int

const (
	Deficient DPA = iota
	Perfect
	Abundant
)

func Sum(xs []int) (sum int) {
	for _, ii := range xs {
		sum = sum + ii
	}
	return sum
}

func ProperDivisors(xx int) (divisors []int) {
	for ii := 1; ii <= xx/2; ii++ {
		if xx%ii == 0 {
			divisors = append(divisors, ii)
		}
	}
	return divisors
}

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

func Text(in DPA) string {
	if in == Deficient {
		return "%v is a deficient number.\n\n"
	} else if in == Abundant {
		return "%v is an abundant number.\n\n"
	}
	return "%v is a perfect number.\n\n"
}

func Calc(in int) string {
    return Text(IsDPA(in))
}
