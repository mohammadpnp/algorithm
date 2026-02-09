<div dir="rtl">

# حل نمونه کامل: Permutations

## صورت سوال

آرایه‌ای از اعداد distinct داده شده. تمام permutationهای ممکن را برگردان.

## تحلیل

1. در هر سطح یک عدد انتخاب می‌کنیم که هنوز استفاده نشده
2. `used[i]` نشان می‌دهد عدد `nums[i]` قبلاً در مسیر فعلی استفاده شده یا نه
3. وقتی طول مسیر برابر `n` شد یک پاسخ کامل داریم

## کد Go

```go
package main

func permute(nums []int) [][]int {
    n := len(nums)
    used := make([]bool, n)
    cur := make([]int, 0, n)
    ans := make([][]int, 0)

    var dfs func()
    dfs = func() {
        if len(cur) == n {
            tmp := append([]int(nil), cur...)
            ans = append(ans, tmp)
            return
        }

        for i := 0; i < n; i++ {
            if used[i] {
                continue
            }
            used[i] = true
            cur = append(cur, nums[i])

            dfs()

            cur = cur[:len(cur)-1]
            used[i] = false
        }
    }

    dfs()
    return ans
}
```

## پیچیدگی

- زمان: `O(n * n!)`
- حافظه: `O(n)` برای recursion + مسیر جاری

## نکته

در سوالات با اعداد تکراری باید از sorting + skip duplicates استفاده کنی.

</div>
