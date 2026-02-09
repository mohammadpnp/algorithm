<div dir="rtl">

# حل نمونه کامل: Longest Substring Without Repeating Characters

## صورت سوال

یک رشته `s` داده شده. طول بلندترین زیررشته‌ای که کاراکتر تکراری ندارد را برگردان.

## تحلیل

این مسئله sliding window متغیر است:

1. `right` را جلو می‌بریم و کاراکتر جدید را وارد پنجره می‌کنیم.
2. اگر این کاراکتر تکراری شد، `left` را جلو می‌بریم تا پنجره دوباره بدون تکرار شود.
3. طول پنجره معتبر را به عنوان پاسخ کاندید می‌گیریم.

برای تشخیص تکرار، از `map[byte]int` به عنوان شمارنده استفاده می‌کنیم.

## کد Go

```go
package main

func lengthOfLongestSubstring(s string) int {
    freq := make(map[byte]int)
    left, ans := 0, 0

    for right := 0; right < len(s); right++ {
        c := s[right]
        freq[c]++

        for freq[c] > 1 {
            freq[s[left]]--
            left++
        }

        if right-left+1 > ans {
            ans = right - left + 1
        }
    }

    return ans
}
```

## Dry Run کوتاه

`s = "abcabcbb"`

1. پنجره تا `"abc"` معتبر است، پاسخ `3`
2. با ورود `a` تکرار ایجاد می‌شود، `left` جلو می‌رود
3. هیچ پنجره‌ای بزرگ‌تر از `3` ساخته نمی‌شود

پاسخ نهایی: `3`

## پیچیدگی

- زمان: `O(n)`
- حافظه: `O(min(n, charset))`

## چرا این روش بهتر از brute force است؟

brute force تمام زیررشته‌ها را چک می‌کند (`O(n^2)` تا `O(n^3)`)،
اما sliding window با یک پیمایش خطی حل می‌کند.

</div>
