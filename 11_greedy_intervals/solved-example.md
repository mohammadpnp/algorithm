<div dir="rtl">

# حل نمونه کامل: Merge Intervals

## صورت سوال

مجموعه‌ای از بازه‌ها داده شده. تمام بازه‌های هم‌پوشان را ادغام کن.

## تحلیل

1. بازه‌ها را بر اساس `start` صعودی مرتب کن.
2. لیست پاسخ را با اولین بازه شروع کن.
3. هر بازه جدید:
   - اگر با آخرین بازه پاسخ overlap دارد، `end` را ماکزیمم کن
   - وگرنه بازه جدید را append کن

## کد Go

```go
package main

import "sort"

func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return nil
    }

    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    ans := [][]int{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        last := ans[len(ans)-1]
        cur := intervals[i]

        if cur[0] <= last[1] {
            if cur[1] > last[1] {
                last[1] = cur[1]
            }
        } else {
            ans = append(ans, cur)
        }
    }

    return ans
}
```

## پیچیدگی

- زمان: `O(n log n)` به خاطر sort
- حافظه: `O(n)` برای خروجی

## نکته

اگر sort انجام نشود، تشخیص هم‌پوشانی به‌صورت خطی ممکن نیست.

</div>
