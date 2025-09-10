# reflect Package - Runtime Reflection 🔍

The `reflect` package implements runtime reflection, allowing a program to manipulate objects with arbitrary types. It's essential for building generic libraries, serialization, and dynamic programming.

## 🎯 Key Concepts

### 1. **Core Types**
- `Type` - Represents a Go type
- `Value` - Represents a Go value
- `Kind` - Represents the specific kind of type
- `StructField` - Represents a struct field
- `StructTag` - Represents a struct tag

### 2. **Type Operations**
- `TypeOf()` - Get type of value
- `New()` - Create new value of type
- `NewAt()` - Create new value at address
- `PtrTo()` - Get pointer type
- `SliceOf()` - Get slice type
- `ArrayOf()` - Get array type
- `MapOf()` - Get map type
- `ChanOf()` - Get channel type

### 3. **Value Operations**
- `ValueOf()` - Get value of interface
- `Indirect()` - Get value pointed to by pointer
- `Elem()` - Get element of pointer/slice/map/channel
- `Addr()` - Get address of value
- `CanAddr()` - Check if value can be addressed
- `CanSet()` - Check if value can be set

### 4. **Kind Types**
- `Invalid` - Invalid type
- `Bool` - Boolean
- `Int`, `Int8`, `Int16`, `Int32`, `Int64` - Integer types
- `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64` - Unsigned integer types
- `Float32`, `Float64` - Floating point types
- `Complex64`, `Complex128` - Complex types
- `Array` - Array
- `Chan` - Channel
- `Func` - Function
- `Interface` - Interface
- `Map` - Map
- `Ptr` - Pointer
- `Slice` - Slice
- `String` - String
- `Struct` - Struct
- `UnsafePointer` - Unsafe pointer

### 5. **Struct Operations**
- `NumField()` - Number of fields
- `Field()` - Get field by index
- `FieldByName()` - Get field by name
- `FieldByIndex()` - Get field by index path
- `FieldByNameFunc()` - Get field by function
- `Tag` - Get struct tag

### 6. **Function Operations**
- `NumIn()` - Number of input parameters
- `NumOut()` - Number of output parameters
- `In()` - Get input parameter type
- `Out()` - Get output parameter type
- `IsVariadic()` - Check if variadic
- `Call()` - Call function
- `CallSlice()` - Call variadic function

### 7. **Map Operations**
- `MapKeys()` - Get map keys
- `MapIndex()` - Get map value by key
- `SetMapIndex()` - Set map value by key
- `MapRange()` - Range over map

### 8. **Slice Operations**
- `Len()` - Get length
- `Cap()` - Get capacity
- `Index()` - Get element by index
- `SetIndex()` - Set element by index
- `Slice()` - Get slice
- `Slice3()` - Get slice with 3 indices

## 🚀 Common Patterns

### Type Inspection
```go
v := reflect.ValueOf(42)
t := reflect.TypeOf(42)
fmt.Printf("Type: %s, Kind: %s\n", t, v.Kind())
```

### Value Manipulation
```go
v := reflect.ValueOf(&x)
if v.CanSet() {
    v.Elem().SetInt(100)
}
```

### Struct Field Access
```go
v := reflect.ValueOf(s)
for i := 0; i < v.NumField(); i++ {
    field := v.Field(i)
    fmt.Printf("Field %d: %s\n", i, field)
}
```

## ⚠️ Common Pitfalls

1. **Panic on invalid operations** - Always check if operations are valid
2. **Performance overhead** - Reflection is slower than direct access
3. **Type safety** - Reflection bypasses compile-time type checking
4. **Memory leaks** - Be careful with reflect.Value references
5. **Interface conversion** - Use proper type assertions

## 🎯 Best Practices

1. **Check validity** - Always check if values are valid
2. **Use type assertions** - Prefer type assertions when possible
3. **Cache reflect.Type** - Store reflect.Type for reuse
4. **Handle panics** - Use recover for reflection panics
5. **Consider alternatives** - Use code generation when possible

## 🔍 Advanced Features

### Custom Type Conversion
```go
func ConvertTo(target reflect.Type, value reflect.Value) reflect.Value {
    if value.Type().ConvertibleTo(target) {
        return value.Convert(target)
    }
    return reflect.Zero(target)
}
```

### Deep Copy
```go
func DeepCopy(src reflect.Value) reflect.Value {
    if !src.IsValid() {
        return reflect.Value{}
    }
    
    switch src.Kind() {
    case reflect.Ptr:
        return reflect.New(src.Type().Elem())
    case reflect.Slice:
        return reflect.MakeSlice(src.Type(), src.Len(), src.Cap())
    // ... handle other types
    }
    return src
}
```

## 📚 Real-world Applications

1. **Serialization** - JSON, XML, binary encoding
2. **ORM** - Object-relational mapping
3. **Configuration** - Dynamic configuration loading
4. **Testing** - Test framework utilities
5. **Code Generation** - Template engines

## 🧠 Memory Tips

- **reflect** = **R**untime **E**xamination **F**or **L**earning **E**very **C**omponent **T**ype
- **Type** = **T**ype information
- **Value** = **V**alue information
- **Kind** = **K**ind of type
- **Field** = **F**ield of struct
- **Tag** = **T**ag of field
- **Call** = **C**all function
- **Elem** = **E**lement of pointer

Remember: The reflect package is your gateway to dynamic programming in Go! 🎯
