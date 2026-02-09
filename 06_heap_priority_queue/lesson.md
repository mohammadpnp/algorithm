<div dir="rtl">

# درسنامه: Heap / Priority Queue

## ایده اصلی

وقتی لازم است "کوچک‌ترین" یا "بزرگ‌ترین" عنصر را بارها برداری یا مجموعه‌ای از top-k را نگه داری، هیپ ابزار اصلی است.

در Go از پکیج `container/heap` استفاده می‌شود.

## سیگنال‌های تشخیص

1. `k`-th smallest/largest
2. top-k frequent
3. stream data با نیاز به کمینه/بیشینه جاری
4. merge چند لیست مرتب

## الگوهای مهم

### 1) Min-Heap برای k بزرگ‌ترین

- اندازه هیپ را `k` نگه می‌داری
- هر مقدار جدید را push کن
- اگر سایز > k شد، min را pop کن
- در پایان ریشه heap همان k-th largest است

### 2) Max-Heap برای استخراج مکرر بیشینه

- برای شبیه‌سازی یا greedyهای خاص

### 3) Two Heaps

- یک max-heap برای نیمه کوچک، یک min-heap برای نیمه بزرگ
- مثال: median from data stream

## پیچیدگی

- push/pop: `O(log n)`
- peek: `O(1)`

## اشتباهات رایج در Go

1. فراموشی پیاده‌سازی `Len`, `Less`, `Swap`, `Push`, `Pop`
2. اشتباه در pointer receiver برای `Push/Pop`
3. confusion بین min-heap پیش‌فرض و max-heap سفارشی

## نمونه اسکلت MinHeap در Go

```go
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
    old := *h
    x := old[len(old)-1]
    *h = old[:len(old)-1]
    return x
}
```

## نکته مصاحبه

به جای حفظ کردن API، یک template آماده در ذهن داشته باش تا سریع تایپ کنی.

</div>
