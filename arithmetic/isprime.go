// Package arithmetic IsPrime function checks if a number is prime or not
// Even function returns true if number is even
package arithmetic

func IsPrime(num int) bool {
	for i := 2; i < int(num/2); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func Even(x uint) bool {
	if x%2 == 0 {
		return true
	}
	return false
}
