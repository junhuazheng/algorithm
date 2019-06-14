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
	kuaisu(buf)
	fmt.Println(time.Since(t))
}

func kuaisu(buf []int) {
	kuai(buf, 0, len(buf)-1)
}

func kuai(a []int, l, r int) {
	if l >= r {
		return
	}
	i, j, key := l, r, a[1]
	for i < j {
		for i < j && a[j] > key {
			j--
		}
		if i < j {
			a[i] = a[j]
			i++
		}
		for i < j && a[i] < key {
			i++
		}
		if i < j {
			a[j] = a[i]
			j--
		}
	}
	a[i] = key
	kuai(a, l, i-1)
	kuai(a, i+1, r)
}