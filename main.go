package main

import (
	"fmt"
)

func main() {
	msg1 := NewBits("01110100011001010111001101110100")
	key1 := NewBits("01100001011000100110001101100100")
	m1 := decodeOTP(key1, encodeOTP(key1, msg1))
	fmt.Println(m1)

	msg2 := 69
	key2 := 420
	m2 := decodeAOTP(key2, encodeAOTP(key2, msg2))
	fmt.Println(m2)

	msg3 := NewBits("01110100011001010111001101110100")
	seed := NewBits("0110000")
	m3 := decodeStreamCip(seed, encodeStreamCip(seed, msg3))
	fmt.Println(m3)
}
