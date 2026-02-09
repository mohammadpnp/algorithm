<div dir="rtl">

# درسنامه: Graph (BFS/DFS/Topological Sort)

## ایده اصلی

گراف مدل رابطه بین نودهاست. بسته به نوع سوال از این ابزارها استفاده می‌کنیم:

1. DFS برای reachability و component
2. BFS برای shortest path در گراف unweighted
3. Topological Sort برای ترتیب وابستگی در DAG

## نمایش گراف

### 1) Adjacency List (پیشنهادی)

- `graph[u] = []v`
- حافظه مناسب برای گراف خلوت

### 2) Adjacency Matrix

- پیاده‌سازی ساده‌تر، ولی حافظه `O(n^2)`

## سیگنال‌های تشخیص

1. "می‌شود از A به B رسید؟"
2. "کمترین تعداد گام"
3. "پیش‌نیاز/وابستگی/ترتیب انجام"

## Topological Sort (Kahn)

1. indegree هر نود را حساب کن
2. نودهای indegree صفر را در صف بگذار
3. پاپ کن، همسایه‌ها را relax کن، اگر indegree صفر شد وارد صف
4. اگر همه نودها پردازش شدند cycle نداریم

## پیچیدگی

- DFS/BFS: `O(V+E)`
- Topological sort: `O(V+E)`

## اشتباهات رایج

1. اشتباه در جهت یال‌ها
2. فراموشی visited
3. شمارش غلط indegree در topo

## نکته مصاحبه

قبل از کدنویسی دقیق مشخص کن directed است یا undirected، weighted است یا نه.

</div>
