package sort

import (
	"os"
	"sort"
	"sync"
)

// copied from my solution for pingcap go talent plan

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	switch os.Getenv("MERGE_SORT") {
	case "mergeOnlySort":
		sorted := mergeOnlySort(src)
		copy(src, sorted)
	case "mergeWithInsertion":
		sorted := mergeWithInsertion(src)
		copy(src, sorted)
	case "mergeSortWithBuf":
		mergeSortWithBufNew(src)
	case "mergeSortWithHalfBuf":
		mergeSortWithHalfBufNew(src)
	case "sort":
		sort.Slice(src, func(i, j int) bool { return src[i] < src[j] })
	default:
		mergeSortParallelWithBufNew(src)
	}
}

func MergeOnlySort(src []int64) {
	sorted := mergeOnlySort(src)
	copy(src, sorted)
}

func MergeWithInsertionSort(src []int64) {
	sorted := mergeWithInsertion(src)
	copy(src, sorted)
}

func MergeSortWithBuf(src []int64) {
	mergeSortWithBufNew(src)
}

func MergeSortWithHalfBuf(src []int64) {
	mergeSortWithHalfBufNew(src)
}

func MergeSortParallelWithBuf(src []int64) {
	mergeSortParallelWithBufNew(src)
}

// mergeOnlySort does not have much optimization.
// It creates a LOT of small slices during merge.
func mergeOnlySort(a []int64) []int64 {
	size := len(a)
	// No need to sort nor merge for [], [0], [1, 2]
	if size <= 2 {
		if size == 2 {
			if a[0] > a[1] {
				a[0], a[1] = a[1], a[0]
			}
		}
		return a
	}

	mid := size / 2
	l := mergeOnlySort(a[0:mid])
	r := mergeOnlySort(a[mid:])
	if l[mid-1] <= r[0] {
		return append(l, r...)
	}
	return merge(l, r)
}

// merge takes two sorted slice and merge into a new sorted one.
func merge(a []int64, b []int64) []int64 {
	sa, sb := len(a), len(b)
	sc := sa + sb
	c := make([]int64, sc)

	i, ia, ib := 0, 0, 0
	for ; i < sc; i++ {
		if ia == sa || ib == sb {
			break
		}
		if a[ia] <= b[ib] {
			c[i] = a[ia]
			ia++
		} else {
			c[i] = b[ib]
			ib++
		}
	}

	// copy the rest
	if i == sc {
		return c
	}
	if ia < sa {
		copy(c[i:], a[ia:])
	}
	if ib < sb {
		copy(c[i:], b[ib:])
	}
	return c
}

// mergeWithInsertion use insertionSort when the size is smaller than 64.
// It still allocates tmp slice, but it should be less because we didn't do merge sort all the way down.
func mergeWithInsertion(a []int64) []int64 {
	size := len(a)
	// No need to sort nor merge for [], [0], [1, 2]
	if size <= 64 {
		//if size <= 1024 { // when size of insertion sort is too large, it take much longer time
		insertionSort(a)
		return a
	}

	mid := size / 2
	l := mergeWithInsertion(a[0:mid])
	r := mergeWithInsertion(a[mid:])
	if l[mid-1] <= r[0] {
		return append(l, r...)
	}
	return merge(l, r)
}

// insertionSort is in place insertion sort
func insertionSort(a []int64) {
	for i := 0; i < len(a); i++ {
		x := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			// keep shifting until find the right slot
			if a[j] > x {
				// shift
				a[j+1] = a[j] // when j = i - 1 we will override a[i], thus we need to keep it as x in the beginning
			} else {
				break
			}
		}
		a[j+1] = x
	}
}

func mergeSortWithBufNew(a []int64) {
	buf := make([]int64, len(a))
	mergeSortWithBuf(a, buf)
}

