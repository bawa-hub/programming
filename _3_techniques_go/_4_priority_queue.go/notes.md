🧠 Step 1 — What Is a Heap?

A heap is just:
    A complete binary tree stored inside an array.

📦 How Tree Is Stored in Array

If index = i

| Relation    | Formula       |
| ----------- | ------------- |
| Left child  | `2*i + 1`     |
| Right child | `2*i + 2`     |
| Parent      | `(i - 1) / 2` |

Index:  0   1   2   3   4   5
Value: 50  30  40  10  20  35

        50
      /    \
     30     40
    /  \    /
   10  20  35


Max Heap:
Parent ≥ children   
       50
     /    \
    30     40

Min Heap:
Parent ≤ children
       10
     /    \
    20     30