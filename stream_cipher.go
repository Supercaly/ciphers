package main

import (
	"log"
	"math/rand"
)

func encodeStreamCip(seed Bits, msg Bits) Bits {
	key := PRGStreamCip(seed, msg.length())
	c, err := xor(key, msg)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func decodeStreamCip(seed Bits, c Bits) Bits {
	key := PRGStreamCip(seed, c.length())
	m, err := xor(key, c)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func PRGStreamCip(seed Bits, length int) Bits {
	rand.Seed(int64(seed.toInt()))
	var vals []bool
	for i := 0; i < length; i++ {
		vals = append(vals, rand.Intn(2) == 1)
	}
	return Bits{vals}
}