// mergeSorWithBuf requires a dynamic buf that has same size as input.
func mergeSortWithBuf(a []int64, buf []int64) {
	size := len(a)
	// No need to sort nor merge for [], [0], [1, 2]
	if size <= 2 {
		if size == 2 {
			if a[0] > a[1] {
				a[0], a[1] = a[1], a[0]
			}
		}
		return
	}

	mid := size / 2
	mergeSortWithBuf(a[0:mid], buf[0:mid])
	mergeSortWithBuf(a[mid:], buf[mid:])
	mergeWithBuf(a, 0, mid, len(a), buf)
}

// mergeWith copies the two sorted slice into the buf and merge onto the original one
func mergeWithBuf(src []int64, start, mid, end int, buf []int64) []int64 {
	copy(buf, src[start:end])

	a := buf[start:mid]
	b := buf[mid:end]
	sa, sb := len(a), len(b)
	sc := sa + sb
	c := src[start:]
	i, ia, ib := 0, 0, 0
	for ; i < sc; i++ {
		if ia == sa || ib == sb {
			break
		}
		if a[ia] <= b[ib] {
			c[i] = a[ia]
			ia++
		} else {
			c[i] = b[ib]
			ib++
		}
	}

	// copy the rest
	if i == sc {
		return c
	}
	if ia < sa {
		copy(c[i:], a[ia:])
	}
	if ib < sb {
		copy(c[i:], b[ib:])
	}
	return c
}

func mergeSortWithHalfBufNew(a []int64) {
	buf := make([]int64, len(a)/2+1)
	mergeSortWithHalfBuf(a, buf)
}

// mergeSortWIthHalfBuf pass a fixed size buf in all recursive calls.
func mergeSortWithHalfBuf(a []int64, buf []int64) {
	size := len(a)
	// No need to sort nor merge for [], [0], [1, 2]
	if size <= 2 {
		if size == 2 {
			if a[0] > a[1] {
				a[0], a[1] = a[1], a[0]
			}
		}
		return
	}

	mid := size / 2
	mergeSortWithHalfBuf(a[0:mid], buf)
	mergeSortWithHalfBuf(a[mid:], buf)
	mergeWithHalfBuf(a, 0, mid, len(a), buf)
}

// mergeWithHalfBuf copies the first sorted slice into the buf and merge onto the original one
func mergeWithHalfBuf(src []int64, start, mid, end int, buf []int64) []int64 {
	copy(buf, src[start:mid]) // NOTE: it's dst, src ... spent a long time on this ...

	a := buf[0 : mid-start]
	b := src[mid:end]
	sa, sb := len(a), len(b)
	sc := sa + sb
	c := src[start:]
	i, ia, ib := 0, 0, 0
	for ; i < sc; i++ {
		if ia == sa || ib == sb {
			break
		}
		if a[ia] <= b[ib] {
			c[i] = a[ia]
			ia++
		} else {
			c[i] = b[ib]
			ib++
		}
	}

	// copy the rest
	if i == sc {
		return c
	}
	if ia < sa {
		copy(c[i:], a[ia:])
	}
	if ib < sb {
		copy(c[i:], b[ib:])
	}
	return c
}

// magic number for 16MB and 8 core CPU. It can be determined at runtime but let's just make it easier for a lab.
const twoMB = 2 << 20

//const twoMB = 2 << 10

func mergeSortParallelWithBufNew(a []int64) {
	buf := make([]int64, len(a))
	mergeSortParallelWithBuf(a, buf)
}

// use multiple go routines until we have too many go routines and switch to serial implementation
func mergeSortParallelWithBuf(a []int64, buf []int64) {
	size := len(a)
	// No need to sort nor merge for [], [0], [1, 2]
	if size <= 2 {
		if size == 2 {
			if a[0] > a[1] {
				a[0], a[1] = a[1], a[0]
			}
		}
		return
	}

	mid := size / 2
	if size < twoMB {
		mergeSortWithBuf(a[0:mid], buf[0:mid])
		mergeSortWithBuf(a[mid:], buf[mid:])
	} else {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			mergeSortWithBuf(a[0:mid], buf[0:mid])
			wg.Done()
		}()
		go func() {
			mergeSortWithBuf(a[mid:], buf[mid:])
			wg.Done()
		}()
		wg.Wait()
	}
	mergeWithBuf(a, 0, mid, len(a), buf)
}
