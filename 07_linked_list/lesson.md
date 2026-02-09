<div dir="rtl">

# درسنامه: Linked List Patterns

## ایده اصلی

در لیست پیوندی دسترسی تصادفی نداریم، اما حذف/درج موضعی با pointer manipulation سریع است.
کلید حل، مدیریت درست اشاره‌گرها و استفاده از `dummy` node است.

## سیگنال‌های تشخیص

1. حذف/درج node با یک پیمایش
2. پیدا کردن میانه یا cycle
3. معکوس‌سازی کل یا بخشی از لیست

## الگوهای مهم

### 1) Dummy Node

- برای ساده‌سازی حذف/درج در ابتدای لیست

### 2) Fast/Slow Pointers

- یافتن mid
- تشخیص cycle (Floyd)

### 3) Reverse Linked List

- با سه pointer: `prev`, `curr`, `next`

### 4) Two Pass vs One Pass

- بعضی مسائل مثل remove nth from end با one-pass و فاصله n قابل حل‌اند

## اشتباهات رایج

1. از دست دادن `next` قبل از تغییر لینک
2. عدم مدیریت head (با dummy قابل حل)
3. nil pointer panic

## اسکلت reverse

```go
func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    cur := head
    for cur != nil {
        nxt := cur.Next
        cur.Next = prev
        prev = cur
        cur = nxt
    }
    return prev
}
```

## نکته مصاحبه

حین کدنویسی با صدای بلند بگو هر pointer الان به چه بخشی اشاره می‌کند. این کار خطا را کم می‌کند.

</div>
