<div dir="rtl">

# حل نمونه کامل: Kth Largest Element in an Array

## صورت سوال

آرایه `nums` و عدد `k` داده شده. عنصر `k`-امین بزرگ را پیدا کن.

## تحلیل

با min-heap اندازه `k`:

1. هر عدد را push می‌کنیم.
2. اگر اندازه از `k` بیشتر شد، کوچک‌ترین را pop می‌کنیم.
3. در پایان، کوچک‌ترین عنصر میان top-k عناصر در هیپ باقی می‌ماند که همان `k`-امین بزرگ است.

## کد Go

```go
package main

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
    old := *h
    x := old[len(old)-1]
    *h = old[:len(old)-1]
    return x
}

func findKthLargest(nums []int, k int) int {
    h := &MinHeap{}
    heap.Init(h)

    for _, x := range nums {
        heap.Push(h, x)
        if h.Len() > k {
            heap.Pop(h)
        }
    }

    return (*h)[0]
}
```

## پیچیدگی

- زمان: `O(n log k)`
- حافظه: `O(k)`

## نکته

اگر `k` خیلی کوچک‌تر از `n` باشد، این روش بهتر از sort کامل (`O(n log n)`) است.

</div>
