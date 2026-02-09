<div dir="rtl">

# درسنامه: Prefix Sum و Difference Array

## ایده اصلی

برای محاسبه سریع مجموع بازه‌ها، از جمع‌های پیشوندی استفاده می‌کنیم:

`prefix[i] = nums[0] + ... + nums[i-1]`

پس مجموع بازه `[l..r]` می‌شود:

`prefix[r+1] - prefix[l]`

## سیگنال‌های تشخیص این تایپ

1. سوالات متعدد جمع بازه‌ای
2. "تعداد زیرآرایه‌هایی که شرط جمع دارند"
3. range update روی آرایه

## الگوهای مهم

### 1) Prefix Sum ساده

- یک‌بار `O(n)` پیش‌پردازش
- هر query جمع بازه `O(1)`

### 2) Prefix + HashMap

- برای شمارش زیرآرایه‌ها با شرط جمع
- مثال: `Subarray Sum Equals K`

### 3) Difference Array

- برای اعمال سریع افزایش روی بازه‌ها
- در پایان با یک prefix آرایه نهایی ساخته می‌شود

## پیچیدگی

- ساخت prefix: `O(n)`
- query: `O(1)`
- حافظه: `O(n)`

## اشتباهات رایج

1. آف‌بای‌وان بین `prefix` و `nums`
2. فراموش کردن مقدار اولیه `prefix[0] = 0`
3. در الگوی شمارش، فراموشی `freq[0] = 1`

## مثال عملی کوتاه

```go
func buildPrefix(nums []int) []int {
    p := make([]int, len(nums)+1)
    for i := 0; i < len(nums); i++ {
        p[i+1] = p[i] + nums[i]
    }
    return p
}

func rangeSum(prefix []int, l, r int) int {
    return prefix[r+1] - prefix[l]
}
```

## نکته مصاحبه

اگر مسئله شامل "count of subarrays" بود، تقریباً همیشه prefix+map یک کاندید جدی است.

</div>
