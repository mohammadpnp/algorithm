<div dir="rtl">

# حل نمونه کامل: Koko Eating Bananas

## صورت سوال

آرایه `piles` تعداد موز هر دسته را می‌دهد. کوکو در هر ساعت با سرعت `k` از یک دسته می‌خورد.
حداقل `k` را پیدا کن که در `h` ساعت همه موزها تمام شود.

## تحلیل

- جواب بین `1` و `max(piles)` است.
- تابع `can(k)`:
  مجموع ساعت لازم برای هر دسته: `ceil(pile/k)`
- اگر `can(k)` true باشد، سرعت‌های بالاتر هم true هستند.

## کد Go

```go
package main

func minEatingSpeed(piles []int, h int) int {
    lo, hi := 1, 0
    for _, p := range piles {
        if p > hi {
            hi = p
        }
    }

    can := func(k int) bool {
        hours := 0
        for _, p := range piles {
            hours += (p + k - 1) / k // ceil(p/k)
            if hours > h {
                return false
            }
        }
        return true
    }

    for lo < hi {
        mid := lo + (hi-lo)/2
        if can(mid) {
            hi = mid
        } else {
            lo = mid + 1
        }
    }

    return lo
}
```

## Dry Run کوتاه

`piles=[3,6,7,11], h=8`

- `k=4` شدنی است
- `k=3` هم شدنی است
- `k=2` نشدنی است

پس کمینه سرعت `4`.

## پیچیدگی

- زمان: `O(n log maxPile)`
- حافظه: `O(1)`

</div>
