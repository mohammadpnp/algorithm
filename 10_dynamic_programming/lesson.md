<div dir="rtl">

# درسنامه: Dynamic Programming

## ایده اصلی

DP زمانی استفاده می‌شود که:

1. مسئله به زیرمسئله‌های همپوشان شکسته شود
2. پاسخ بهینه از پاسخ زیرمسئله‌ها ساخته شود

فرمول کلاسیک:

`dp[state] = best(dp[prev states] + transition cost)`

## سیگنال‌های تشخیص

1. "حداکثر/حداقل تعداد/هزینه/روش"
2. brute force بازگشتی کند به خاطر تکرار state
3. قید ترتیب (prefix-based)

## روش طراحی DP

1. state را تعریف کن (`dp[i]` یعنی چه؟)
2. transition بنویس
3. base caseها را مشخص کن
4. ترتیب پر کردن table را تعیین کن
5. در صورت امکان، حافظه را بهینه کن

## الگوهای رایج

1. 1D DP (مثل house robber)
2. 2D DP (مثل edit distance)
3. knapsack-style
4. interval DP
5. sequence DP (LCS/LIS)

## اشتباهات رایج

1. state مبهم که انتقال را سخت می‌کند
2. base case ناقص
3. ترتیب iteration غلط
4. استفاده از DP وقتی greedy کافی است

## نمونه کوتاه: Fibonacci با فضای O(1)

```go
func fib(n int) int {
    if n <= 1 {
        return n
    }
    a, b := 0, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}
```

## نکته مصاحبه

اگر stuck شدی، state را با جمله طبیعی تعریف کن:
"`dp[i]` یعنی بهترین جواب برای prefix تا i".

</div>
