<div dir="rtl">

# حل نمونه کامل: Subarray Sum Equals K

## صورت سوال

آرایه `nums` (ممکن است منفی هم داشته باشد) و عدد `k` داده شده.
تعداد زیرآرایه‌هایی که مجموعشان دقیقاً `k` است را برگردان.

## تحلیل

روش sliding window با اعداد منفی جواب نمی‌دهد. الگوی درست:

1. `prefix` مجموع تا اندیس جاری
2. برای هر موقعیت، تعداد prefixهایی که مقدارشان `prefix-k` بوده را اضافه می‌کنیم
3. map فراوانی prefixها را نگه می‌دارد

## کد Go

```go
package main

func subarraySum(nums []int, k int) int {
    freq := map[int]int{0: 1}
    prefix, ans := 0, 0

    for _, x := range nums {
        prefix += x
        ans += freq[prefix-k]
        freq[prefix]++
    }

    return ans
}
```

## Dry Run

`nums=[1,1,1], k=2`

1. prefix=1 -> نیاز `-1`، صفر بار
2. prefix=2 -> نیاز `0`، یک بار -> ans=1
3. prefix=3 -> نیاز `1`، یک بار -> ans=2

## پیچیدگی

- زمان: `O(n)`
- حافظه: `O(n)`

## نکته

کلید طلایی: `sum(i..j)=k` معادل `prefix[j]-prefix[i-1]=k`.

</div>
