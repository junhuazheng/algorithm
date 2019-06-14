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
	maopao(buf)
	fmt.Println(time.Since(t))
}

func maopao(buf []int) {
	times := 0
	for i := 0; i < len(buf)-1; i++ {
		flag := false
		for j := 1; j < len(buf)-i; j++ {
			if buf[j-1] > buf[j] {
				times++
				tmp := buf[j-1]
				buf[j-1] = buf[j]
				buf[j] = tmp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	fmt.Println("maopao times: ", times)
}