<div dir="rtl">

# حل نمونه کامل: Two Sum

## صورت سوال

آرایه `nums` و عدد `target` داده شده. دو اندیس پیدا کن که مجموع مقادیرشان برابر `target` باشد.
فرض کن دقیقاً یک پاسخ وجود دارد و نمی‌توان یک عنصر را دوبار استفاده کرد.

## تحلیل

brute force برابر `O(n^2)` است. بهینه:

1. با پیمایش هر `x`، مکمل آن را حساب کن: `need = target - x`
2. اگر `need` در map دیده شده باشد، پاسخ پیدا شده
3. در غیر این صورت `x` و اندیسش را ذخیره کن

## کد Go

```go
package main

func twoSum(nums []int, target int) []int {
    idx := make(map[int]int)

    for i, x := range nums {
        need := target - x
        if j, ok := idx[need]; ok {
            return []int{j, i}
        }
        idx[x] = i
    }

    return nil
}
```

## Dry Run

`nums = [2,7,11,15], target = 9`

1. `2` را می‌بینیم، نیاز `7` است، نیست -> `2:0` ذخیره می‌شود
2. `7` را می‌بینیم، نیاز `2` است، هست -> پاسخ `[0,1]`

## پیچیدگی

- زمان: `O(n)`
- حافظه: `O(n)`

## نکته

اول check کن بعد insert؛ وگرنه در حالت `target=2*x` ممکن است همان عنصر دوبار انتخاب شود.

</div>
