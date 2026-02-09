<div dir="rtl">

# حل نمونه کامل: Course Schedule

## صورت سوال

`numCourses` درس داریم و آرایه `prerequisites` که هر زوج `[a,b]` یعنی برای گرفتن `a` باید `b` را قبلاً گرفته باشی.
بررسی کن آیا می‌توان همه درس‌ها را پاس کرد یا نه.

## تحلیل

این همان تشخیص cycle در directed graph است.
از Kahn Topological Sort استفاده می‌کنیم:

1. گراف و indegree را می‌سازیم
2. همه indegree صفرها در صف
3. تا صف خالی نشده پردازش می‌کنیم
4. اگر تعداد پردازش‌شده‌ها == `numCourses`، cycle نداریم

## کد Go

```go
package main

func canFinish(numCourses int, prerequisites [][]int) bool {
    g := make([][]int, numCourses)
    indeg := make([]int, numCourses)

    for _, p := range prerequisites {
        a, b := p[0], p[1]
        g[b] = append(g[b], a)
        indeg[a]++
    }

    q := make([]int, 0)
    for i := 0; i < numCourses; i++ {
        if indeg[i] == 0 {
            q = append(q, i)
        }
    }

    processed := 0
    for head := 0; head < len(q); head++ {
        u := q[head]
        processed++
        for _, v := range g[u] {
            indeg[v]--
            if indeg[v] == 0 {
                q = append(q, v)
            }
        }
    }

    return processed == numCourses
}
```

## پیچیدگی

- زمان: `O(V+E)`
- حافظه: `O(V+E)`

</div>
