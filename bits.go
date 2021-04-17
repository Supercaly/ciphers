package main

import (
	"bytes"
	"errors"
	"log"
)

type Bits struct {
	value []bool
}

func NewBits(bits string) Bits {
	var values []bool
	for i, bit := range bits {
		switch bit {
		case rune(48):
			values = append(values, false)
		case rune(49):
			values = append(values, true)
		default:
			log.Fatalf("Invalid character %x at %d", bit, i)
		}
	}
	return Bits{values}
}

func (b Bits) String() (ret string) {
	var buffer bytes.Buffer
	for _, bit := range b.value {
		strBit := "0"
		if bit {
			strBit = "1"
		}
		buffer.WriteString(strBit)
	}
	return buffer.String()
}

func (b Bits) length() int {
	return len(b.value)
}

func xor(a Bits, b Bits) (Bits, error) {
	if a.length() != b.length() {
		return Bits{}, errors.New("a and b must have the same lenght")
	}

	var bits []bool
	for i := 0; i < a.length(); i++ {
		if a.value[i] == b.value[i] {
			bits = append(bits, false)
		} else {
			bits = append(bits, true)
		}
	}

	return Bits{bits}, nil
}

func (b Bits) toInt() (ret int) {
	if b.value[0] {
		ret++
	}
	for i := 1; i < len(b.value); i++ {
		ret = ret << 1
		if b.value[i] {
			ret++
		}
	}
	return ret
}
