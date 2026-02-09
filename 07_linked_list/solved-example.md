<div dir="rtl">

# حل نمونه کامل: Linked List Cycle

## صورت سوال

سر لیست پیوندی داده شده. مشخص کن آیا لیست cycle دارد یا خیر.

## تحلیل

از الگوریتم Floyd (tortoise and hare):

1. `slow` هر بار یک گام، `fast` هر بار دو گام
2. اگر cycle باشد، این دو نهایتاً ملاقات می‌کنند
3. اگر `fast` یا `fast.Next` nil شود، cycle نداریم

## کد Go

```go
package main

type ListNode struct {
    Val  int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }

    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }

    return false
}
```

## پیچیدگی

- زمان: `O(n)`
- حافظه: `O(1)`

## چرا درست است؟

در حالت cycle، fast نسبت به slow یک گام اضافه در هر دور می‌گیرد و روی دایره بالاخره به slow می‌رسد.

</div>
