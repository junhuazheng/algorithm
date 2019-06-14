package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	num = 100000
	rangeNum = 100000
)

func main() {
	randSeed := rand.New(rand.NewSource(time.Now().Unix() + time.Now().UnixNano()))
	var buf []int
	for i := 0; i < num; i++ {
		buf = append(buf, randSeed.Intn(rangeNum))
	}
	t := time.Now()
	xuanze(buf)
	fmt.Println(time.Since(t))
}

func xuanze(buf []int) {
	times := 0
	for i := 0; i < len(buf)-1; i++ {
		min := i 
		for j := i; j < len(buf); j++ {
			times++
			if buf[min] > buf[j] {
				min = j
			}
		}
		if min != 1 {
			tmp := buf[i]
			buf[i] = buf[min]
			buf[min] = tmp
		}
	}
	fmt.Println("xuanze times: ", times)
}