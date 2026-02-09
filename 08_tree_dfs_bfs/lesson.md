<div dir="rtl">

# درسنامه: Tree Traversal (DFS/BFS)

## ایده اصلی

درخت‌ها معمولاً با recursion (DFS) یا queue (BFS) حل می‌شوند.
برای هر گره تصمیم می‌گیری چه اطلاعاتی از فرزندان لازم داری و چه اطلاعاتی برمی‌گردانی.

## سیگنال‌های تشخیص

1. مسئله ذاتاً سلسله‌مراتبی (parent/child)
2. نیاز به پردازش سطح‌به‌سطح (BFS)
3. نیاز به پاسخ مبتنی بر زیرمسئله‌های subtree (DFS)

## الگوهای DFS

1. Preorder: node -> left -> right
2. Inorder: left -> node -> right (مفید برای BST)
3. Postorder: left -> right -> node

## الگوهای BFS

- صف نگه می‌داری
- برای level-order هر دور اندازه صف را جداگانه پردازش می‌کنی

## قالب DFS بازگشتی

```text
func dfs(node):
    if node == nil: return base
    left = dfs(node.Left)
    right = dfs(node.Right)
    return combine(node, left, right)
```

## اشتباهات رایج

1. base case ناقص (`nil`)
2. اشتباه در state مشترک بین فراخوانی‌ها
3. stack overflow در درخت خیلی عمیق (در بعضی مسائل iterative بهتر است)

## نکته مصاحبه

قبل از کدنویسی بگو: "من تابع DFS را تعریف می‌کنم که چه چیزی را برای هر node برمی‌گرداند."

</div>
