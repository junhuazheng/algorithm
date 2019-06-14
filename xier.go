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
	xier(buf)
	fmt.Println(time.Since(t))
}

func xier(buf []int) {
	times := 0
	tmp := 0
	length := len(buf)
	incre := length

	for {
		incre /= 2
		for k := 0; k < incre; k++ {
			for i := k + incre; i < length; i += incre {
				for j := i; j > k; j -= incre {
					times++
					if buf[j] < buf[j-incre] {
						tmp = buf[j-incre]
						buf[j-incre] = buf[j]
						buf[j] = tmp
					} else {
						break
					}
				}
			}
		}
		if incre == 1 {
			break
		}
	}
	fmt.Println("xier times: ", times)
}