<div dir="rtl">

# درسنامه: HashMap / Set / Frequency Counting

## ایده اصلی

وقتی لازم است membership، شمارش تکرار، یا lookup سریع داشته باشی، `map` و `set` بهترین ابزار هستند.
در Go، `map[key]value` به طور میانگین `O(1)` برای دسترسی دارد.

## سیگنال‌های تشخیص این تایپ

1. "آیا قبلاً این مقدار را دیده‌ایم؟"
2. "تعداد وقوع هر عنصر"
3. "pair/triplet با target"
4. deduplicate یا membership سریع

## الگوهای مهم

### 1) Seen Set

- هدف: فقط بررسی وجود
- ساختار: `map[int]bool` یا `map[int]struct{}`

### 2) Frequency Map

- هدف: شمارش تکرار
- ساختار: `map[T]int`

### 3) Index Map

- هدف: ذخیره آخرین/اولین اندیس دیده‌شده
- مثال: longest substring, two sum

### 4) Signature Map

- برای grouping (مثل anagram) یک signature می‌سازی
- مثال: sort string یا count-vector

## پیچیدگی

- زمان: میانگین `O(n)`
- حافظه: `O(n)`

## اشتباهات رایج

1. اعتماد مطلق به ترتیب map (در Go ترتیب map پایدار نیست)
2. فراموشی مقدار پیش‌فرض صفر برای کلید غایب
3. نادیده گرفتن برخوردهای hash از دید تئوری
4. استفاده بی‌دلیل از map وقتی آرایه شمارشی کافی است

## مثال عملی کوتاه: شمارش فراوانی

```go
func frequency(nums []int) map[int]int {
    freq := make(map[int]int)
    for _, x := range nums {
        freq[x]++
    }
    return freq
}
```

## نکته مصاحبه

در جواب شفاهی بگو: "من trade-off حافظه در برابر زمان می‌دهم تا از `O(n^2)` به `O(n)` برسم."

</div>
