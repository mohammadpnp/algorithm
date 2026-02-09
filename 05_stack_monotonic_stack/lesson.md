<div dir="rtl">

# درسنامه: Stack و Monotonic Stack

## ایده اصلی

پشته برای مسائلی عالی است که "آخرین کاندید باز" باید سریعاً قابل دسترسی باشد.
Monotonic stack نسخه‌ای از پشته است که اعضا را با ترتیب صعودی/نزولی نگه می‌دارد تا next greater/smaller را سریع پیدا کنیم.

## سیگنال‌های تشخیص

1. "عنصر بعدی بزرگ‌تر/کوچک‌تر"
2. "تا وقتی شرط برقرار است pop کن"
3. مسئله‌ای که به nearest previous/next وابسته است

## الگوهای مهم

### 1) پشته معمولی

- پرانتز معتبر
- undo/redo
- evaluation expression

### 2) Monotonic Decreasing Stack

- برای next greater element
- وقتی عنصر بزرگ‌تر می‌آید، اعضای کوچک‌تر pop می‌شوند

### 3) Monotonic Increasing Stack

- برای next smaller یا histogram

## قالب ذهنی

```text
for i over array:
    while stack not empty and arr[i] violates monotonic condition:
        popped = stack.pop()
        update answer for popped
    stack.push(i)
```

## پیچیدگی

- زمان: `O(n)` چون هر عنصر حداکثر یک بار push و یک بار pop می‌شود
- حافظه: `O(n)`

## اشتباهات رایج

1. ذخیره مقدار به جای اندیس وقتی اندیس لازم است
2. نادیده گرفتن عناصر باقی‌مانده در پشته بعد از پایان حلقه
3. استفاده اشتباه از `>` و `>=` (تفاوت در handling عناصر مساوی)

## مثال عملی کوتاه: Valid Parentheses

```go
func isValid(s string) bool {
    st := make([]byte, 0, len(s))
    match := map[byte]byte{')': '(', ']': '[', '}': '{'}

    for i := 0; i < len(s); i++ {
        c := s[i]
        if c == '(' || c == '[' || c == '{' {
            st = append(st, c)
            continue
        }
        if len(st) == 0 || st[len(st)-1] != match[c] {
            return false
        }
        st = st[:len(st)-1]
    }
    return len(st) == 0
}
```

## نکته مصاحبه

بگو invariant چیست: "اندیس‌های داخل پشته همیشه به صورت یکنواخت نزولی/صعودی نگه داشته می‌شوند."

</div>
