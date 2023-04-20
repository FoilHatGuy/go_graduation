package security

import (
	"crypto/rand"
	"fmt"
	"strconv"
)

func generateNumber(length int) (number int64, err error) {
	b := make([]byte, length)
	_, err = rand.Read(b)
	if err != nil {
		return 0, fmt.Errorf("[security:utils] while generating random number\n%s", err)
	}
	num, err := strconv.ParseInt(string(b), 10, 0)
	if err != nil {
		return 0, fmt.Errorf("[security:utils] while generating random number\n%s", err)
	}
	return num, nil
}
func validateLuhn(number int64) bool {
	return (number%10+checksum(number/10))%10 == 0
}
func calculateLuhn(number int64) int64 {
	checkNumber := checksum(number)

	if checkNumber == 0 {
		return number * 10
	}
	return number*10 + (10 - checkNumber)
}

func checksum(number int64) int64 {
	var luhn int64

	for i := 0; number > 0; i++ {
		cur := number % 10

		if i%2 == 0 { // even
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}

		luhn += cur
		number = number / 10
	}
	return luhn % 10
}
