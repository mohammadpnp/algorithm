<div dir="rtl">

# حل نمونه کامل: Coin Change

## صورت سوال

آرایه `coins` و مقدار `amount` داده شده. کمترین تعداد سکه لازم برای ساخت `amount` را پیدا کن.
اگر ممکن نیست `-1`.

## تحلیل DP

state:

- `dp[x]` = حداقل تعداد سکه برای ساخت `x`

transition:

- `dp[x] = min(dp[x], dp[x-c] + 1)` برای هر سکه `c` که `x-c >= 0`

base:

- `dp[0] = 0`
- بقیه ابتدا `INF`

## کد Go

```go
package main

func coinChange(coins []int, amount int) int {
    const INF = int(1e9)
    dp := make([]int, amount+1)
    for i := 1; i <= amount; i++ {
        dp[i] = INF
    }

    for x := 1; x <= amount; x++ {
        for _, c := range coins {
            if x-c >= 0 && dp[x-c] != INF {
                if dp[x-c]+1 < dp[x] {
                    dp[x] = dp[x-c] + 1
                }
            }
        }
    }

    if dp[amount] == INF {
        return -1
    }
    return dp[amount]
}
```

## پیچیدگی

- زمان: `O(amount * len(coins))`
- حافظه: `O(amount)`

## نکته

این نسخه "unbounded knapsack" است چون هر سکه را بی‌نهایت بار می‌توان استفاده کرد.

</div>
