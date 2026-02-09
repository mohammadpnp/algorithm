<div dir="rtl">

# حل نمونه کامل: Daily Temperatures

## صورت سوال

آرایه `temperatures` داده شده. برای هر روز مشخص کن چند روز باید صبر کنی تا دمای بالاتری بیاید.
اگر روز گرم‌تری وجود ندارد، `0`.

## تحلیل

Monotonic decreasing stack از اندیس روزها نگه می‌داریم:

1. تا وقتی دمای امروز از دمای اندیس بالای پشته بیشتر است، آن اندیس را pop می‌کنیم.
2. فاصله روزها پاسخ آن اندیس است.
3. اندیس امروز را push می‌کنیم.

## کد Go

```go
package main

func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    ans := make([]int, n)
    st := make([]int, 0, n) // stack of indices

    for i := 0; i < n; i++ {
        for len(st) > 0 && temperatures[i] > temperatures[st[len(st)-1]] {
            j := st[len(st)-1]
            st = st[:len(st)-1]
            ans[j] = i - j
        }
        st = append(st, i)
    }

    return ans
}
```

## Dry Run

`[73,74,75,71,69,72,76,73]`

- روز 0 با دیدن 74 حل می‌شود (`1` روز)
- روز 1 با دیدن 75 حل می‌شود
- روز 2 با دیدن 76 حل می‌شود (`4` روز)
- ...

پاسخ: `[1,1,4,2,1,1,0,0]`

## پیچیدگی

- زمان: `O(n)`
- حافظه: `O(n)`

</div>
