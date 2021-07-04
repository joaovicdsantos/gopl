package main

import "fmt"

// pc[i] é a população de i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount devolve a população (número de bits definidos de x)
func PopCount(x uint64) int {
	var res int
	for i := 0; i < 8; i++ {
		res += int(pc[byte(x>>(i*8))])
	}
	return res
}

// PopCountBitBit devolve a população (número de bits definidos de x)
func PopCountBitBit(x uint64) int {
	var res int
	for i := 0; i < 64; i++ {
		res += int(pc[byte(x>>(i))])
	}
	return res
}

func main() {
	fmt.Println(PopCount(649))
	fmt.Println(PopCountBitBit(649))
}
