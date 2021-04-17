package main

import "log"

func encodeOTP(key Bits, msg Bits) Bits {
	c, err := xor(key, msg)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func decodeOTP(key Bits, c Bits) Bits {
	m, err := xor(key, c)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func encodeAOTP(key int, msg int) int {
	return (msg + key) % 10000
}

func decodeAOTP(key int, c int) int {
	return (c - key) % 10000
}
