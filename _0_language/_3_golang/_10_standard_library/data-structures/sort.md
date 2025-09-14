# sort Package - Sorting Algorithms 📊

The `sort` package provides primitives for sorting slices and user-defined collections. It's essential for data organization, searching, and algorithm implementation.

## 🎯 Key Concepts

### 1. **Basic Sorting Functions**
- `Sort()` - Sort slice in ascending order
- `Stable()` - Stable sort (preserves equal elements' order)
- `IsSorted()` - Check if slice is sorted
- `Search()` - Binary search in sorted slice
- `SearchInts()` - Binary search for ints
- `SearchFloat64s()` - Binary search for float64s
- `SearchStrings()` - Binary search for strings

### 2. **Sorting Interfaces**
- `Interface` - Interface for sorting
- `IntSlice` - Int slice with sorting methods
- `Float64Slice` - Float64 slice with sorting methods
- `StringSlice` - String slice with sorting methods

### 3. **Interface Methods**
- `Len()` - Length of collection
- `Less(i, j int) bool` - Compare elements at indices i and j
- `Swap(i, j int)` - Swap elements at indices i and j

### 4. **Custom Sorting**
- Implement `sort.Interface` for custom types
- Use `sort.Slice()` for function-based sorting
- Use `sort.SliceStable()` for stable function-based sorting

### 5. **Search Functions**
- `Search()` - Binary search with custom comparison
- `SearchInts()` - Binary search for int slice
- `SearchFloat64s()` - Binary search for float64 slice
- `SearchStrings()` - Binary search for string slice

## 🚀 Common Patterns

### Basic Sorting
```go
numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
sort.Ints(numbers)
fmt.Println(numbers) // [1 1 2 3 4 5 6 9]
```

### Custom Sorting
```go
type Person struct {
    Name string
    Age  int
}

people := []Person{{"Alice", 30}, {"Bob", 25}}
sort.Slice(people, func(i, j int) bool {
    return people[i].Age < people[j].Age
})
```

### Binary Search
```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
index := sort.SearchInts(numbers, 5)
fmt.Println(index) // 4
```

## ⚠️ Common Pitfalls

1. **Not implementing interface correctly** - Ensure all methods are implemented
2. **Inconsistent comparison** - Less function must be consistent
3. **Searching unsorted data** - Binary search requires sorted data
4. **Index out of bounds** - Check indices in Less and Swap methods
5. **Performance considerations** - Choose appropriate sorting algorithm

## 🎯 Best Practices

1. **Use built-in types** - Prefer Ints, Float64s, Strings for simple cases
2. **Implement interface correctly** - Ensure all methods work together
3. **Use stable sort** - When order of equal elements matters
4. **Check if sorted** - Use IsSorted before binary search
5. **Consider performance** - Choose appropriate sorting method

## 🔍 Advanced Features

### Custom Sort Interface
```go
type CustomSlice []int

func (c CustomSlice) Len() int           { return len(c) }
func (c CustomSlice) Less(i, j int) bool { return c[i] < c[j] }
func (c CustomSlice) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
```

### Function-based Sorting
```go
sort.Slice(data, func(i, j int) bool {
    return data[i].Field < data[j].Field
})
```

### Reverse Sorting
```go
sort.Sort(sort.Reverse(data))
```

## 📚 Real-world Applications

1. **Data Analysis** - Sorting datasets for analysis
2. **Search Optimization** - Preparing data for binary search
3. **User Interfaces** - Sorting tables and lists
4. **Algorithm Implementation** - Building sorting algorithms
5. **Performance Optimization** - Choosing optimal sorting method

## 🧠 Memory Tips

- **sort** = **S**orting **O**perations **R**eference **T**oolkit
- **Sort** = **S**ort slice
- **Stable** = **S**table sort
- **Search** = **S**earch sorted data
- **Interface** = **I**nterface for sorting
- **Less** = **L**ess comparison
- **Swap** = **S**wap elements

Remember: The sort package is your gateway to data organization in Go! 🎯
