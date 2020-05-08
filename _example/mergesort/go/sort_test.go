package sort_test

import (
	"math/rand"
	stdsort "sort"
	"testing"
	"time"

	"github.com/benchhub/benchhub/_example/sort"
)

func BenchmarkStd(b *testing.B) {
	sorter := func(src []int64) {
		stdsort.Slice(src, func(i, j int) bool { return src[i] < src[j] })
	}
	bench(b, sorter)
}

func BenchmarkMergeOnlySort(b *testing.B) {
	bench(b, sort.MergeOnlySort)
}

func BenchmarkMergeWithInsertionSort(b *testing.B) {
	bench(b, sort.MergeWithInsertionSort)
}

func BenchmarkMergeSortWithBuf(b *testing.B) {
	bench(b, sort.MergeSortWithBuf)
}

func BenchmarkMergeSortWithHalfBuf(b *testing.B) {
	bench(b, sort.MergeSortWithHalfBuf)
}

func BenchmarkMergeSortParallelWithBuf(b *testing.B) {
	bench(b, sort.MergeSortParallelWithBuf)
}

type inplaceSorter = func(src []int64)

func bench(b *testing.B, sorter inplaceSorter) {
	unsorted := newUnsorted(DefaultElements)
	scracth := make([]int64, DefaultElements)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(scracth, unsorted)
		b.StartTimer()
		sorter(scracth)
	}
}

// TODO: benchhub should allow explore parameters so we can see how an algorithm scales, i.e. O(n^2) vs O(nlogn), need to refer to some hyper parameter tuning in ML and db
const DefaultElements = 16 << 20

func newUnsorted(cnt int) []int64 {
	arr := make([]int64, cnt)
	rand.Seed(time.Now().Unix())
	for i := range arr {
		arr[i] = rand.Int63()
	}
	return arr
}
