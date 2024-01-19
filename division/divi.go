package division

func Divide(num1 int, num2 int) (int, int) {
	quotient := num2 / num1
	remainder := num2 % num1
	return quotient, remainder
}
