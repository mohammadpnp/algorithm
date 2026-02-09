<div dir="rtl">

# حل نمونه کامل: Lowest Common Ancestor of a Binary Tree

## صورت سوال

در یک درخت دودویی، دو گره `p` و `q` داده شده‌اند.
کمترین جد مشترک (LCA) آن‌ها را پیدا کن.

## تحلیل DFS

برای هر node:

1. اگر node خودش `p` یا `q` باشد، خود node را برگردان.
2. از چپ و راست نتیجه می‌گیریم.
3. اگر هر دو سمت غیر nil بودند، node فعلی LCA است.
4. در غیر این صورت همان نتیجه غیر nil را برمی‌گردانیم.

## کد Go

```go
package main

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || root == p || root == q {
        return root
    }

    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)

    if left != nil && right != nil {
        return root
    }
    if left != nil {
        return left
    }
    return right
}
```

## پیچیدگی

- زمان: `O(n)`
- حافظه: `O(h)` برای recursion stack

`h` ارتفاع درخت است.

</div>
