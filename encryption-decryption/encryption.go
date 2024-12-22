package main

import (
	"fmt"
	"math/big"
	"math/rand"
)

func GenerateOtp() (string, error) {
	min := big.NewInt(1000)
	randomNum := rand.Int()

	// Add 1000 to ensure the OTP is always 4 digits
	randomNumBigInt := big.NewInt(int64(randomNum))
	otp := new(big.Int).Add(randomNumBigInt, min)
	return otp.String(), nil
}

func main() {
	str, _ := GenerateOtp()
	fmt.Println(str)
}
