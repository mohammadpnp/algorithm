<div dir="rtl">

# درسنامه: Backtracking

## ایده اصلی

در بک‌ترکینگ فضای جستجو را به شکل درخت تصمیم طی می‌کنی:

1. یک انتخاب انجام می‌دهی
2. وارد سطح بعد می‌شوی
3. بعد از برگشت، اثر انتخاب را rollback می‌کنی

برای مسائل "تولید تمام حالت‌های معتبر" بسیار مناسب است.

## سیگنال‌های تشخیص

1. "تمام permutation/combination/subsetها"
2. قیدهایی که حین ساخت جواب باید بررسی شوند
3. نیاز به enumerate پاسخ‌ها

## قالب ذهنی

```text
dfs(state):
    if goal:
        collect answer
        return

    for choice in choices:
        if not valid(choice):
            continue
        apply(choice)
        dfs(next_state)
        rollback(choice)
```

## تکنیک‌های pruning

1. برش با قیدها (مثلاً sum از target گذشت)
2. sort قبل از شروع برای حذف تکراری‌ها
3. visited برای permutation
4. memoization در صورت وجود state تکراری

## اشتباهات رایج

1. rollback ناقص
2. append مستقیم slice بدون کپی در پاسخ
3. branchهای تکراری

## مثال کوتاه: subsets

```go
func subsets(nums []int) [][]int {
    ans := [][]int{}
    cur := []int{}

    var dfs func(int)
    dfs = func(i int) {
        if i == len(nums) {
            tmp := append([]int(nil), cur...)
            ans = append(ans, tmp)
            return
        }

        dfs(i + 1)
        cur = append(cur, nums[i])
        dfs(i + 1)
        cur = cur[:len(cur)-1]
    }

    dfs(0)
    return ans
}
```

## نکته مصاحبه

در شروع حل بگو: "من درخت تصمیم را تعریف می‌کنم و هر سطح یعنی انتخاب چه چیزی."

</div>
