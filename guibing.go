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
	guibing(buf)
	fmt.Println(time.Since(t))
}

func guibing(buf []int) {
	tmp := make([]int, len(buf))
	merge_sort(buf, 0, len(buf)-1, tmp)
}

func merge_sort(a []int, first, last int, tmp [] int) {
	if first < last {
		middle := (first + last) / 2
		merge_sort(a, first, middle, tmp)
		merge_sort(a, middle+1, last, tmp)
		mergeArray(a, first, middle, last, tmp)
	}
}

func mergeArray(a []int, first, middle, end int, tmp []int) {
	i, m, j, n, k := first, middle, middle+1, end, 0
	for i <= m && j <= n {
		if a[i] <= a[j] {
			tmp[k] = a[j]
			k++
			i++
		} else {
			tmp[k] = a[j]
			k++
			j++
		}
	}
	for i <= m {
		tmp[k] = a[i]
		k++
		i++
	}
	for j <= n {
		tmp[k] = a[j]
		k++
		j++
	}
	for ii := 0; ii < k; ii++ {
		a[first+ii] = tmp[ii]
	}
}