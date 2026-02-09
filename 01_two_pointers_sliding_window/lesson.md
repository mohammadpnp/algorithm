<div dir="rtl">

# درسنامه: Two Pointers و Sliding Window

## ایده اصلی

وقتی ورودی یک آرایه/رشته خطی است و دنبال بازه یا زوج اندیس هستی، به جای بررسی همه حالت‌ها (`O(n^2)`)، با دو اشاره‌گر (`left`, `right`) یک‌بار داده را پیمایش می‌کنی.

## سیگنال‌های تشخیص این تایپ

1. عبارت‌هایی مثل "longest/shortest subarray/substring"
2. قیدهای "at most / at least / exactly k"
3. نیاز به شمارش/کنترل اعضای یک بازه پیوسته
4. داده sorted است و دنبال pair با شرط خاص هستی

## الگوهای مهم

### 1) Opposite Pointers روی آرایه مرتب

- یک اشاره‌گر ابتدای آرایه، یکی انتهای آرایه
- با توجه به مقایسه با target یکی را جابه‌جا می‌کنی
- مثال: `Two Sum II`

### 2) Fast/Slow Pointers

- هر دو از ابتدا حرکت می‌کنند ولی با سرعت متفاوت
- مثال: تشخیص cycle در linked list

### 3) Sliding Window با اندازه ثابت

- اندازه پنجره `k` ثابت است
- با اضافه کردن راست و حذف چپ، مقدار aggregate را آپدیت می‌کنی
- مثال: بیشینه مجموع زیرآرایه طول `k`

### 4) Sliding Window با اندازه متغیر

- `right` جلو می‌رود تا شرط نقض شود
- با حرکت `left` شرط را برمی‌گردانی
- مثال: کوچک‌ترین زیرآرایه با مجموع حداقل `target`

## قالب ذهنی استاندارد

```text
left = 0
for right in [0..n-1]:
    add nums[right]
    while window invalid:
        remove nums[left]
        left++
    update answer
```

## پیچیدگی

- زمان: معمولاً `O(n)` چون هر عنصر حداکثر یک‌بار وارد و یک‌بار خارج پنجره می‌شود
- حافظه: بسته به نیاز مسئله (`O(1)` یا `O(alphabet)` یا `O(n)`)

## اشتباهات رایج

1. اشتباه در شرط حلقه داخلی `while`
2. فراموش کردن آپدیت پاسخ بعد از معتبر شدن پنجره
3. آف‌بای‌وان در محاسبه طول: `right-left+1`
4. جابه‌جایی دیرهنگام یا زودهنگام `left`

## مثال عملی کوتاه

مسئله: کمترین طول زیرآرایه با مجموع حداقل `target`.

```go
func minSubArrayLen(target int, nums []int) int {
    n := len(nums)
    left, sum := 0, 0
    ans := n + 1

    for right := 0; right < n; right++ {
        sum += nums[right]
        for sum >= target {
            if right-left+1 < ans {
                ans = right - left + 1
            }
            sum -= nums[left]
            left++
        }
    }

    if ans == n+1 {
        return 0
    }
    return ans
}
```

## نکته مصاحبه

همیشه شفاف بگو پنجره چه invariantی را نگه می‌دارد؛ مثلاً:
"در هر لحظه پنجره‌ی `[left..right]` فاقد کاراکتر تکراری است."

</div>
